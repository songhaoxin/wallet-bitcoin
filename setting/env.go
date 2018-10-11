package setting

var IsDebugEnv bool = true

// 负责 发送交易的 结点地址
const BtcPortHostDebug = "http://test:test@47.75.194.231:8332"
const BtcPortHost = "http://test:test@47.75.194.231:8332"

// 负责 快速查询交易余额、交易列表信息 的服务站
const BtcApiServerHostDebug = "https://testnet.blockchain.info/"
const BtcApiServerHost = "https://blockchain.info/"

// 保存交易信息的数据库地址
//const ServerDBConnectString  = "root:current@tcp(47.75.194.231)/multi_wallet?charset=utf8&parseTime=True&loc=Local"
const ServerDBConnectString  = "root:current@tcp(47.75.194.231)/test_wallet?charset=utf8&parseTime=True&loc=Local"


/// 获取Btc结点地址
func GetBtcPortHost() string  {
	if IsDebugEnv {
		return BtcPortHostDebug
	}
	return BtcPortHost
}

func GetBtcApiHost() string {
	if IsDebugEnv {
		return BtcApiServerHostDebug
	}
	return BtcApiServerHost
}