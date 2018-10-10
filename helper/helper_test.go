package helper

import (
	"testing"
	"wallet-bitcoin/setting"
	"log"
)

func TestRpcBtcPort(t *testing.T) {
	var parmas = []string{"0100000001b67e3229cf7b48ad7a46a1760f23a1a70d75e657e6d92db3b93d435be665f860010000006a47304402206b23eba98e334009ededc169d938670b5bfd6b29db86e995bb250958f164b37802206723942f7290f9d6c5dd3d3124dee2cc133c25f4e1211cabf2a266aac25b05ce0121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803dffffffff02204e00000000000017a914af4072741af0fc265a6a8e946ed25c3088327c9887d05a0700000000001976a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac00000000"}
	res,err := RpcBtcPort(setting.BtcPortHostDebug,"sendrawtransaction",parmas)
	log.Println(err)

	log.Println(res)


	if resMap,ok := res.(map[string]interface{});ok {
		result := resMap["result"]
		if nil == result {
			log.Println("no nono no")
		}
	}


}