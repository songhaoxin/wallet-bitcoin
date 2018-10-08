package model

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"log"
)

type Prev_out struct {
	Spent bool `json:"spent"`
	Tx_index json.Number `json:"tx_index"`
	Type json.Number `json:"type"`
	Addr string `json:"addr"`
	Value json.Number `json:"value"`
	N json.Number `json:"n"`
	Script string `json:"script"`
}

type InputsElement struct {
	Sequence json.Number `json:"sequence"`
	Witness string `json:"witness"`
	Prev_out Prev_out `json:"prev_out"`
	Script string `json:"script"`
}


type OutElement struct {
	Spent bool `json:"spent"`
	Tx_index json.Number `json:"tx_index"`
	Type json.Number `json:"type"`
	Addr string `json:"addr"`
	Value json.Number `json:"value"`
	N json.Number `json:"n"`
	Script string `json:"script"`
}

type Tx struct {
	Ver json.Number `json:"ver"`
	Inputs []InputsElement `json:"inputs"`
	Weight json.Number `json:"weight"`
	Block_height json.Number `json:"block_height"`
	Relayed_by string `json:"relayed_by"`
	Out []OutElement `json:"out"`
	Lock_time json.Number `json:"lock_time"`
	Result json.Number `json:"result"`
	Size json.Number `json:"size"`
	Time json.Number `json:"time"`
	Tx_index json.Number `json:"tx_index"`
	Vin_sz json.Number `json:"vin_sz"`
	Hash string `json:"hash"`
	Vout_sz json.Number `json:"vout_sz"`
}

type TxInfo struct {
	Tash160 string `json:"tash_160"`
	Address string `json:"address"`
	N_tx json.Number `json:"n_tx"`
	Total_received json.Number `json:"total_received"`
	Total_sent json.Number `json:"total_sent"`
	Final_balance json.Number `json:"final_balance"`
	Txs []Tx `json:"txs"`
}


func BalanceByAddress(address string,txInfo TxInfo) (balance decimal.Decimal,err error)  {

	if "" == address {
		err = fmt.Errorf("%s", "地址不能为空")
		return
	}

	balance, _ = decimal.NewFromString("0")
	base,_ := decimal.NewFromString("100000000")

	balance,err = decimal.NewFromString(txInfo.Final_balance.String())
	if nil != err {return }

	balance = balance.Div(base)

	shouldChangeValue := CalculateBalanceShouldChangedByAddress(address,txInfo.Txs)

	balance = balance.Add(shouldChangeValue)

	//balance,_ = strconv.ParseFloat(fmt.Sprintf("%.8f", balance), 64)

	fmt.Println("余额：",balance.String())

	return
}

func CalculateBalanceShouldChangedByAddress(address string,txs []Tx) (shouldChangeValue decimal.Decimal) {

	if 0 == len(txs) {
		log.Println("交易信息为空")
		return
	}

	//balanceChanged,_ = decimal.NewFromString("0")
	// 需要调整的余额数量
	shouldChangeValue,_ = decimal.NewFromString("0")

	base,_ := decimal.NewFromString("100000000")

	for _,tx  := range txs {
		if tx.Block_height != "" {
			continue
		}

		inValue,_ :=  decimal.NewFromString("0")
		outValue,_ := decimal.NewFromString("0")

		// 遍历输入的数组
		for _,inputElm := range tx.Inputs {
			if inputElm.Prev_out.Addr == address {
				per_inValue,err := decimal.NewFromString(inputElm.Prev_out.Value.String())
				if nil != err {
					fmt.Println(err)
					continue
				}
				per_inValue = per_inValue.Div(base).Round(8)
				inValue = inValue.Add(per_inValue)
			}
		}
		fmt.Println("Input:",inValue)

		// 遍历输出的数组
		for _,outElm := range tx.Out {
			if outElm.Addr == address {
				per_outValue,err := decimal.NewFromString(outElm.Value.String())
				if nil != err {
					fmt.Println(err)
					continue
				}
				per_outValue = per_outValue.Div(base).Round(8)
				outValue = outValue.Add(per_outValue)
			}
		}
		fmt.Println("Output:",outValue)


		subValue := outValue.Sub(inValue)
		// 是收到了没有确认的交易
		if subValue.GreaterThan(decimal.Zero) {
			shouldChangeValue = shouldChangeValue.Sub(subValue)
		} else {
			// 是发送了还没有确认的交易
			sendValue := inValue.Sub(outValue)
			shouldChangeValue = shouldChangeValue.Add(sendValue)
		}
	}

	log.Println("需要调整的余额",shouldChangeValue.String())

	return
}





































type TransactionModelJson struct {
	Address string `json:"address"`
} 