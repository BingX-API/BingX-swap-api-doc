Bingbon交易所官方API文档
==================================================
[Bingbon][]交易所开发者文档([English Docs][])。

<!-- TOC -->

- [Websocket 简介](#Websocket 简介)
    - [接入方式](#接入方式)
    - [数据压缩](#数据压缩)
    - [心跳信息](#心跳信息)
    - [订阅方式](#订阅方式)
    - [取消订阅](#取消订阅)
- [Websocket 行情推送](#Websocket 行情推送)
    - [订阅合约交易深度](#1-订阅合约交易深度)
    - [订单最新成交记录](#2-订单最新成交记录)

<!-- /TOC -->

# Websocket 介绍

## 接入方式

行情Websocket的接入URL：`wss://open-ws-swap.bingbon.pro/ws`

## 数据压缩

WebSocket 行情接口返回的所有数据都进行了 GZIP 压缩，需要 client 在收到数据之后解压。

## 心跳信息

当用户的Websocket客户端连接到Bingbon Websocket服务器后，服务器会定期（当前设为5秒）向其发送心跳字符串Ping，

当用户的Websocket客户端接收到此心跳消息后，应返回字符串Pong消息

## 订阅方式

成功建立与Websocket服务器的连接后，Websocket客户端发送如下请求以订阅特定主题：

{
  "id": "id1",
  "reqType": "sub",
  "dataType": "data to sub",
}

成功订阅后，Websocket客户端将收到确认：

{
  "id": "id1",
  "code": 0,
  "msg": "",
}
之后, 一旦所订阅的数据有更新，Websocket客户端将收到服务器推送的更新消息

## 取消订阅
取消订阅的格式如下：

{
  "id": "id1",
  "reqType": "unsub",
  "dataType": "data to unsub",
}

取消订阅成功确认：

{
  "id": "id1",
  "code": 0,
  "msg": "",
}


# Websocket 行情推送

## 1. 订阅合约交易深度

    订阅合约对盘口深度的数据的推送，此主题发送最新市场深度快照。快照频率为每秒1次。

**订阅类型**

    dataType 为 market.depth.$Symbol.$Step.$Level，比如market.depth.BTC-USDT.step0.level5

**订阅参数**  


| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| symbol | String | 是 | 合约名称中需有"-"，如BTC-USDT |
| step | String | 是 | 合并深度类型，step0，step1，step2，step3，step4，step5 |
| level | String | 是 | 档数, 如 level5，level10，level20，level50，level100 |

"step" 合并深度类型
| 参数名 | 描述 |
| ----- |----|
| step0 | 不合并深度 |
| step1 | 按价格最小精度乘以10合并深度数据 |
| step2 | 按价格最小精度乘以100合并深度数据 |
| step3 | 按价格最小精度乘以1000合并深度数据 |
| step4 | 按价格最小精度乘以10000合并深度数据 |
| step5 | 按价格最小精度乘以100000合并深度数据 |

"level" 深度档数定义
| 参数名 | 描述 |
| ----- |----|
| level5 | 5档 |
| level10 | 10档 |
| level20 | 20档 |
| level50 | 50档 |
| level100 | 100档 |

**推送数据** 

| 返回字段|字段说明|  
| ------------- |----|
| code   | 是否有错误信息，0为正常，1为有错误 |
| dataType | 订阅的数据类型，例如 market.depth.BTC-USDT.step0.level5 |
| data | struct | 是 | 推送内容 |
| asks   | 卖方深度 |
| bids   | 买方深度 |
| p      | price价格  | float64
| v      | volume数量 | float64

```javascript
    # Response
    {
        "code": 0,
        "dataType": "market.depth.BTC-USDT.step0.level5",
        "data": {
            "asks": [
                {
                    "p": 5319.94,
                    "v": 0.05483456
                },{
                    "p": 5320.19,
                    "v": 1.05734545
                },{
                    "p": 5320.39,
                    "v": 1.16307999
                },{
                    "p": 5319.94,
                    "v": 0.05483456
                },{
                    "p": 5320.19,
                    "v": 1.05734545
                },{
                    "p": 5320.39,
                    "v": 1.16307999
                },
            ],
            "bids": [
                {
                    "p": 5319.94,
                    "v": 0.05483456
                },{
                    "p": 5320.19,
                    "v": 1.05734545
                },{
                    "p": 5320.39,
                    "v": 1.16307999
                },{
                    "p": 5319.94,
                    "v": 0.05483456
                },{
                    "p": 5320.19,
                    "v": 1.05734545
                },{
                    "p": 5320.39,
                    "v": 1.16307999
                },
            ],
        }
    }
```


## 2. 订单最新成交记录

    订阅单个合约的逐笔成交明细

**订阅类型**

    dataType 为 market.trade.detail.$Symbol，比如market.trade.detail.BTC-USDT

**订阅参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 是 |合约名称| 合约名称中需有"-"，如BTC-USDT |

**推送数据** 

| 返回字段|字段说明|  
| ------------- |----|
| code   | 是否有错误信息，0为正常，1为有错误 |
| dataType | 订阅的数据类型，例如 market.tradeDetail.BTC-USDT |
| data | struct | 是 | 推送内容 |
| trades    | 逐笔成交 |
| time      | data   |    | 成交时间 |
| makerSide | String |    | 吃单方向(Bid / Ask 买/卖) |
| price     | String |    | 成交价格 |
| volume    | String |    | 成交数量 |

   ```javascript
    # Response
    {
        "code": 0,
        "dataType": "market.tradeDetail.BTC-USDT",
        "data": {
            "trades": [
                {
                    "time": "2018-04-25T15:00:51.999Z",
                    "makerSide": "Bid",
                    "price": 0.279563,
                    "volume": 100,
                },
                {
                    "time": "2018-04-25T15:00:51.000Z",
                    "makerSide": "Ask",
                    "price": 0.279563,
                    "volume": 300,
                }
            ]
        }
    }
   ```


  **备注**

    更多返回错误代码请看首页的错误代码描述

    
[Bingbon]: https://bingbon.pro
[English Docs]: https://bingbon.pro
[Unix Epoch]: https://en.wikipedia.org/wiki/Unix_time
