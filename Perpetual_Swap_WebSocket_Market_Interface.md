Official API Documentation for the Bingx Trading Platform- Websocket
==================================================
Bingx Developer Documentation

<!-- TOC -->

- [Official API Documentation for the bingx Trading Platform- Websocket](#official-api-documentation-for-the-bingx-trading-platform--websocket)
- [Introduction](#introduction)
  - [Access](#access)
  - [Data Compression](#data-compression)
  - [Heartbeats](#heartbeats)
  - [Subscriptions](#subscriptions)
  - [Unsubscribe](#unsubscribe)
- [Perpetual Swap Websocket Market Data](#perpetual-swap-websocket-market-data)
  - [1. Subscribe Market Depth Data](#1-subscribe-market-depth-data)
  - [2. Subscribe the Latest Trade Detail](#2-subscribe-the-latest-trade-detail)
  - [3. Subscribe K-Line Data](#3-subscribe-k-line-data)
- [Websocket Account Data](#Websocket-account-data)
  - [Subscribe to account balance push](#1-Subscribe-to-account-balance-push)

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
| data | Push Data |
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
            "asks": [{
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
            "bids": [{
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
| data | Push Data |
| trades    | Deal by deal |
| time      | Closing Time |
| makerSide | Direction ( Bid / Ask) |
| price     | Closing Price |
| volume    | Filled Amount |

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

## 3. Subscribe K-Line Data

    Subscribe to market k-line data of one trading pair

**Subscription Type**

    The dataType is market.kline.$Symbol.$KlineType. 
    E.g. market.kline.BTC-USDT.1min

**Subscription Parameters**

| Parameters | Type   | Required | Field Description   | Description                                                  |
| ---------- | ------ | -------- | ------------------- | ------------------------------------------------------------ |
| symbol     | String | YES      | Trading pair symbol | There must be a hyphen/ "-" in the trading pair symbol. eg: BTC-USDT |
| klineType  | String | YES      | K-Line Type         | The type of K-Line ( minutes, hours, weeks etc.)             |

**Remarks**

| klineType | Field Description |
|-----------|-------------------|
| 1min      | 1 min Kline       |
| 3min      | 3 min Kline       |
| 5min      | 5 min Kline       |
| 15min     | 15 min Kline      |
| 30min     | 30 min Kline      |
| 1hour     | 1-hour Kline      |
| 2hour     | 2-hour Kline      |
| 4hour     | 4-hour Kline      |
| 6hour     | 6-hour Kline      |
| 12hour    | 12-hour Kline     |
| 1day      | 1-Day Kline       |
| 1week     | 1-Week Kline      |
| 1month    | 1-Month Kline     |

**Push Data**

| Return Parameters | Field Description                                            |
| ----------------- | ------------------------------------------------------------ |
| code              | With regards to error messages, 0 means normal, and 1 means error |
| data              | Push Data                                                    |
| dataType          | Data Type                                                    |
| klineInfosVo      | Kline Data                                                   |
| close             | Closing Price                                                |
| high              | High                                                         |
| low               | Low                                                          |
| open              | Opening Price                                                |
| statDate          | Kline Date                                                   |
| time              | The timestamp of K-Line，Unit: ms                            |
| volume            | Volume                                                       |

```javascript
 # Response
 {
     "code": 0,
     "data": {
         "klineInfosVo": [
             {
                 "close": 54564.31, 
                 "high": 54711.73,
                 "low": 54418.27,
                 "open": 54577.41, 
                 "statDate": "2021-04-29T11:00:00.000+0800", 
                 "time": 1619665200000, 
                 "volume": 1607.0727000000002
             }
         ]
     },
     "dataType": "market.kline.BTC-USDT.30min"
 }
```

# Websocket Account Data

- Note that websocket authentication is required to obtain such information, use listenKey, and see the [Rest interface documentation for details](https://github.com/BingX-API/BingX-swap-api-doc/blob/master/Perpetual_Swap_API_Documentation.md#other-interface).
- The base URL of Websocket Market Data is: `wss://open-ws-swap.bingbon.pro/ws`
- User Data Streams are accessed at `/$listenKey`

```
wss://open-ws-swap.bingbon.pro/ws/94bE3nW8BuyGCUsvjRKPPRt1lDomEeJlEO8ABMLxYM6rT92u
```

## 1. Subscribe to account balance push

**Subscription Type**
```
dataType is ACCOUNT_UPDATE
```

**Subscription Example**
```
{"id":"gdfg2311-d0f6-4a70-8d5a-043e4c741b40","dataType":"ACCOUNT_UPDATE"}
```

```
The field "m" represents the reason for the launch of the event, including the following possible types:

    - DEPOSIT
    - WITHDRAW
    - FUNDING_FEE
    - ORDER
```

**Push Data**

| return field | field description                      |  
|----|---------------------------   |
| e  | Event Type             |
| E  | Event Time             |
| m  | Event reason type              |
| a  | Asset             |
| wb   | Wallet Balance          |
| cw   | Cross Wallet Balance          |
| bc  | Balance Change except PnL and Commission             |

```
# Response
{
	"e": "ACCOUNT_UPDATE", // Event Type
	"E": 1671159080000,  // Event Time
	"a": {
		"B": [{
			"a": "USDT", // Asset
			"bc": "-66", // Balance Change except PnL and Commission
			"cw": "4470.890393795533", // Cross Wallet Balance
			"wb": "4499.53918561" // Wallet Balance
		}],
		"m": "WITHDRAW" // Event reason type 
	}
}
```

**Remarks**

    For more about return error codes, please see the error code description on the homepage.
