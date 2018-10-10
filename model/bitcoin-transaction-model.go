package model

import (
	"github.com/jinzhu/gorm"
	"log"
	"github.com/shopspring/decimal"
	"wallet-bitcoin/setting"
)


type BtcTransactionInfo struct {
	gorm.Model
	Hash string `gorm:"type:varchar(100);not null;unique"`
	FromAddress string `gorm:"type:varchar(100);not null"`
	ToAddress string `gorm:"type:varchar(100);not null"`
	Fee string `gorm:"type:varchar(100);not null"`
	Amount string `gorm:"type:varchar(100);not null"`
	Remark string `gorm:"type:varchar(100);not null"`
	IsConfirm bool `gorm:"default:0"`
}

var db *gorm.DB

func init() {
	var err error
	db,err = gorm.Open("mysql",setting.ServerDBConnectString)
	if nil != err{
		log.Println(err)
		return
	}

	db.AutoMigrate(&BtcTransactionInfo{})
}

func InitBtcTransactionInfo() *BtcTransactionInfo {
	return &BtcTransactionInfo{}
}

/// 持久化到数据库
func (btf *BtcTransactionInfo) Store2Db()  error {
	if err := db.Create(btf).Error;err != nil {
		return err
	}
	return nil
}



func GetRemarkInfo(hash string) string  {
	var info = &BtcTransactionInfo{}
	db.Where("hash = ?", hash).First(info)

	return info.Remark
}

/////////////////////////////////////////////////////////////////////////////////
type BtcTransaction struct{
	Hash string
	FromAddress string
	ToAddress string
	Fee string
	Amount string
	Remark string
	IsConfirm bool
}

func InitBtcTransacton() *BtcTransaction {
	return &BtcTransaction{}
}

func GetBtcTransaction(address string,info TxInfo) (trs []BtcTransaction,err error)  {

	trs = make([]BtcTransaction,0)

	for _,tx := range info.Txs {

		btcTx := BtcTransaction{}
		btcTx.Hash = tx.Hash
		btcTx.FromAddress,btcTx.ToAddress = GetTransactionParti(address,tx)

		//过虑掉别人发的币，但还没有被打包确认的记录
		if btcTx.ToAddress == address && "" == tx.Block_height {
			continue
		}

		btcTx.Fee,btcTx.Amount = GetFeeAmount(address,tx)
		btcTx.IsConfirm = (tx.Block_height.String() != "")

		// 从自己平台获取交易的备注信息
		btcTx.Remark = GetRemarkInfo(btcTx.Hash)

		trs = append(trs,btcTx)
	}

	return

}

// 获取交易的双方地址
// 假设所有的交易只有一笔
func GetTransactionParti(address string,tx Tx) (from string,to string)  {
	from = ""
	to = ""

	// 获取 发送方 的地址
	var otherAddress string = ""
	for _,inputElmt := range tx.Inputs {
		if inputElmt.Prev_out.Addr == address {
			from = address
			break
		}
		otherAddress = inputElmt.Prev_out.Addr
	}

	// 获取 接收方 的地址
	// WARNING: 仅考虑笔交易的情况
	if "" != from {
		for _,outElmt := range tx.Out {
			if outElmt.Addr != address {
				to = outElmt.Addr
				break;
			}
		}
		return
	}

	from = otherAddress
	to = address

	return
}

// 获取 矿工费 及 交易额
func GetFeeAmount(address string,tx Tx) (fee string,amount string) {
	fee = "0"
	amount = "0"

	amountIn,_ := decimal.NewFromString("0")
	amountOut,_ := decimal.NewFromString("0")
	feeIn,_ := decimal.NewFromString("0")
	feeOut,_ := decimal.NewFromString("0")

	// 计算输入总额
	for _,inputElmt := range tx.Inputs {
		perInValue,err := decimal.NewFromString(inputElmt.Prev_out.Value.String())
		if nil == err {
			feeIn = feeIn.Add(perInValue)
		}

		// 额度
		if inputElmt.Prev_out.Addr == address {
			perAmountValueIn,err := decimal.NewFromString(inputElmt.Prev_out.Value.String())
			if nil == err {
				amountIn = amountIn.Add(perAmountValueIn)
			}
		}
	}

	// 计算输出总额
	for _,outElmt := range tx.Out {
		perOutValue,err := decimal.NewFromString(outElmt.Value.String())
		if nil == err {
			feeOut = feeOut.Add(perOutValue)
		}

		if outElmt.Addr == address {
			perAmountValueOut,err := decimal.NewFromString(outElmt.Value.String())
			if nil == err {
				amountOut = amountOut.Add(perAmountValueOut)
			}
		}
	}


	feeV := feeOut.Sub(feeIn).Abs()
	fee = feeV.String()

	amountV := amountOut.Sub(amountIn)

	// 发币的情况
	if amountV.LessThan(decimal.Zero) {
		amountV = amountV.Abs().Sub(feeV)
	}

	amount = amountV.String()

	return
}
