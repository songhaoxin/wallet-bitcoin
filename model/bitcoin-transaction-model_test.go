package model

import (
	"testing"
	"fmt"
)

func TestInitBitTransactionInfo(t *testing.T) {
	bt := InitBtcTransactionInfo()
	fmt.Println(bt)

	bt2 := InitBtcTransacton()
	fmt.Println(bt2)
}

func TestBtcTransactionInfo_Store2Db(t *testing.T) {
	bt := InitBtcTransactionInfo()
	bt.FromAddress = "a"
	bt.ToAddress = "b"
	bt.Remark = "remark"
	bt.Amount = "100"
	bt.Fee = "0.003"
	bt.Hash = "11"

	if nil == bt.Store2Db() {

	}else {
		t.Error("not pass")
	}

}

func TestGetRemarkInfo(t *testing.T) {
	remark := GetRemarkInfo("11")
	fmt.Println("remark:",remark)
}