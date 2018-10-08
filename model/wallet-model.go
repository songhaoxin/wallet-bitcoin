package model

import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"github.com/jmoiron/sqlx"
	"time"
	"regexp"
)

const ServerDBConnectString  = "root:current@tcp(47.75.194.231)/multi_wallet?charset=utf8&parseTime=True&loc=Local"

var Db *sqlx.DB

func init() {
	database,err := sqlx.Open("mysql",ServerDBConnectString)
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



func (w *Wallet) AddBtcAccount(walletId int,address string) bool {
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
func (w *Wallet) IsExistBtc(walletId int,address string) bool  {

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

func (w *Wallet) IsExistWallet(walletId int) bool {

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