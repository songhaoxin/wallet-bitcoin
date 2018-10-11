package helper

import (
	"log"
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func RpcBtcPort(host string,method string,params []string) (res interface{}, err error) {

	data := make(map[string]interface{})
	data["jsonrpc"] = "1.0"
	data["id"] = 1
	data["method"] = method
	data["params"] = params

	resBytes, err2 := BtcRpcPost(host, data)
	if err2 != nil {
		log.Println(err.Error())
		res,err = nil,err2
	}

	err = json.Unmarshal(resBytes,&res)

	return
}



func BtcRpcPost(url string, send map[string]interface{}) ([]byte, error) {
	bytesData, err := json.Marshal(send)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println("rpc post:", err.Error())
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	//fmt.Println("resp:", resp)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	//log.Println(string(respBytes))
	//byte数组直接转成string，优化内存
	return respBytes, nil

}

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprint("print:%v",p)

}