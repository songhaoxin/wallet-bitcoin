package model

import (
	"testing"
	"fmt"
)

func TestWallet_IsExistBtc(t *testing.T) {
	w := Wallet{}
	exist := w.IsExistBtc(370,"abc")
	if exist {
		fmt.Println("has one")
	} else {
		fmt.Println("no no")
	}
}

func TestWallet_IsExistWallet(t *testing.T) {
	w := Wallet{}
	exist := w.IsExistWallet(372)
	if exist {
		fmt.Println("has one")
	} else {
		fmt.Println("no no")
	}
}

func TestWallet_AddBtcAccount(t *testing.T) {
	w := Wallet{}
	exist := w.IsExistBtc(370,"abc")
	if exist {
		fmt.Println("has one")
		return
	} else {
		fmt.Println("no no")
	}

	state := w.AddBtcAccount(370,"abc")
	if state {
		fmt.Println("添加BTC帐户成功！")
	} else {
		fmt.Println("添加BTC帐户失败！")
	}
}

func TestCheckPhoneNumber(t *testing.T) {
	CheckPhoneNumber("1312957030")
}

func TestWallet_UpdatePhoneNumber(t *testing.T) {
	w := Wallet{}
	w.UpdatePhoneNumber("xxxx",373)
}