// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-10-10 15:15:54.88516188 +0800 CST m=+0.335442180

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server Petstore server.",
        "title": "Bitcoin Wallet API",
        "contact": {},
        "license": {},
        "version": "1.0"
    },
    "paths": {
        "/btc/v1/account/balance/": {
            "get": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取余额",
                "parameters": [
                    {
                        "type": "string",
                        "description": "比特币帐户地址",
                        "name": "address",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "failed",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/btc/v1/account/transactions/": {
            "get": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取交易列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "比特币帐户地址",
                        "name": "address",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "起始页",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "最多获取的记录总数",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "failed",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "发送交易",
                "parameters": [
                    {
                        "type": "string",
                        "description": "交易的签名后字符串",
                        "name": "tx",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "发送方地址",
                        "name": "from",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "接收方地址",
                        "name": "to",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "矿工费",
                        "name": "fee",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "交易额",
                        "name": "amount",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "备注",
                        "name": "remark",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "failed",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/btc/v1/wallet/accounts/": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "向钱包增加BTC帐户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "钱包的Id",
                        "name": "walletid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "BTC地址",
                        "name": "address",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "添加BTC帐户成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "非法参数",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "失败",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/btc/v1/wallet/phone/": {
            "put": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "修改钱包的电话号码",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "钱包的Id",
                        "name": "walletid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "电话号码",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "修改电话号码成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "非法参数",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "失败",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
