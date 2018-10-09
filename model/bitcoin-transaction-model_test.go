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
