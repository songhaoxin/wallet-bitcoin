package main

import (
	_ "./docs"
	. "wallet-bitcoin/router"
)


// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termOfService http://swagger.io/terms/

func main() {


	router := InitRouter()
	router.Run(":8788")
}