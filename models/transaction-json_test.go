package models

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestBalanceByAddress(t *testing.T) {
	jsonTxs := `{
    "hash160": "fd2fdebf6a9662e8e85673569251a9997d737f0e",
    "address": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
    "n_tx": 17,
    "total_received": 7100000,
    "total_sent": 2360000,
    "final_balance": 4740000,
    "txs": [
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285836793,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 4780000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "473044022068a534aabdefc78414f8c01669f64439bfab380a230ce56664d425c51b648a9e022067e643e1529af10f17e732d08633893f1dc329fe56f7bbfc3ec82ad6281de7370121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285849324,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 10000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": false,
                    "tx_index": 285849324,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 4740000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": 0,
            "size": 223,
            "time": 1538884195,
            "tx_index": 285849324,
            "vin_sz": 1,
            "hash": "7dc60155c0f24339e313705f5b29ead553e037840e6d159045aacd688440e3c8",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285836652,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 4820000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "473044022022e5b1b81c4e0c4aae91e7ce9bf47530ccb6db3b7c9a54669151f3265324f9ae02203d840623ffd7ed788c9d5e505663f92521d2bab8ecfacdabc5f1ec3aef9e24130121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285836793,
                    "type": 0,
                    "addr": "2N9DsSiVVYt6Cmb5xex7kfxJo5oLxFE6Nmn",
                    "value": 10000,
                    "n": 0,
                    "script": "a914af4072741af0fc265a6a8e946ed25c3088327c9887"
                },
                {
                    "spent": true,
                    "tx_index": 285836793,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 4780000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 223,
            "time": 1538751428,
            "tx_index": 285836793,
            "vin_sz": 1,
            "hash": "595d11a29e14951b426a93e61ab216f8ee4c4486c0f48c69ce1b3fd914c0dd39",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285835881,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 4860000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "4730440220330b4be456f216c632964c52e06471eeb9615be06de83cece91b8d0a3294694402201a1c157b8ad563a3acea955dd62f7523b597a0919f9702d212d4899bfd2a2e430121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1415078,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285836652,
                    "type": 0,
                    "addr": "2N9DsSiVVYt6Cmb5xex7kfxJo5oLxFE6Nmn",
                    "value": 10000,
                    "n": 0,
                    "script": "a914af4072741af0fc265a6a8e946ed25c3088327c9887"
                },
                {
                    "spent": true,
                    "tx_index": 285836652,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 4820000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 223,
            "time": 1538750608,
            "tx_index": 285836652,
            "vin_sz": 1,
            "hash": "f1ad3a126d13b9e98260c77b3aa4c106d8d7f3c213dc6725affac5d2ebacd31d",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285835503,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 4900000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "47304402200a3247c401928638acefa0ae22f8867bea70ed8fbb96dadb2c912116a0cade81022032fdc5f2bbd6aa02da0c46786274ee04e7641aa67b1adc66e570d375f6ca59660121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1415074,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285835881,
                    "type": 0,
                    "addr": "2N9DsSiVVYt6Cmb5xex7kfxJo5oLxFE6Nmn",
                    "value": 10000,
                    "n": 0,
                    "script": "a914af4072741af0fc265a6a8e946ed25c3088327c9887"
                },
                {
                    "spent": true,
                    "tx_index": 285835881,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 4860000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 223,
            "time": 1538746156,
            "tx_index": 285835881,
            "vin_sz": 1,
            "hash": "1821a895dfa45ff8c5b6fbba8be91226daed93165b18a8995d5eed1d196b66f9",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285835098,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 4940000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "473044022004704a1d615a114d061a14dc49ec3fb2788e5ab244d4063bb96a9edb39afdfb902202aa6981ce06d9502e12b9eb053245146fd6e68a5625d097e56fe3e8632a4b8b80121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1415071,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285835503,
                    "type": 0,
                    "addr": "2N9DsSiVVYt6Cmb5xex7kfxJo5oLxFE6Nmn",
                    "value": 10000,
                    "n": 0,
                    "script": "a914af4072741af0fc265a6a8e946ed25c3088327c9887"
                },
                {
                    "spent": true,
                    "tx_index": 285835503,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 4900000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 223,
            "time": 1538743677,
            "tx_index": 285835503,
            "vin_sz": 1,
            "hash": "aec213e2538c389b5d3b037bd5dcb5bdfe2f3c9d18d3a606bb80c602cb157ac7",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285834438,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 4980000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "47304402206401e369558f76b88f1e2843e781313fecce73655d7dc221e86b743f74d233fd02200301de6353008e2014dcf2c9e6dc53ea6041b3130c9c3facf32d3ce8e68ffc510121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1415069,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285835098,
                    "type": 0,
                    "addr": "2N9DsSiVVYt6Cmb5xex7kfxJo5oLxFE6Nmn",
                    "value": 10000,
                    "n": 0,
                    "script": "a914af4072741af0fc265a6a8e946ed25c3088327c9887"
                },
                {
                    "spent": true,
                    "tx_index": 285835098,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 4940000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 223,
            "time": 1538741730,
            "tx_index": 285835098,
            "vin_sz": 1,
            "hash": "d5db3c7edd63464a32f7ff4d3662d3744126ecf2680e4780a9cf669cd6a9256d",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285834426,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5020000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "4830450221008d7368b36d6a4b1fe3da0d146d9f0a8ae0b633132bf21eec3fecf9b2f61b104c02203973aba097f9e8b83971bbf684a0498248b52a73718cff6df96f4baacfd0a0f00121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 804,
            "block_height": 1415066,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": true,
                    "tx_index": 285834438,
                    "type": 0,
                    "value": 10000,
                    "n": 0,
                    "script": ""
                },
                {
                    "spent": true,
                    "tx_index": 285834438,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 4980000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 201,
            "time": 1538738142,
            "tx_index": 285834438,
            "vin_sz": 1,
            "hash": "bd009194646a1aad33fa5dc8ee285b06e3fccc5a0bc834f2b420d01e850a562a",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285828339,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5060000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "473044022003529012d55cde8367458269fa74a767394047c02c0e73de23f45cc92f534e60022063c5da565124f47f5263c9e25852e84d3b028cd05fb63b8788eb965187c5f88d0121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1415066,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285834426,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 10000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": true,
                    "tx_index": 285834426,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5020000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 223,
            "time": 1538738047,
            "tx_index": 285834426,
            "vin_sz": 1,
            "hash": "3ec7d66f856e84d3390bd75f2d30104d6048fdaa2e3d80df2fe3abd904bfefc0",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285784788,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5190000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "47304402204dfda54165bfd7ebce06e5ba8f89fb70e4b386afad04b4ea6817e6caa570242902205d10a632c50db6d140f2743b2876a8107c277fbf62df268ca2f861521e315a8c0121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1415046,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285828339,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 100000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": true,
                    "tx_index": 285828339,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5060000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 223,
            "time": 1538713884,
            "tx_index": 285828339,
            "vin_sz": 1,
            "hash": "1703f6df4eb66ead73b93603f1e79a1e1f4018ec10291cfc8ea8f78685c95a74",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285728418,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5320000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "483045022100b1ad41e4f5ee35b008b087b43bcf35c2733c080bcf6dbec75e640581198c0276022008793c0aeabbe7a2b4ac686304882e052b025df8d3a973f34ceb358b777cbc260121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 904,
            "block_height": 1414856,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285784788,
                    "type": 0,
                    "addr": "mqDn1EhCupHYtNVhwiUgXenCGLPTGsPH4b",
                    "value": 100000,
                    "n": 0,
                    "script": "76a9146a7198cc6fc9230c5bb8ae6e4e59dfd5682f7adb88ac"
                },
                {
                    "spent": true,
                    "tx_index": 285784788,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5190000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -130000,
            "size": 226,
            "time": 1538495841,
            "tx_index": 285784788,
            "vin_sz": 1,
            "hash": "079af63c40f226b4f0a41d44068f16f19713be1496ba5682a02faa5821d6b338",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285728286,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5450000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "47304402204fc6751d992c56f72e1cf542342261afde63204f895c894261dffdd26923a899022035a8494a63054eb0e530604b70e6e62f195ef40c56064c92c487ba60ff2583bb0121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1414616,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285728418,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 100000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": true,
                    "tx_index": 285728418,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5320000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -130000,
            "size": 223,
            "time": 1538211779,
            "tx_index": 285728418,
            "vin_sz": 1,
            "hash": "e8dc71f4460d1825c4d02a6474a3bec4fe893f5b75279cf442f4f4b199a66475",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285724842,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5580000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "4830450221009d40334d1333eaa0c25632103cf4258a9b28d6abaf2e08d8483a99863387345c02206649cfe243c0f3e60f204c5b5b8b59f380ac0633d7d19fc38a18b1b414bf5cb20121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 896,
            "block_height": 1414613,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285728286,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 100000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": true,
                    "tx_index": 285728286,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5450000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -130000,
            "size": 224,
            "time": 1538211248,
            "tx_index": 285728286,
            "vin_sz": 1,
            "hash": "f54630e8be2a799cac4b05a78941639750a1db1480e12726185fea64431a5469",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285724065,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5710000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "483045022100c67dac814f50d4d76c703fa728436e1dfa011f00c8e4d8814c0686f84034e0060220726dfb6c23c4cdcfa0813e688640446458bd640bdb7c0eb892ccff2cea7bbcf30121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 896,
            "block_height": 1414605,
            "relayed_by": "0.0.0.0",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285724842,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 100000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": true,
                    "tx_index": 285724842,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5580000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -130000,
            "size": 224,
            "time": 1538201774,
            "tx_index": 285724842,
            "vin_sz": 1,
            "hash": "fdcb30ad3c9049243736b867ec813edd700a117731cdf5259af92f7b67386466",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285723120,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5840000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "47304402205d0f1f931a733a4534ee00a9f14d26d3c8f70a16bfd4f9ee45a21efb4f2fc36202201d93bc49f265df35a4f510ccca6bc8ab16d5d5aafd2c3f03ca1f5f56f787e26a0121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1414603,
            "relayed_by": "0.0.0.0",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285724065,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 100000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": true,
                    "tx_index": 285724065,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5710000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -130000,
            "size": 223,
            "time": 1538199344,
            "tx_index": 285724065,
            "vin_sz": 1,
            "hash": "42c8f69710cf481a590a932e4233a7f16f22d52c087c2068a4d10c4e20108af3",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285722412,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5970000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "473044022061cfa988bfac85301dc289fb7e83919fc1c56d45f1bc837263f599c04cbf30a8022054254a1d00943110fac8621dbabc539bdf9ff8c9d0efffc4efcc45666435ebfc0121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1414599,
            "relayed_by": "0.0.0.0",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285723120,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 100000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": true,
                    "tx_index": 285723120,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5840000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -130000,
            "size": 223,
            "time": 1538194873,
            "tx_index": 285723120,
            "vin_sz": 1,
            "hash": "993762cb64cbc42f13a3736b78e4194d5741a7faaae994300a26bba481765f66",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285703527,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 7100000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "483045022100a4515e4c9395d9f36316a5a2db58a1b1f59b5309cbdec7ca352393e8f8e2add602207032c4b60309c546804245c28657e04f7c99cb8f5e4d46f1556d57ac94436b530121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 896,
            "block_height": 1414597,
            "relayed_by": "0.0.0.0",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285722412,
                    "type": 0,
                    "addr": "2N9DsSiVVYt6Cmb5xex7kfxJo5oLxFE6Nmn",
                    "value": 1100000,
                    "n": 0,
                    "script": "a914af4072741af0fc265a6a8e946ed25c3088327c9887"
                },
                {
                    "spent": true,
                    "tx_index": 285722412,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5970000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -130000,
            "size": 224,
            "time": 1538192437,
            "tx_index": 285722412,
            "vin_sz": 1,
            "hash": "f841e495e6cd3436956e5f31fd3189d108890cccfbf873357fabff2beead3831",
            "vout_sz": 2
        },
        {
            "ver": 2,
            "inputs": [
                {
                    "sequence": 4294967294,
                    "witness": "0247304402203bc34f956870f92f31cde6c01fe6fc3b0ed5d738560d272ad68a917bc93ccf7002200b2c972a41a49733aeb0cc4d25a724feb86732530c0d094bb1332baee0e8b7270121033a11b6e3e7629fcdf8de902a2bd2f8a2c840dffb64e62e861d042547758f9c27",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285703113,
                        "type": 0,
                        "addr": "2N8MALSUksHK8RqAujyPYNBUzxduYfR7Vpg",
                        "value": 26203758292,
                        "n": 0,
                        "script": "a914a5a955662fc14526c3185cb7c18a7a2939bb353587"
                    },
                    "script": "160014afb215093568b8c7e07dcc1241b00f2ad9ad10c1"
                }
            ],
            "weight": 669,
            "block_height": 1414548,
            "relayed_by": "0.0.0.0",
            "out": [
                {
                    "spent": true,
                    "tx_index": 285703527,
                    "type": 0,
                    "addr": "2NF4gm8h4tDxqPsoFnLtRm4ZAJocnVcNbrp",
                    "value": 26196657391,
                    "n": 0,
                    "script": "a914ef548a75dc51f9e11cbb9f59d92905e3ed5a1dc387"
                },
                {
                    "spent": true,
                    "tx_index": 285703527,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 7100000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 1414547,
            "result": -1130000,
            "size": 249,
            "time": 1538133394,
            "tx_index": 285703527,
            "vin_sz": 1,
            "hash": "01c6bc398ec9467654b64681addc89f59cfa71c440c83bc6f2cc9be95d5252ee",
            "vout_sz": 2
        }
    ]
}`

	var txInfo TxInfo

	err := json.Unmarshal([]byte(jsonTxs),&txInfo)
	if err != nil {
		fmt.Println("err was %v",err)
	}
	//fmt.Println(txs)


	//address := "mum6KQEwAHGDRr7HRXx7hufcsyqiGPUCtN"
	address := "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB"
	balance,err := BalanceByAddress(address,txInfo)
	if nil != err {
		fmt.Println(err)
		return
	}

	fmt.Println(balance)
}

func TestCalculateBalanceShouldChangedByAddress(t *testing.T) {

	jsonTxs := `{"txs": [
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285836793,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 4780000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "473044022068a534aabdefc78414f8c01669f64439bfab380a230ce56664d425c51b648a9e022067e643e1529af10f17e732d08633893f1dc329fe56f7bbfc3ec82ad6281de7370121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285849324,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 10000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": false,
                    "tx_index": 285849324,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 4740000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": 0,
            "size": 223,
            "time": 1538884195,
            "tx_index": 285849324,
            "vin_sz": 1,
            "hash": "7dc60155c0f24339e313705f5b29ead553e037840e6d159045aacd688440e3c8",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285836652,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 4820000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "473044022022e5b1b81c4e0c4aae91e7ce9bf47530ccb6db3b7c9a54669151f3265324f9ae02203d840623ffd7ed788c9d5e505663f92521d2bab8ecfacdabc5f1ec3aef9e24130121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1415079,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285836793,
                    "type": 0,
                    "addr": "2N9DsSiVVYt6Cmb5xex7kfxJo5oLxFE6Nmn",
                    "value": 10000,
                    "n": 0,
                    "script": "a914af4072741af0fc265a6a8e946ed25c3088327c9887"
                },
                {
                    "spent": true,
                    "tx_index": 285836793,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 4780000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 223,
            "time": 1538751428,
            "tx_index": 285836793,
            "vin_sz": 1,
            "hash": "595d11a29e14951b426a93e61ab216f8ee4c4486c0f48c69ce1b3fd914c0dd39",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285835881,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 4860000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "4730440220330b4be456f216c632964c52e06471eeb9615be06de83cece91b8d0a3294694402201a1c157b8ad563a3acea955dd62f7523b597a0919f9702d212d4899bfd2a2e430121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1415078,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285836652,
                    "type": 0,
                    "addr": "2N9DsSiVVYt6Cmb5xex7kfxJo5oLxFE6Nmn",
                    "value": 10000,
                    "n": 0,
                    "script": "a914af4072741af0fc265a6a8e946ed25c3088327c9887"
                },
                {
                    "spent": true,
                    "tx_index": 285836652,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 4820000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 223,
            "time": 1538750608,
            "tx_index": 285836652,
            "vin_sz": 1,
            "hash": "f1ad3a126d13b9e98260c77b3aa4c106d8d7f3c213dc6725affac5d2ebacd31d",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285835503,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 4900000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "47304402200a3247c401928638acefa0ae22f8867bea70ed8fbb96dadb2c912116a0cade81022032fdc5f2bbd6aa02da0c46786274ee04e7641aa67b1adc66e570d375f6ca59660121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1415074,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285835881,
                    "type": 0,
                    "addr": "2N9DsSiVVYt6Cmb5xex7kfxJo5oLxFE6Nmn",
                    "value": 10000,
                    "n": 0,
                    "script": "a914af4072741af0fc265a6a8e946ed25c3088327c9887"
                },
                {
                    "spent": true,
                    "tx_index": 285835881,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 4860000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 223,
            "time": 1538746156,
            "tx_index": 285835881,
            "vin_sz": 1,
            "hash": "1821a895dfa45ff8c5b6fbba8be91226daed93165b18a8995d5eed1d196b66f9",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285835098,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 4940000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "473044022004704a1d615a114d061a14dc49ec3fb2788e5ab244d4063bb96a9edb39afdfb902202aa6981ce06d9502e12b9eb053245146fd6e68a5625d097e56fe3e8632a4b8b80121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1415071,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285835503,
                    "type": 0,
                    "addr": "2N9DsSiVVYt6Cmb5xex7kfxJo5oLxFE6Nmn",
                    "value": 10000,
                    "n": 0,
                    "script": "a914af4072741af0fc265a6a8e946ed25c3088327c9887"
                },
                {
                    "spent": true,
                    "tx_index": 285835503,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 4900000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 223,
            "time": 1538743677,
            "tx_index": 285835503,
            "vin_sz": 1,
            "hash": "aec213e2538c389b5d3b037bd5dcb5bdfe2f3c9d18d3a606bb80c602cb157ac7",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285834438,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 4980000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "47304402206401e369558f76b88f1e2843e781313fecce73655d7dc221e86b743f74d233fd02200301de6353008e2014dcf2c9e6dc53ea6041b3130c9c3facf32d3ce8e68ffc510121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1415069,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285835098,
                    "type": 0,
                    "addr": "2N9DsSiVVYt6Cmb5xex7kfxJo5oLxFE6Nmn",
                    "value": 10000,
                    "n": 0,
                    "script": "a914af4072741af0fc265a6a8e946ed25c3088327c9887"
                },
                {
                    "spent": true,
                    "tx_index": 285835098,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 4940000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 223,
            "time": 1538741730,
            "tx_index": 285835098,
            "vin_sz": 1,
            "hash": "d5db3c7edd63464a32f7ff4d3662d3744126ecf2680e4780a9cf669cd6a9256d",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285834426,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5020000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "4830450221008d7368b36d6a4b1fe3da0d146d9f0a8ae0b633132bf21eec3fecf9b2f61b104c02203973aba097f9e8b83971bbf684a0498248b52a73718cff6df96f4baacfd0a0f00121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 804,
            "block_height": 1415066,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": true,
                    "tx_index": 285834438,
                    "type": 0,
                    "value": 10000,
                    "n": 0,
                    "script": ""
                },
                {
                    "spent": true,
                    "tx_index": 285834438,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 4980000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 201,
            "time": 1538738142,
            "tx_index": 285834438,
            "vin_sz": 1,
            "hash": "bd009194646a1aad33fa5dc8ee285b06e3fccc5a0bc834f2b420d01e850a562a",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285828339,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5060000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "473044022003529012d55cde8367458269fa74a767394047c02c0e73de23f45cc92f534e60022063c5da565124f47f5263c9e25852e84d3b028cd05fb63b8788eb965187c5f88d0121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1415066,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285834426,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 10000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": true,
                    "tx_index": 285834426,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5020000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 223,
            "time": 1538738047,
            "tx_index": 285834426,
            "vin_sz": 1,
            "hash": "3ec7d66f856e84d3390bd75f2d30104d6048fdaa2e3d80df2fe3abd904bfefc0",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285784788,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5190000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "47304402204dfda54165bfd7ebce06e5ba8f89fb70e4b386afad04b4ea6817e6caa570242902205d10a632c50db6d140f2743b2876a8107c277fbf62df268ca2f861521e315a8c0121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1415046,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285828339,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 100000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": true,
                    "tx_index": 285828339,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5060000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -40000,
            "size": 223,
            "time": 1538713884,
            "tx_index": 285828339,
            "vin_sz": 1,
            "hash": "1703f6df4eb66ead73b93603f1e79a1e1f4018ec10291cfc8ea8f78685c95a74",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285728418,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5320000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "483045022100b1ad41e4f5ee35b008b087b43bcf35c2733c080bcf6dbec75e640581198c0276022008793c0aeabbe7a2b4ac686304882e052b025df8d3a973f34ceb358b777cbc260121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 904,
            "block_height": 1414856,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285784788,
                    "type": 0,
                    "addr": "mqDn1EhCupHYtNVhwiUgXenCGLPTGsPH4b",
                    "value": 100000,
                    "n": 0,
                    "script": "76a9146a7198cc6fc9230c5bb8ae6e4e59dfd5682f7adb88ac"
                },
                {
                    "spent": true,
                    "tx_index": 285784788,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5190000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -130000,
            "size": 226,
            "time": 1538495841,
            "tx_index": 285784788,
            "vin_sz": 1,
            "hash": "079af63c40f226b4f0a41d44068f16f19713be1496ba5682a02faa5821d6b338",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285728286,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5450000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "47304402204fc6751d992c56f72e1cf542342261afde63204f895c894261dffdd26923a899022035a8494a63054eb0e530604b70e6e62f195ef40c56064c92c487ba60ff2583bb0121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1414616,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285728418,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 100000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": true,
                    "tx_index": 285728418,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5320000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -130000,
            "size": 223,
            "time": 1538211779,
            "tx_index": 285728418,
            "vin_sz": 1,
            "hash": "e8dc71f4460d1825c4d02a6474a3bec4fe893f5b75279cf442f4f4b199a66475",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285724842,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5580000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "4830450221009d40334d1333eaa0c25632103cf4258a9b28d6abaf2e08d8483a99863387345c02206649cfe243c0f3e60f204c5b5b8b59f380ac0633d7d19fc38a18b1b414bf5cb20121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 896,
            "block_height": 1414613,
            "relayed_by": "127.0.0.1",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285728286,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 100000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": true,
                    "tx_index": 285728286,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5450000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -130000,
            "size": 224,
            "time": 1538211248,
            "tx_index": 285728286,
            "vin_sz": 1,
            "hash": "f54630e8be2a799cac4b05a78941639750a1db1480e12726185fea64431a5469",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285724065,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5710000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "483045022100c67dac814f50d4d76c703fa728436e1dfa011f00c8e4d8814c0686f84034e0060220726dfb6c23c4cdcfa0813e688640446458bd640bdb7c0eb892ccff2cea7bbcf30121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 896,
            "block_height": 1414605,
            "relayed_by": "0.0.0.0",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285724842,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 100000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": true,
                    "tx_index": 285724842,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5580000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -130000,
            "size": 224,
            "time": 1538201774,
            "tx_index": 285724842,
            "vin_sz": 1,
            "hash": "fdcb30ad3c9049243736b867ec813edd700a117731cdf5259af92f7b67386466",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285723120,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5840000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "47304402205d0f1f931a733a4534ee00a9f14d26d3c8f70a16bfd4f9ee45a21efb4f2fc36202201d93bc49f265df35a4f510ccca6bc8ab16d5d5aafd2c3f03ca1f5f56f787e26a0121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1414603,
            "relayed_by": "0.0.0.0",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285724065,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 100000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": true,
                    "tx_index": 285724065,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5710000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -130000,
            "size": 223,
            "time": 1538199344,
            "tx_index": 285724065,
            "vin_sz": 1,
            "hash": "42c8f69710cf481a590a932e4233a7f16f22d52c087c2068a4d10c4e20108af3",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285722412,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 5970000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "473044022061cfa988bfac85301dc289fb7e83919fc1c56d45f1bc837263f599c04cbf30a8022054254a1d00943110fac8621dbabc539bdf9ff8c9d0efffc4efcc45666435ebfc0121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 892,
            "block_height": 1414599,
            "relayed_by": "0.0.0.0",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285723120,
                    "type": 0,
                    "addr": "2MzYvk4iYKgHEuh1XFXLGeyeShgcMwJ2dCJ",
                    "value": 100000,
                    "n": 0,
                    "script": "a9145021f8e6801e5499431e16d1e189e150fa553ec987"
                },
                {
                    "spent": true,
                    "tx_index": 285723120,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5840000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -130000,
            "size": 223,
            "time": 1538194873,
            "tx_index": 285723120,
            "vin_sz": 1,
            "hash": "993762cb64cbc42f13a3736b78e4194d5741a7faaae994300a26bba481765f66",
            "vout_sz": 2
        },
        {
            "ver": 1,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285703527,
                        "type": 0,
                        "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                        "value": 7100000,
                        "n": 1,
                        "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                    },
                    "script": "483045022100a4515e4c9395d9f36316a5a2db58a1b1f59b5309cbdec7ca352393e8f8e2add602207032c4b60309c546804245c28657e04f7c99cb8f5e4d46f1556d57ac94436b530121030c7c21083b0fe357a3380ba8c1c9b315f27f035a22103d98a98fa217e052803d"
                }
            ],
            "weight": 896,
            "block_height": 1414597,
            "relayed_by": "0.0.0.0",
            "out": [
                {
                    "spent": false,
                    "tx_index": 285722412,
                    "type": 0,
                    "addr": "2N9DsSiVVYt6Cmb5xex7kfxJo5oLxFE6Nmn",
                    "value": 1100000,
                    "n": 0,
                    "script": "a914af4072741af0fc265a6a8e946ed25c3088327c9887"
                },
                {
                    "spent": true,
                    "tx_index": 285722412,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 5970000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 0,
            "result": -130000,
            "size": 224,
            "time": 1538192437,
            "tx_index": 285722412,
            "vin_sz": 1,
            "hash": "f841e495e6cd3436956e5f31fd3189d108890cccfbf873357fabff2beead3831",
            "vout_sz": 2
        },
        {
            "ver": 2,
            "inputs": [
                {
                    "sequence": 4294967294,
                    "witness": "0247304402203bc34f956870f92f31cde6c01fe6fc3b0ed5d738560d272ad68a917bc93ccf7002200b2c972a41a49733aeb0cc4d25a724feb86732530c0d094bb1332baee0e8b7270121033a11b6e3e7629fcdf8de902a2bd2f8a2c840dffb64e62e861d042547758f9c27",
                    "prev_out": {
                        "spent": true,
                        "tx_index": 285703113,
                        "type": 0,
                        "addr": "2N8MALSUksHK8RqAujyPYNBUzxduYfR7Vpg",
                        "value": 26203758292,
                        "n": 0,
                        "script": "a914a5a955662fc14526c3185cb7c18a7a2939bb353587"
                    },
                    "script": "160014afb215093568b8c7e07dcc1241b00f2ad9ad10c1"
                }
            ],
            "weight": 669,
            "block_height": 1414548,
            "relayed_by": "0.0.0.0",
            "out": [
                {
                    "spent": true,
                    "tx_index": 285703527,
                    "type": 0,
                    "addr": "2NF4gm8h4tDxqPsoFnLtRm4ZAJocnVcNbrp",
                    "value": 26196657391,
                    "n": 0,
                    "script": "a914ef548a75dc51f9e11cbb9f59d92905e3ed5a1dc387"
                },
                {
                    "spent": true,
                    "tx_index": 285703527,
                    "type": 0,
                    "addr": "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB",
                    "value": 7100000,
                    "n": 1,
                    "script": "76a914fd2fdebf6a9662e8e85673569251a9997d737f0e88ac"
                }
            ],
            "lock_time": 1414547,
            "result": -1130000,
            "size": 249,
            "time": 1538133394,
            "tx_index": 285703527,
            "vin_sz": 1,
            "hash": "01c6bc398ec9467654b64681addc89f59cfa71c440c83bc6f2cc9be95d5252ee",
            "vout_sz": 2
        }
    ]}`

	var txInfo TxInfo

	err := json.Unmarshal([]byte(jsonTxs),&txInfo)
	if err != nil {
		fmt.Println("err was %v",err)
	}
	//fmt.Println(txs)


	//address := "mum6KQEwAHGDRr7HRXx7hufcsyqiGPUCtN"
	address := "n4bgbfGM1WVWicaHJhbqtXFiP7rj1CmrFB"
	changeValue := CalculateBalanceShouldChangedByAddress(address,txInfo.Txs)

	fmt.Println(changeValue)

	//fmt.Println(txs.Txs)
}






















