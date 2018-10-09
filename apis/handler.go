package apis

import (
	"github.com/gin-gonic/gin"
	"log"
	"encoding/json"
	"github.com/kirinlabs/HttpRequest"
	"net/http"
	"wallet-bitcoin/setting"
	"wallet-bitcoin/model"
	"github.com/shopspring/decimal"
	"strconv"
	"fmt"

)




// @Summary 向钱包增加BTC帐户
// @Produce  json
// @Accept  application/x-www-form-urlencoded
// @Param   walletid     formData    int     true        "钱包的Id"
// @Param   address 	 formData	 string		true 		"BTC地址"
// @Success 201 {string} string    "添加BTC帐户成功"
// @Failure 400 {string} string "非法参数"
// @Failure 500 {string} string "失败"
// @Router /btc/v1/wallet/accounts/ [post]
func AddBtcAccountApi(c *gin.Context)  {
	walletId := c.PostForm("walletid")
	log.Println("传入的钱包Id:",walletId)

	if "" == walletId {
		log.Println("钱包Id不能为空")
		c.JSON(http.StatusBadRequest,gin.H{
			"status": 1,
			"msg":    "缺少参数：walletid",
		})
		return
	}
	walletId_Int,err := strconv.Atoi(walletId)
	if nil != err {
		log.Println(err)
		c.JSON(http.StatusBadRequest,gin.H{
			"status": 1,
			"msg":    "非法的walletid",
		})
		return
	}

	wallet := model.Wallet{}

	if !wallet.IsExistWallet(walletId_Int) {
		c.JSON(http.StatusBadRequest,gin.H{
			"status": 1,
			"msg":    "walletid 不存在",
		})
		return
	}

	//---------------------

	address := c.PostForm("address")
	log.Println("账户的地址：",address)
	if "" == address {
		log.Println("账户的地址不能为空")
		c.JSON(http.StatusBadRequest,gin.H{
			"status": 1,
			"msg":    "缺少参数：address",
		})
		return
	}

	////////////////////////////////////////////////////////////


	if wallet.IsExistBtc(walletId_Int,address) {
		c.JSON(http.StatusBadRequest,gin.H{
			"status": 1,
			"msg":    "钱包中已经存在该地址的BTC帐户",
		})
		return
	}

	state := wallet.AddBtcAccount(walletId_Int,address)
	if !state {
		c.JSON(http.StatusInternalServerError,gin.H{
			"status": 1,
			"msg":    "增加BTC帐户失败",
		})
	}

	c.JSON(http.StatusCreated,gin.H{
		"status": 0,
		"msg":    "增加BTC帐户成功",
	})
	return

}


// @Summary 修改钱包的电话号码
// @Produce  json
// @Accept  application/x-www-form-urlencoded
// @Param   walletid     formData    int     true        "钱包的Id"
// @Param   phone 	 formData	 string		true 		"电话号码"
// @Success 200 {string} string    "修改电话号码成功"
// @Failure 400 {string} string "非法参数"
// @Failure 500 {string} string "失败"
// @Router /btc/v1/wallet/phone/ [put]
func UpdatePhoneNumberApi(c *gin.Context)  {
	walletId := c.PostForm("walletid")
	log.Println("传入的钱包Id:",walletId)

	if "" == walletId {
		log.Println("钱包Id不能为空")
		c.JSON(http.StatusBadRequest,gin.H{
			"status": 1,
			"msg":    "缺少参数：walletid",
		})
		return
	}

	phoneNumber := c.PostForm("phone")
	log.Println("传入的电话号码：",phoneNumber)

	if !model.CheckPhoneNumber(phoneNumber) {
		log.Println("非法的电话号码")
		c.JSON(http.StatusBadRequest,gin.H{
			"status": 1,
			"msg":    "非法的电话号码",
		})
		return
	}

	walletId_Int,err := strconv.Atoi(walletId)
	if nil != err {
		log.Println(err)
		c.JSON(http.StatusBadRequest,gin.H{
			"status": 1,
			"msg":    "非法的walletid",
		})
		return
	}

	wallet := model.Wallet{}

	if !wallet.UpdatePhoneNumber(phoneNumber,walletId_Int) {
		c.JSON(http.StatusInternalServerError,gin.H{
			"status": 1,
			"msg":    "修改电话号码失败",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"status": 0,
		"msg":    "修改电话号码成功",
	})

}


// @Summary 获取余额
// @Produce  json
// @Accept  application/x-www-form-urlencoded
// @Param   address     query    string     true        "比特币帐户地址"
// @Success 200 {string} string    "ok"
// @Failure 400 {string} string "failed"
// @Failure 500 {string} string "failed"
// @Router /btc/v1/account/balance/ [get]
func GetBtcBalanceApi(c *gin.Context)  {

	//address := c.PostForm("address")
	address := c.Request.URL.Query().Get("address")
	log.Println("查询余额的地址：",address)
	if "" == address {
		log.Println("查询地址不能为空")
		c.JSON(http.StatusBadRequest,gin.H{
			"status": 1,
			"msg":    "缺少参数：address",
		})
		return
	}


	var api string
	if setting.IsDebugEnv {
		api = setting.BthApiServerHostDebug + "rawaddr/"
	} else {
		api = setting.BthApiServerHost + "rawaddr/"
	}
	api = api + address

	req := HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded", //这也是HttpRequest包的默认设置
	})
	res,err := req.Get(api, map[string]interface{}{})

	if nil != err {
		log.Println(err)
		var code int
		var errStr string = err.Error()
		if nil == res {
			code = 500
			errStr = "无法连接到服务"
		} else {
			log.Println(res.StatusCode())
			code = res.StatusCode()
		}

		c.JSON(code,gin.H{
			"status": 1,
			"err":errStr,
			"msg":    "查询失败！",
		})
		return
	}

	log.Println(res.StatusCode())

	body,err := res.Body()
	if nil != err {
		log.Println(err)
		c.JSON(res.StatusCode(),gin.H{
			"status": 1,
			"err":err,
			"msg":    "查询失败！",
		})
		return
	}
	log.Println(body)


	var txInfo model.TxInfo
	err = json.Unmarshal(body,&txInfo)
	if err != nil {
		c.JSON(res.StatusCode(),gin.H{
			"status": 1,
			"err":err,
			"msg":    "查询失败！",
		})
		return
	}

	finalBalance,err := model.BalanceByAddress(address,txInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"status": 1,
			"err":err,
			"msg":    "查询失败！",
		})
		return
	}

	log.Println("余额：",finalBalance.String())



	btcBalance := model.BitcoinBalance{}
	btcBalance.Num = 0.0
	btcBalance.Price_USD = 0.0
	btcBalance.Total_USD = 0.0
	btcBalance.Price_CNY = 0.0
	btcBalance.Total_CNY = 0.0

	finalBalanceFloat,_ := finalBalance.Round(8).Float64()
	ammount,err := strconv.ParseFloat(fmt.Sprintf("%.8f", finalBalanceFloat), 64)
	if nil == err {
		btcBalance.Num = ammount
	}


	// 获取价格信息
	priceMap := getPrice("USD","BTC")
	// 设置USD价格信息
	usdPrice,exists := priceMap["USD"]
	if exists {
		if "" != usdPrice {
			// 把字符串转换成decimal
			usdPriceDecimal,err := decimal.NewFromString(usdPrice)
			if nil == err {
				usdPriceFloatTmp,_ := usdPriceDecimal.Round(2).Float64()
				usdPriceFloat,err := strconv.ParseFloat(fmt.Sprintf("%.8f", usdPriceFloatTmp), 64)
				if nil == err {
					btcBalance.Price_USD = usdPriceFloat
				}

				total_USDDecimal := usdPriceDecimal.Mul(finalBalance)
				total_USD_FloatTmp,_ := total_USDDecimal.Round(2).Float64()
				total_USD_Float, err := strconv.ParseFloat(fmt.Sprintf("%.8f", total_USD_FloatTmp), 64)
				if nil == err {
					btcBalance.Total_USD = total_USD_Float
				}
			}
		}
	}

	// 设置CNY价格信息
	priceMap = getPrice("CNY","BTC")
	cnyPrice,exists := priceMap["CNY"]
	if exists {
		if "" != usdPrice {
			// 把字符串转换成decimal
			cnyPriceDecimal,err := decimal.NewFromString(cnyPrice)
			if nil == err {
				cnyPriceFloatTmp,_ := cnyPriceDecimal.Round(2).Float64()
				cnyPriceFloat,err := strconv.ParseFloat(fmt.Sprintf("%.8f", cnyPriceFloatTmp), 64)
				if nil == err {
					btcBalance.Price_CNY = cnyPriceFloat
				}

				total_CNYDecimal := cnyPriceDecimal.Mul(finalBalance)
				total_CNY_FloatTmp,_  := total_CNYDecimal.Round(2).Float64()
				total_CNY_Float,err := strconv.ParseFloat(fmt.Sprintf("%.8f", total_CNY_FloatTmp), 64)
				if nil == err {
					btcBalance.Total_CNY = total_CNY_Float
				}
			}
		}
	}


	c.JSON(res.StatusCode(),gin.H{
		"status": 0,
		"msg":    "查询成功！",
		"data":btcBalance,
	})


}

func getPrice(currency string,symbols string) map[string]string {
	var priceMap  = make(map[string]string)
	if "" == currency || "" == symbols {
		log.Println("参数不能为空")
		return priceMap
	}

	req := HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded", //这也是HttpRequest包的默认设置
	})
	res,err := req.Get("https://api.trustwalletapp.com/prices?currency=" + currency + "&symbols=" + symbols, map[string]interface{}{})
	if nil != err {
		log.Println(err)
		return priceMap
	}
	body,err := res.Body()
	if nil != err {
		return priceMap
	}
	var resMap map[string]interface{}
	err = json.Unmarshal(body,&resMap)
	if nil == err {
		if response,ok := resMap["response"].([]interface{});ok {
			for _,perContent := range response {
				if content,ok := perContent.(map[string]interface{});ok {
					price := content["price"].(string)
					priceMap[currency] = price
				}
			}
		}
	}

	return priceMap
}

// @Summary 获取交易列表
// @Produce  json
// @Accept  application/x-www-form-urlencoded
// @Param   address     query    string     true        "比特币帐户地址"
// @Param   offset      query    string     false        "起始页"
// @Param   limit      	query    string     false        "最多获取的记录总数"
// @Success 200 {string} string    "ok"
// @Failure 400 {string} string "failed"
// @Failure 500 {string} string "failed"
// @Router /btc/v1/account/transactions/ [get]
func GetBtcTransactionsApi(c *gin.Context)  {

	address := c.Request.URL.Query().Get("address")
	log.Println("输入地址：",address)
	if "" == address {
		log.Println("查询地址不能为空")
		c.JSON(http.StatusBadRequest,gin.H{
			"status": 1,
			"msg":    "缺少参数：address",
		})
		return
	}

	offset := c.Request.URL.Query().Get("offset")
	limit := c.Request.URL.Query().Get("limit")


	var api string
	if setting.IsDebugEnv {
		api = setting.BthApiServerHostDebug + "rawaddr/"
	} else {
		api = setting.BthApiServerHost + "rawaddr/"
	}
	api = api + address
	if "" != offset {
		api = api + "?offset=" + offset
	}
	if "" != limit {
		if "" == offset {
			api = api + "?limit=" + limit
		} else {
			api = api + "&limit=" + limit
		}
	}
	log.Println(api)

	req := HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded", //这也是HttpRequest包的默认设置
	})
	res,err := req.Get(api, map[string]interface{}{})

	if nil != err {
		log.Println(err)
		var code int
		var errStr string = err.Error()
		if nil == res {
			code = 500
			errStr = "无法连接到服务"
		} else {
			log.Println(res.StatusCode())
			code = res.StatusCode()
		}

		c.JSON(code,gin.H{
			"status": 1,
			"err":errStr,
			"msg":    "查询失败！",
		})
		return
	}

	log.Println(res.StatusCode())

	body,err := res.Body()
	if nil != err {
		log.Println(err)
		c.JSON(res.StatusCode(),gin.H{
			"status": 1,
			"err":err,
			"msg":    "查询失败！",
		})
		return
	}
	//log.Println(body)


	var txInfo model.TxInfo
	err = json.Unmarshal(body,&txInfo)
	if err != nil {
		c.JSON(res.StatusCode(),gin.H{
			"status": 1,
			"err":err,
			"msg":    "查询失败！",
		})
		return
	}

	// 获取具体的交易信息列表
	trs,err := model.GetBtcTransaction(address,txInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"status": 1,
			"err":"服务错误",
			"msg":    "查询失败！",
		})
		return
	}


	c.JSON(http.StatusOK,gin.H{
		"status": 0,
		"data":trs,
		"msg":    "查询成功！",
	})


}


