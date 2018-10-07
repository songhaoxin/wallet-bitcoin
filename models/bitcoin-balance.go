package models

/*
"bgft_balance": 53.4,
   "num": "0.113918",
    "price_CNY": "1544.36",
    "price_CNY_total": 175.93,
    "price_USD": "224.71",
    "price_USD_total": 25.6,
    "url": "http://47.75.194.231:8781/images/tokenicon/eth.png"
 */

 type BitcoinBalance struct {
 	num float64
 	price_CNY float64
 	price_USD float64
 	total_CNY float64
 	total_USD float64
 	url string
 }
