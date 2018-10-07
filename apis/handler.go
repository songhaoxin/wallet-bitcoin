package apis

import (
	"github.com/gin-gonic/gin"
	"log"
	"encoding/json"
	"github.com/kirinlabs/HttpRequest"
	"net/http"
	"wallet-bitcoin/setting"
	"wallet-bitcoin/models"
)

// @Summary 获取余额(比特币)
// @Produce  json
// @Accept  application/x-www-form-urlencoded
// @Param   address     formData    string     true        "比特币帐户地址"
// @Success 200 {string} string    "ok"
// @Failure 400 {string} string "failed"
// @Failure 500 {string} string "failed"
// @Router /bth/balance/ [post]
func GetBalanceApi(c *gin.Context)  {
	address := c.PostForm("address")
	log.Println("查询余额的地址：",address)
	if "" == address {
		log.Println("查询地址不能为空")
		c.JSON(http.StatusBadRequest,gin.H{
			"status": 1,
			"msg":    "缺少参数：address",
		})
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
	res,err := req.Get(api, map[string]interface{}{
		//"address":address,
	})

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


	resMap := make(map[string]interface{})
	err = json.Unmarshal(body,&resMap)
	if nil != err {
		log.Println(err)
		log.Println(err)
		c.JSON(res.StatusCode(),gin.H{
			"status": 1,
			"err":err,
			"msg":    "查询失败！",
		})
		return
	}


	var txInfo models.TxInfo
	err = json.Unmarshal(body,&txInfo)
	if err != nil {
		c.JSON(res.StatusCode(),gin.H{
			"status": 1,
			"err":err,
			"msg":    "查询失败！",
		})
		return
	}

	finalBalance,err := models.BalanceByAddress(address,txInfo)
	if err != nil {
		c.JSON(res.StatusCode(),gin.H{
			"status": 1,
			"err":err,
			"msg":    "查询失败！",
		})
		return
	}


	log.Println(resMap)
	c.JSON(res.StatusCode(),gin.H{
		"status": 0,
		"msg":    "查询成功！",
		"data":finalBalance,
	})


}
