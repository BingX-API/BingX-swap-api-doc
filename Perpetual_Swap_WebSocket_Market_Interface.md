Official API Documentation for the Bingbon Trading Platform- Websocket
==================================================
[Bingbon][] Developer Documentation ([English Docs][])

<!-- TOC -->

- [Introduction](#Introduction)
    - [Access](#Access)
    - [Data Compression](#Data Compression)
    - [Heartbeats](#Heartbeats)
    - [Subscriptions](#Subscriptions)
    - [Unsubscribe](#Unsubscribe)
- [Perpetual Swap Websocket Market Data](#Perpetual Swap Websocket Market Data Reference) 
    
    1. [Subscribe Market Depth Data](#Subscribe Market Depth Data)
    2. [Subscribe the Latest Trade Detail](#Subscribe the Latest Trade Detail)

<!-- /TOC -->

# Introduction

## Access

the base URL of Websocket Market Data ：`wss://open-ws-swap.bingbon.pro/ws`

## Data Compression

All response data from Websocket server are compressed into GZIP format. Clients have to decompress them for further use.

## Heartbeats

Once the Websocket Client and Websocket Server get connected, the server will send a heartbeat- Ping message every 5 seconds (the frequency might change).

When the Websocket Client receives this heartbeat message, it should return Pong message.

## Subscriptions

After successfully establishing a connection with the Websocket server, the Websocket client sends the following request to subscribe to a specific topic:

{
  "id": "id1",
  "reqType": "sub",
  "dataType": "data to sub",
}

After a successful subscription, the Websocket client will receive a confirmation message:

{
  "id": "id1",
  "code": 0,
  "msg": "",
}
After that, once the subscribed data is updated, the Websocket client will receive the update message pushed by the server.

## Unsubscribe
The format of unsubscription is as follows:

{
  "id": "id1",
  "reqType": "unsub",
  "dataType": "data to unsub",
}

Confirmation of Unsubscription:

{
  "id": "id1",
  "code": 0,
  "msg": "",
}


# Perpetual Swap Websocket Market Data

## 1. Subscribe Market Depth Data

    Subscribe to the push of a certain trading pair's market depth data; this topic sends the latest market depth as a snapshot. Snapshots are sent at a frequency of once every 1 second.

**Subscription Type**

    The dataType is market.depth.$Symbol.$Step.$Level. E.g. market.depth.BTC-USDT.step0.level5

**Subscription Parameters**  

| Parameters | Type | Required | Description |
| ------------- |----|----|----|
| symbol | String | YES | There must be a hyphen/ "-" in the trading pair symbol. eg: BTC-USDT |
| step | String | YES | Merged depth, step0,step1,step2,step3,step4,step5 |
| level | String | YES | Depth level, such as level5,level10,level20,level50,level100 |

"step" Merged Depth
| Parameters | Description |
| ----- |----|
| step0 | Depth data will not be merged. |
| step1 | Multiply the minimum precision of the price by 10 to merge the depth data |
| step2 | Multiply the minimum precision of the price by 100 to merge the depth data |
| step3 | Multiply the minimum precision of the price by 1,000 to merge the depth data |
| step4 | Multiply the minimum precision of the price by 10,000 to merge the depth data |
| step5 | Multiply the minimum precision of the price by 100,000 to merge the depth data |

"level" Depth Level
| Parameters | Description |
| -------- |----|
| level5   | level 5  |
| level10  | level 10 |
| level20  | level 20 |
| level50  | level 50 |
| level100 | level 100|

**Push Data** 

| Return Parameters | Description |
| ------------- |----|
| code   | With regards to error messages, 0 means normal, and 1 means error |
| dataType | The type of subscribed data, such as market.depth.BTC-USDT.step0.level5 |
| data | struct |
| asks   | Sell side depth |
| bids   | Buy side depth |
| p | price |
| v | volume |
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


## 2. Subscribe the Latest Trade Detail

    Subscribe to the trade detail data of a trading pair

**Subscription Type**

    The dataType is market.trade.detail.$Symbol. 
    E.g. market.trade.detail.BTC-USDT

**Subscription Parameters**

| Parameters | Type | Required | Field description | Description |
| -------|--------|--- |-------|------|
| symbol | String | YES | Trading pair symbol | There must be a hyphen/ "-" in the trading pair symbol. eg: BTC-USDT |

**Push Data**

| Return Parameters | Description |
| ------------- |----|
| code   | With regards to error messages, 0 means normal, and 1 means error |
| dataType | The type of data subscribed, such as market.tradeDetail.BTC-USDT |
| data | struct |
| trades    | Deal by deal |
| time      | data   |
| makerSide | String |
| price     | String |
| volume    | String |

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

**Remarks**

    For more about return error codes, please see the error code description on the homepage.


​    
[Bingbon]: https://bingbon.pro
[English Docs]: https://bingbon.pro
[Unix Epoch]: https://en.wikipedia.org/wiki/Unix_time
