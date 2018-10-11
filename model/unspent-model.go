package model

import "encoding/json"

type UnspentOutputs struct {
	Unspent_outputs []UnspentTx `json:"unspent_outputs"`
}

type UnspentTx struct {
	Tx_hash string `json:"tx_hash"`
	Tx_hash_big_endian string `json:"tx_hash_big_endian"`
	Tx_index json.Number `json:"tx_index"`
	Tx_output_n json.Number `json:"tx_output_n"`
	Script string `json:"script"`
	Value json.Number `json:"value"`
	Value_hex string `json:"value_hex"`
	Confirmations json.Number `json:"confirmations"`
}