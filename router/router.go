package router

import (

	"github.com/gin-gonic/gin"
	. "wallet-bitcoin/apis"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/bth/balance/",GetBalanceApi)
	return router
}
