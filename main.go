package main

import (
	_ "./docs"
	. "wallet-bitcoin/router"
)


// @title Bitcoin Wallet API
// @version 1.0
// @description 超链 多币种钱包 -- Btc接口
// @termOfService http://swagger.io/terms/

func main() {
	router := InitRouter()

	router.Run(":8788")
}