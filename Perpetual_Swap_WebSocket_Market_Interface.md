Bingbon Exchange Contract Official API documentation
==================================================
[Bingbon][]Developer documentation([English Docs][])。

<!-- TOC -->

- [Introduction](#Introduction)
- [Perpetual Swap Websocket Market Data](#Perpetual Swap Websocket Market Data Reference) 
    - [Perpetual Swap Websocket Market API](#Perpetual Swap Websocket Market API)
        - [1. Subscribe to the Trading Depth of Swap Trading Pairs](#1-Subscribe to the Trading Depth of Swap Trading Pairs)
        - [2. Subscribe to The Latest Transaction Records of a Single Swap](#2-Subscribe to The Latest Transaction Records of a Single Swap)

<!-- /TOC -->

# Introduction

## The Base URL

the base URL of Market Websocket：`wss://open-ws-swap.bingbon.pro/ws`

## Data Compression

All data returned by the WebSocket market interface is GZIP compressed, and the client needs to decompress after receiving the data.

## Heartbeat

When the user's Websocket client connects to the Bingbon Websocket server, the server will periodically (currently set as 50 seconds) to send a heartbeat string Ping to it.

When the user's Websocket client receives this heartbeat message, it should return the string Pong message

## Subscribe

After successfully establishing a connection with the Websocket server, the Websocket client sends the following request to subscribe to a specific topic:

{
  "id": "id1",
  "reqType": "sub",
  "dataType": "data to sub",
}

After a successful subscription, the Websocket client will receive a confirmation:

{
  "id": "id1",
  "code": 0,
  "msg": "",
}
After that, once the subscribed data is updated, the Websocket client will receive the update message pushed by the server

## Unsubscribe
The format of unsubscription is as follows:

{
  "id": "id1",
  "reqType": "unsub",
  "dataType": "data to unsub",
}

Confirmation of cancellation of subscription:

{
  "id": "id1",
  "code": 0,
  "msg": "",
}


# Perpetual Swap Websocket API Reference

## Perpetual Swap Websocket Market

### 1. Subscribe to the Trading Depth of Swap Trading Pairs

     Subscribe to the push of contract trading pair market depth data, this topic sends the latest market depth snapshot. The frequency of snapshot  is 1 time per second.

**Subscription Type**

    dataType is market.depth.$Symbol.$Step.$Level，such as market.depth.BTC-USDT.step0.level5

**Subscription Parameters**  

| Name | Type | Mandatory | Description |
| ------------- |----|----|----|
| symbol | String | YES | Trading pair, like BTC-USDT |
| step | String | YES | Combine depth type, step0,step1,step2,step3,step4,step5 |  
| level | String | YES |  Number of levels, such as level5,level10,level20,level50,level100 | 

"step" Combine Depth Type
| Name | Description | 
| ----- |----|
| step0 | Do not merge depth |   
| step1 | Multiply the minimum precision of the price by 10 to merge the depth data |  
| step2 | Multiply the minimum precision of the price by 100 to merge the depth data | 
| step3 | Multiply the minimum precision of the price by 1000 to merge the depth data |
| step4 | Multiply the minimum precision of the price by 10000 to merge the depth data |
| step5 | Multiply the minimum precision of the price by 100000 to merge the depth data |

"level" Depth Level Definition
| Name | Description |
| -------- |----|
| level5   | level 5  |
| level10  | level 10 |
| level20  | level 20 |
| level50  | level 50 |
| level100 | level 100|

**Push Data** 

| Return Field | Description |  
| ------------- |----|
| code   | For error messages, 0 means normal, 1 means error|
| dataType | The type of subscribed data, such as market.depth.BTC-USDT.step0.level5 |  
| data | struct | YES | Push content |
| asks   | Sell side depth |  
| bids   | Buy side depth |
| p      | price  | float64 
| v      | volume | float64 

```javascript
    # Response
    {
        "code": 0,
        "dataType": "market.depth.BTC-USDT.step0.level5",
        "data": {
            "asks": [[
                    "p": 5319.94,
                    "v": 0.05483456
                ],[
                    "p": 5320.19,
                    "v": 1.05734545
                ],[
                    "p": 5320.39,
                    "v": 1.16307999
                ],[
                    "p": 5319.94,
                    "v": 0.05483456
                ],[
                    "p": 5320.19,
                    "v": 1.05734545
                ],[
                    "p": 5320.39,
                    "v": 1.16307999
                ],
            ],
            "bids": [[
                    "p": 5319.94,
                    "v": 0.05483456
                ],[
                    "p": 5320.19,
                    "v": 1.05734545
                ],[
                    "p": 5320.39,
                    "v": 1.16307999
                ],[
                    "p": 5319.94,
                    "v": 0.05483456
                ],[
                    "p": 5320.19,
                    "v": 1.05734545
                ],[
                    "p": 5320.39,
                    "v": 1.16307999
                ],
            ],
        }
    }
```


### 2. Subscribe to The Latest Transaction Records of a Single Swap

    Subscribe to the transaction details of a single Swap

**Subscription Type**

    dataType is market.trade.detail.$Symbol，such as market.trade.detail.BTC-USDT

**Subscription Parameters**

| Name | Type | Mandatory | Field description | Description |
| -------|--------|--- |-------|------|
| symbol | String | YES | Swap name | The Swap name needs to be underlined (BTC-USDT) |

**Push Data**

| Return Field | Description |  
| ------------- |----|
| code   | For error messages, 0 means normal, 1 means error | 
| dataType | The type of data subscribed, such as market.tradeDetail.BTC-USDT |
| data | struct | YES | Push content |
| trades    | Deal by deal | 
| time      | data   |    | Closing Time |
| makerSide | String |    | The direction of Swap (Bid / Ask) |
| price     | String |    | Closing Price |
| volume    | String |    | Closing Amount |

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


**Remark**

    For more return error codes, please see the error code description on the homepage

    
[Bingbon]: https://bingbon.pro
[English Docs]: https://bingbon.pro
[Unix Epoch]: https://en.wikipedia.org/wiki/Unix_time
