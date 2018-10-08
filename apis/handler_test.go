package apis

import (
	"testing"
	"fmt"
)

func TestGetBalanceApi(t *testing.T) {
	priceMap := getPrice("CNY","BTC,ETH")
	fmt.Println(priceMap)
}

func TestAddBtcAccount(t *testing.T) {
	AddBtcAccount(nil)
	fmt.Println("abc")
}