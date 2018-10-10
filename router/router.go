package router

import (

	"github.com/gin-gonic/gin"
	. "wallet-bitcoin/apis"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	/// Swagger UI 相关API
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	/// 获取BTC帐户的余额
	router.GET("/btc/v1/account/balance/", GetBtcBalanceApi)

	/// 获取交易列表
	router.GET("/btc/v1/account/transactions/",GetBtcTransactionsApi)

	/// 发送交易
	router.POST("/btc/v1/account/transactions/",SendTransactionApi)

	/// 向指定的钱包添加BTC帐户
	router.POST("/btc/v1/wallet/accounts/",AddBtcAccountApi)

	/// 修改钱包的电话号码
	router.PUT("/btc/v1/wallet/phone/",UpdatePhoneNumberApi)

	return router
}
