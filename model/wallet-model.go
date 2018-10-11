package model

import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"github.com/jmoiron/sqlx"
	"time"
	"regexp"
	"wallet-bitcoin/setting"
)



var Db *sqlx.DB

func init() {
	database,err := sqlx.Open("mysql",setting.ServerDBConnectString)
	if nil != err {
		log.Println("open mysql failed",err)
		return
	}
	Db = database
}


type Wallet struct {
	db *sql.DB
}

func Init() *Wallet {
	return &Wallet{}
}

// 根据钱包 id 删除钱包
func RemoveWallet(walletId int64) bool {
	conn,err := Db.Begin()

	if nil != err {
		log.Println(err)
		return false
	}
	// 先删除对应钱包的所有的币种信息列表
	stmt1,_ := Db.Prepare("DELETE FROM coin_address WHERE wallet_id_id=?")
	defer stmt1.Close()
	ret,err := stmt1.Exec(walletId)
	if nil != err {
		log.Println(err)
		conn.Rollback()
		return false
	}

	// 再删除钱包信息
	stmt2,_ := Db.Prepare("DELETE FROM wallet WHERE id=?")
	defer stmt2.Close()
	ret,err = stmt2.Exec(walletId)
	if nil != err {
		log.Println(err)
		conn.Rollback()
		return false
	}

	if RowsAffected, err := ret.RowsAffected(); nil == err {
		log.Println("RowsAffected:", RowsAffected)
	}

	return true
}


// 根据钱包 ID 及 币种(ETH/BTC) 删除相关信息 实际上用不上
func RemoveAccount(walletId int64,coinType string) bool  {
	if (coinType != "BTC") && (coinType != "ETH") { return false}
	//var walletIdStr string = strconv.FormatInt(walletId,10)

	conn,err := Db.Begin()

	if nil != err {
		log.Println(err)
		return false
	}

	stmt,_ := Db.Prepare("DELETE FROM coin_address WHERE wallet_id_id=?")
	defer stmt.Close()

	//执行删除操作
	ret,err := stmt.Exec(walletId)
	if nil != err {
		log.Println(err)
		conn.Rollback()
		return false
	}

	if RowsAffected, err := ret.RowsAffected(); nil == err {
		log.Println("RowsAffected:", RowsAffected)
	}

	// 查看钱包中是否还存在 帐户，
	var address string
	err = Db.QueryRow("SELECT wallet_id_id FROM coin_address WHERE wallet_id_id=? LIMIT 1",walletId).Scan(&address)
	if nil != err {
		return false
	}
	if "" == address {
		ret,err := Db.Exec("DELETE FROM wallet where id=?",walletId)
		if nil != err {
			log.Println(err)
			conn.Rollback()
		}
		if RowsAffected, err := ret.RowsAffected(); nil == err {
			log.Println("RowsAffected:", RowsAffected)
		}
	}

	return true
}




func (w *Wallet) AddBtcAccount(walletId int64,address string) bool {
	stmt,_ := Db.Prepare("INSERT INTO coin_address (`wallet_id_id`,`address`,`type`,`created_time`) VALUES (?,?,?,?)")
	defer stmt.Close()


	ret,err := stmt.Exec(walletId,address,"BTC",time.Now())
	if nil != err {
		log.Println("增加BTC帐户失败:%v",err)
		return false
	}

	if LastInsertId, err := ret.LastInsertId(); nil == err {
		log.Println("LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		log.Println("RowsAffected:", RowsAffected)
	}
	return true
}

// 指定的地址是否存在了BTC
func (w *Wallet) IsExistBtc(walletId int64,address string) bool  {

	if nil == Db {
		log.Println("open mysql failed")
		return false
	}

	var address2 string

	err := Db.QueryRow("SELECT address FROM coin_address WHERE wallet_id_id=? AND type='BTC' AND address=? LIMIT 1",walletId,address).Scan(&address2)
	if nil != err {
		return false
	}

	return true
}

func (w *Wallet) IsExistWallet(walletId int64) bool {

	if nil == Db {
		log.Println("open mysql failed")
		return false
	}

	var id string
	err := Db.QueryRow("SELECT `id` FROM wallet WHERE `id`=? LIMIT 1", walletId).Scan(&id)
	if nil != err {
		return false
	}

	return true
}


////////////////////////////////////////////////////////////////////////////////////////////
func (w *Wallet) UpdatePhoneNumber(phoneNumber string,walletId int) bool {
	if nil == Db {
		log.Println("open mysql failed")
		return false
	}

	stmt,_ := Db.Prepare("UPDATE `wallet` SET `phone_num`=? WHERE `id`=?")
	defer stmt.Close()

	ret,err := stmt.Exec(phoneNumber,walletId)
	if nil != err {
		log.Println("修改电话号码失败:%v",err)
		return false
	}

	if LastInsertId, err := ret.LastInsertId(); nil == err {
		log.Println("LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		log.Println("RowsAffected:", RowsAffected)
	}

	return true
}

func CheckPhoneNumber(phoneNumber string) bool {
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phoneNumber)
}