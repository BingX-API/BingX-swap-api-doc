Bingbon Exchange Contract Official API documentation
==================================================
[Bingbon][]Developer documentation([English Docs][])。

<!-- TOC -->

- [Introduction](#Introduction)
- [API interface encryption verification](#API interface encryption verification)
    - [Generate API Key](#Generate api-key)
    - [Make Requests](#Make Requests)
    - [Signature](#Signature)
    - [Select Timestamp](#Select Timestamp)
    - [Request interaction](#Request interaction)
        - [Requests](#Requests)
        - [Paging](#Paging)
    - [Standard Specification](#Standard Specification)
        - [Timestamp](#Timestamp)Timestamp
        - [Example](#Example)
        - [Numbers](#Numbers)
        - [Rate Limit](#Rate Limit)
            - [REST API](#rest-api)
- [Perpetual Swap API Reference](#Perpetual Swap API Reference)
    - [Perpetual Swap Market API](#Perpetual Swap Market API)
        - [1. Get All the Information of Contract Trading Pairs](#1-Get All the Information of Contract Trading Pairs)
        - [2. Get Trading Depths of Contract Trading Pairs](#2-Get Trading Depths of Contract Trading Pairs)
        - [3. Get the latest transaction record of a single contract](#3-Get the latest transaction record of a single contract)
    - [Perpetual Swap Account API](#Perpetual Swap Account API)
        - [1. Place Orders](#1-Place Orders)
        - [2. Cancel open orders](#2-Cancel open orders)
        - [3. Query Pending Order](#3-Query Pending Order)
        - [4. Query the Details of a Single Order](#4-Query the Details of a Single Order)
        - [5. Get Information of User's Account Asset ](#5-Get Information of User's Account Asset )
        - [6. Query User's Current Position](#6-Query User's Current Position) 
        - [7. Get Server Time](#7-Get Server Time) 

<!-- /TOC -->

# Introduction

Welcome to Use [Bingbon][] Development Document

This document provides an introduction to the use of related APIs such as account management, market query, and trading functions for Perpetual Swap Service.
The market API provides an open quotation data endpoints of the market. Account/Trades APIs require identity verification, which provide functions such as placing orders, canceling orders, and querying orders and account information.


# API Interface Encryption Verification
## Generate API Key

Before signing any request, you must create an API key through the Bingbon website [User Center]-[API]. After creating the key, you will get 2 pieces of information that must be remembered:
* API Key
* Secret Key


API Key and Secret Key will be randomly generated and provided.

## Make Requests

All REST requests must contain the following parameters:
* API Key -  as a string
* sign - A signature obtained by using a certain algorithm (signature information for reference).
* timestamp - As the timestamp of your request.
* All requests should contain application/json type content and be valid JSON.

## signature
sign parameter is obtained by encrypting the specific string using the **HMAC SHA256** method.The specific string is obtained after sorting all parameters (including timestamp) according to the dictionary,like key1=value1 + key2=value2 ... + Secret Key.(+ means string concatenation).

* method is the request method (POST/GET/PUT/DELETE) with all uppercase letters.


**For example: Sign the following request parameters**

```bash
curl "https://api-swap-rest.bingbon.pro/api/v1/user/getBalance"
      
```
* Get the balance information of a user's asset. Take apiKey=bingbonOneUser123, secretKey=bingbonSecondUser456 As an example
```
timestamp = 1540286290170
apiKey = bingbonOneUser123
currency = BTC
```

After the parameters are sorted by dictionary,the result is as follows
```
apiKey = bingbonOneUser123
currency = BTC
timestamp = 1540286290170
```

Generate a string to be signed

```
originString = 'apiKey=bingbonOneUser123&currency=BTC&timestamp=1540286290170'
  
```

Then, add the private key parameter to the string to be signed to generate the final string to be signed.


E.g:
```
Signature = HmacSHA256(secretkey, originString)

Signature = HmacSHA256("bingbonSecondUser456", "apiKey=bingbonOneUser123&currency=BTC&timestamp=1540286290170")

```
If the result of Signature is bingbonHashxxxxyyyyzzzz, the url query parameter after the signature is as follows
```
apiKey = ABC
currency = BTC
timestamp = 1540286290170
sign = aabbbbccccffffeeeeffff

That is, the final API request sent to the server should be:
"https://api-swap-rest.bingbon.pro/api/v1/user/getBalance?apiKey=bingbonOneUser123&currency=BTC&timestamp=1540286290170&sign=bingbonHashxxxxyyyyzzzz"
```

## Request interaction

Root URL for REST access：`https://api-swap-rest.bingbon.pro` 

### Request 

All requests are based on the HTTPS protocol, and the Content-Type in the request header information needs to be uniformly set to:'application/json'.

**Request Interaction Description**

1、Request Parameters: encapsulate parameters according to the interface request parameters.

2、Submit Request Parameters: submit the encapsulated request parameters to the server through POST/GET/DELETE, etc.

3、Server Response: The server first performs parameter security verification on the user request data, and returns the response data to the user in JSON format according to the business logic after the verification

4、Data Processing: processing the server response data.

**Success**

HTTP status code 200 indicates a successful response and may contain content. If the response contains content, it will be displayed in the corresponding return content.

**Common HTTP Error Codes**

* 4XX error codes are used to indicate wrong request content, behavior, format.

* 5XX error codes are used to indicate problems on the Bingbon service.

* 400 Bad Request – Invalid request format 

* 401 Unauthorized – Invalid API Key 

* 403 Forbidden – You do not have access to the requested resource

* 404 Not Found 

* 429 Too Many Requests - Return code is used when breaking a request rate limit.

* 418 return code is used when an IP has been auto-banned for continuing to send requests after receiving 429 codes.

* 500 Internal Server Error – We had a problem with our server

* 504 return code means that the API server has submitted a request to the business core but failed to get a response. It should be noted that the 504 return code does not mean that the request failed, but is unknown. It may have been executed, or it may fail, and further confirmation is required.

* If it fails, the response body will contain an error description

* Every interface may throw exceptions


## Standard Specification

### Timestamp

Unless otherwise specified, all timestamps in the API are returned in microseconds.

The timestamp of the request must be within 30 seconds of the API service time, otherwise the request will be considered expired and rejected. If there is a large deviation between the local server time and the API server time, we recommend that you update the HTTP header by querying the API server time.

### Example

1587091154123

### Numbers

In order to maintain the integrity of precision across platforms, decimal numbers are returned as strings. It is recommended that you also convert numbers to strings when making a request to avoid truncation and precision errors.

Integers (such as transaction number and sequence) are used without quotation marks.

### Rate Limit

If the requests are too frequent, the system will automatically limit the requests.

##### REST API

* Market Interface：We restrict the call of the public interface by IP, up to 10 requests per 1 second.

* Account and Transaction Interface: We restrict the call of the private interface by user ID, up to 10 requests per 1 second.

* The special restrictions of some interfaces are indicated on the specific interface

# Perpetual Swap API Reference

## Perpetual Swap Market API

### 1. Get All the Information of Contract Trading Pairs

**HTTP Request**

```http
    # Request
    GET api/v1/getAllContracts
    
    example： https://api-swap-rest.bingbon.pro/api/v1/market/getAllContracts
```
```javascript
    # Response
    {
        "code": 0,
        "msg": "",
        "data": [{
            "contractId": "100",
            "symbol": "BTC-USDT",
            "name": "BTC合约",
            "size": "0.0001",
            "volumePrecision": 0,
            "pricePrecision": 2,
            "feeRate": 0.001,
            "tradeMinLimit": 1,
            "currency": "USDT",
            "asset": "BTC"
        }, {
            "contractId": "101",
            "symbol": "ETH_USDT",
            "name": "ETH合约",
            "size": "0.01",
            "volumePrecision": 0,
            "pricePrecision": 2,
            "feeRate": 0.001,
            "tradeMinLimit": 1,
            "currency": "USDT",
            "asset": "ETH"
        }],
        ...
   } 
```



**Return Value Description**  


| Return Field | Description |
| ---------- |:-------:|
| code       | For error messages, 0 means normal |
| msg        | Error message description |
| contractId | ContractId |
| symbol     | Contract symbol, returned in the form of A_B  |
| name       | The name of contract product   |
| size       | Contract size, for example 0.0001 BTC |
| volumePrecision  | The precision of trading volume  |
| pricePrecision   | The precision of price |
| feeRate          | Trading fees |
| tradeMinLimit       | Minimum trading unit |
| currency   | Margin currency assets used for settlement |
| asset      | Contract transaction assets |

### 2. Get Trading Depths of Contract Trading Pairs

    Get the list of requests for the depth of the market.

**HTTP Request**

```http
    # Request
    GET api/v1/market/getMarketDepth
```

**Request Parameters**  


| Parameter name | Type  | Mandatory | Description |
| ------------- |----|----|----|
| symbol | String | YES | Trading pair symbol,like BTC-USDT |


```javascript
    # Response
    {
        "code": 0,
        "msg": "",
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
**Return Value Description**  

|Return Field| Description|  
| ------------- |----|
| code   | For error messages, 0 means normal, 1 means error|
| msg    |  Error message description |
| asks   | Sell side depth |
| bids   | Sell side depth |
| p      | price  | float64 |
| v      | volume | float64 |

### 3. Get the latest transaction record of a single contract

    Get the latest transaction record of a single contract

**HTTP Request**

  ```http
    # Request
    GET api/v1/market/getMarketTrades
```

**Request Parameters**  

| Name | Type | Mandatory | Field description | Description |
| -------|--------|--- |-------|------|
| symbol | String | YES |Contract name| The contract name needs to be underlined (BTC-USDT) |

   ```javascript
    # Response
    {
        "code": 0,
        "msg": "",
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

   **Return Value Description**  

| Return Field | Type | Mandatory | Description |
| ------------- |----|----|----|
| time      | data   |    | Closing Time |
| makerSide | String |    | The direction of contract (Bid / Ask) |
| price     | String |    | Closing Price |
| volume    | String |    | Closing Amount |


  **Remark**


    For more return error codes, please see the error code description on the homepage

### 4. Get the history of single funding fee rate 

    Get the history of single funding fee rate

**HTTP Request**

  ```http
    # Request
    GET api/v1/market/getHistoryFunding
```

**Request Parameters**  

| Name | Type | Mandatory | Field description | Description |
| -------|--------|--- |-------|------|
| symbol | String | YES |Contract name| The contract name needs to be underlined (BTC-USDT) |

   ```javascript
    # Response
    {
        "code": 0,
        "msg": "",
        "data": {
            "fundings": [
                {
                    "historyId": "687",
                    "symbol": "ETH-USDT",
                    "fundingRate": "0.3000",
                    "fairPrice": "182.73",
                    "interval": "8",
                    "time": "2019-10-28T16:00:00.000Z"
                },
                {
                    "historyId": "686",
                    "symbol": "ETH-USDT",
                    "fundingRate": "0.3000",
                    "fairPrice": "182.90",
                    "interval": "8",
                    "time": "2019-10-28T15:00:00.000Z"
                }
            ]
        }
    }
   ```

   **Return Value Description** 

| Return Field | Type | Mandatory | Description |
| ------------- |----|----|----|
| historyId     | String |    | historyId |
| fundingRate   | String |    | Funding fee rate |
| fairPrice     | String |    | Mark Price |
| interval      | String |    | The cycle of Funding fee rate settlement , unit: hour |
| time          | data   |    |  Settlement Time |


  **Remark**

    For more return error codes, please see the error code description on the homepage

### 5. Get the latest Kline/Candlestick Data

    Get the latest Kline/Candlestick Data

**HTTP Request**

  ```http
    # Request
    GET api/v1/market/getLatestKline
```

**Request Parameters**  

| Name | Type | Mandatory | Field description | Description |
| -------|--------|--- |-------|------|
| symbol | String | YES |Contract name| The contract name needs to be underlined (BTC-USDT) |
| klineType | String | YES |The type of Kline| The type of Kline(minutes，hours，weeks等等) |

**Remark**

| klineType | Field description |
| ----------|----|
| 1	        | 1min Kline |
| 3         | 3min Kline |
| 5         | 5min Kline |
| 15        | 15min Kline |
| 30        | 30min Kline |
| 60        | 1h Kline |
| 120       | 2h Kline |
| 240       | 4h Kline |
| 360       | 6h Kline |
| 720       | 12h Kline |
| 1D        | 1D Kline |
| 1W        | 1W Kline |
| 1M        | 1M Kline |

```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
            "kline": {
                "ts": 1572253500000,
                "open": 181.41,
                "close": 181.54,
                "high": 181.54,
                "low": 181.39,
                "volume": 281
            }
        }
    }
   ```

**Return Value Description** 

| Return Field | Type | Mandatory | Description |
| ------------- |----|----|----|
| open     | float64 |    | Open |
| close    | float64 |    | Close |
| high     | float64 |    | High |
| low      | float64 |    | Low |
| volume   | float64 |    | Volume |
| ts       | int64  |    | The timestamp of Kline，Unit: ms |

**Remark**
   For more return error codes, please see the error code description on the homepage

### 6. Get the history of K-line data

    Get the history of K-line data

**HTTP Request**

  ```http
    # Request
    GET api/v1/market/getHistoryKlines
```

**Request Parameters**  

| Name | Type | Mandatory | Field description | Description |
| -------|--------|--- |-------|------|
| symbol | String | YES |Contract name| The contract name needs to be underlined (BTC-USDT) |
| klineType | String | YES |The type of Kline| The type of Kline(minutes，hours，weeks等等) |
| startTs       | int64  |    | Starting timestamp, Unit: ms |
| endTs       | int64  |    | End timestamp, Unit: ms |

**Remark**

| klineType | Field description |
| ----------|----|
| 1         | 1min Kline |
| 3         | 3min Kline |
| 5         | 5min Kline |
| 15        | 15min Kline |
| 30        | 30min Kline |
| 60        | 1h Kline |
| 120       | 2h Kline |
| 240       | 4h Kline |
| 360       | 6h Kline |
| 720       | 12h Kline |
| 1D        | 1D Kline |
| 1W        | 1W Kline |
| 1M        | 1M Kline |

```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
            "klines": [
                {
                    "ts": 1572253140000,
                    "open": 181.89,
                    "close": 181.97,
                    "high": 182.04,
                    "low": 181.89,
                    "volume": 2136
                },
                {
                    "ts": 1572253200000,
                    "open": 181.94,
                    "close": 181.72,
                    "high": 181.94,
                    "low": 181.72,
                    "volume": 965
                },
                {
                    "ts": 1572253260000,
                    "open": 181.69,
                    "close": 181.72,
                    "high": 181.72,
                    "low": 181.56,
                    "volume": 1245
                },
                {
                    "ts": 1572253320000,
                    "open": 181.72,
                    "close": 181.73,
                    "high": 181.81,
                    "low": 181.69,
                    "volume": 541
                },
                {
                    "ts": 1572253380000,
                    "open": 181.77,
                    "close": 181.59,
                    "high": 181.77,
                    "low": 181.53,
                    "volume": 933
                },
                {
                    "ts": 1572253440000,
                    "open": 181.59,
                    "close": 181.38,
                    "high": 181.62,
                    "low": 181.38,
                    "volume": 1425
                },
                {
                    "ts": 1572253500000,
                    "open": 181.41,
                    "close": 181.64,
                    "high": 181.64,
                    "low": 181.39,
                    "volume": 923
                }
            ]
        }
    }
   ```

**Return Value Description** 

| Return Field | Type | Mandatory | Description |
| ------------- |----|----|----|
| klines   | array     |    | Kline data | 
| open     | float64 |    | Open |
| close    | float64 |    | Close |
| high     | float64 |    | High |
| low      | float64 |    | Low |
| volume   | float64 |    | Volume |
| ts       | int64  |    | The timestamp of Kline，Unit: ms |


**Remark**
     For more return error codes, please see the error code description on the homepage

## Trades APIs


### 1. New order interface

     Place a new order

**HTTP Request**

     
           
```http
    # Request
    POST api/v1/user/trade
```
**Request method**

    POST

**Request Parameters**  

| Name | Type | Mandatory | Description |
| ------------- |----|----|----|
| symbol | String | YES | The symbol of contract(BTC-USDT) |
| apiKey | String | YES | Interface key |
| side | String | YES | (Bid/Ask) |
| entrustPrice | float64 | YES | Price  |
| entrustVolume | float64 | YES | Volume |
| tradeType | String | YES | Market/Limit |
| action | String | YES | Open/Close  |


```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
            "orderId": "11141",
        }
    }
```
 
**Return Value Description** 

| Return Field | Type | Mandatory | Description |
| ---- |---- | ---- | ---- |
| orderId | String | YES | Order ID |



### 2. Cancel Order
 
       Cancel Order

**HTTP Request**
 
   
   
```http
    # Request
    POST api/v1/user/cancelOrder
```

**Request method**

    POST

**Request Parameters**  

| Name | Type | Mandatory | Description |
| ------------- |----|----|----|
| orderId   | String | YES | Order ID |
| symbol    | String | YES | The symbol of contract(BTC-USDT) |
| apiKey | String | YES | Interface key|


```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
            "orderId": "11141",
        }
    }
```
**Return Value Description** 
| Return Field| Type | Mandatory | Description |
| ---- |---- | ---- | ---- |
| orderId | String | YES | Order ID |


### 3. Query Current Open Order

    Query Current Open Order

**HTTP Request**

```http
    # Request
    POST api/v1/user/pendingOrders
```

**Request method**

    POST

**Request Parameters**  

| Name | Type | Mandatory | Description |
| ------------- |----|----|----|
| symbol | String | YES | The symbol of contract(BTC-USDT)，If it‘s null, return all |
| apiKey | String | YES | Interface key |

 ```javascript

# Response
    {
       "code": 0,
       "msg": "",
       "data": {
            "orders": [
                {
                    "entrustTm": "2018-04-25T15:00:51.000Z",
                    "side": "Bid",
                    "tradeType": "Limit",
                    "action": "Open",
                    "entrustPrice": 6.021954,
                    "entrustVolume": 18.098,
                    "filledVolume": 0,
                    "avgFilledPrice": 0,
                    "orderId": "6030",
                    "symbol": "BTC-USDT",
                },
                {
                    "entrustTm": "2018-04-25T15:00:51.999Z",
                    "side": "Ask",
                    "tradeType": "Limit",
                    "action": "Close",
                    "entrustPrice": 6.021954,
                    "entrustVolume": 18.098,
                    "filledVolume": 0,
                    "avgFilledPrice": 0,
                    "orderId": "6030",
                    "symbol": "ETH-USDT",
                },
            ]
        }
    }
 ```
 
**Return Value Description** 
| Return Field| Type | Mandatory | Description |
| ------------- |----|----|----|
| entrustTm     | String  | YES | Trigger time of order |
| side          | String  | YES | The direction of trading(Bid/Ask) |
| tradeType     | String  | YES | Order Type(Market/Limit) |
| action        | String  | YES | Open/Close |
| entrustPrice  | Float64 | YES | Order Price|
| entrustVolume | Float64 | YES | Order Amount |
| avgFilledPrice| Float64 | YES | Ave. Closing Price |
| filledVolume  | Float64 | YES | Closing Amount |
| orderId       | String  | YES | Order ID |

  **Remark**
  
    For more return error codes, please see the error code description on the homepage
  
### 4. Query the details of a single order
   
    Query the details of a single order

**HTTP Request**
   ```http
    # Request
    POST api/v1/user/queryOrderStatus
```
**Request method**

    POST

**Request Parameters**  

| Name | Type | Mandatory | Description |
| ------------- |----|----|----|
| apiKey | String | YES | Interface key |
| symbol | String | YES | The symbol of contract(BTC-USDT) |
| orderId | String | YES | Order ID |

```javascript
    # Response
    {
        "code": 0,
        "msg": "",
        "data": {
            "entrustTm": "2018-04-25T15:00:51.000Z",
            "side": "Ask",
            "tradeType": "Limit",
            "action": "Close",
            "entrustPrice": 6.021954,
            "entrustVolume": 18.098,
            "filledVolume": 0,
            "avgFilledPrice": 0,
            "orderId": "6030",
            "status": "Filled"
     	}
     }
```

**Return Value Description**

| Return Field | Type | Mandatory | Description |
| ------------- |----|----|----|
| entrustTm     | String  | YES | Order Time |
| side          | String  | YES | The direction of trading(Bid/Ask) |
| tradeType     | String  | YES | Order type(Market/Limit) |
| action        | String  | YES | Open/Close|
| entrustPrice  | Float64 | YES | Order Price |
| entrustVolume | Float64 | YES | Order Amount |
| avgFilledPrice| Float64 | YES | Ave. Closing Price |
| filledVolume  | Float64 | YES | Closing Amount |
| orderId       | String  | YES | Order No. |
| status        | String  | YES| The status of Order(Filled or PartiallyFilled, Pending, Cancelled, Failed) |
  


| Status | Description |
| ----------|----|
| Pending           | Order that has not been closed |
| PartiallyFilled   | Order that has been Partially filled |
| Cancelled         | Cancelled|
| Filled            | Filled  |
| Failed            | Failed  |

  **Remark**
  
    For more return error codes, please see the error code description on the homepage

### 5. Get Information of user‘s account asset 

    Get Information of user‘s account asset 
          
**HTTP Request**
             
    ```http
        # Request
        POST api/v1/user/getBalance
    ```

**Request method**

    POST

**Request Parameters**  

| Name | Type | Mandatory | Description |
| ------------- |----|----|---|---- |
| apiKey | String | YES | Interface key | |
| currency  | String | YES  | contract asset | |

    ```javascript
        # Response
            {
                "code": 0,
                "msg": "",
                "data": {
                    "userId": "123",
                    "currency": "USDT",
                    "balance": 123.33,
                    "equity": 128.99,
                    "unrealisedPNL": 1.22,
                    "realisedPNL": 8.1,
                    "availableMargin": 123.33,
                    "usedMargin": 2.2,
                    "freezedMargin": 3.3,
                    "longLeverage": 10,
                    "shortLeverage": 10,
                }
            }
    ```

**Return Value Description**

| Return Field | Type | Mandatory | Description |
| ------------- |----|----|----|
| userId	    | String | YES | User's ID |
| currency   | String | YES | User‘s asset |
| balance    | Float64 | YES | Asset Balance |
| equity     | Float64 | YES | Net Asset Value |
| unrealisedPNL  | Float64 | YES | Unrealized Profit/Loss |
| realisedPNL    | Float64 | YES | realized Profit/Loss |
| availableMargin| Float64 | YES | Available Margin |
| usedMargin     | Float64 | YES | Used Margin |
| freezedMargin  | Float64 | YES | Freezed Margin |
| longLeverage   | Float64 | YES | long Leverage |
| shortLeverage  | Float64 | YES | short Leverage |

  **Remark**
  
    For more return error codes, please see the error code description on the homepage

### 6. Query User's Current Position 

    Query User's Current Position 

**HTTP Request**

```http
    # Request
    POST api/v1/user/getPositions
```

**Request method**

    POST

**Request Parameters**  

| Name | Type | Mandatory | Description |
| ------------- |----|----|----|
| symbol | String | YES |  The symbol of contract(BTC-USDT)，If it‘s null, return all |
| apiKey | String | YES | Interface key | 

 ```javascript

# Response
    {
       "code": 0,
       "msg": "",
       "data": {
            "positions": [
                {
                    "symbol": "BTC-USDT",
                    "currency": "USDT",
                    "positionSide": "Long",
                    "marginMode": "Cross",
                    "volume": 123.33,
                    "availableVolume": 128.99,
                    "unrealisedPNL": 1.22,
                    "realisedPNL": 8.1,
                    "margin": 123.33,
                    "avgPrice": 2.2,
                    "liquidatedPrice": 2.2,
                    "leverage": 10,
                },
                {
                    "symbol": "ETH-USDT",
                    "currency": "USDT",
                    "positionSide": "Short",
                    "marginMode": "Isolated",
                    "volume": 123.33,
                    "availableVolume": 128.99,
                    "unrealisedPNL": 1.22,
                    "realisedPNL": 8.1,
                    "margin": 123.33,
                    "avgPrice": 2.2,
                    "liquidatedPrice": 2.2,
                    "leverage": 10,
                },
            ]
        }
    }
 ```
 
**Return Value Description**

| Return Field | Type | Mandatory | Description |
| ------------- |----|----|----|
| symbol         | String  | YES | Contract Type |
| currency       | String  | YES | User Assets |
| positionSide   | String  | YES | Position direction Long/Short |
| marginMode     | String  | YES | margin mode Cross/Isolated |
| volume         | Float64 | YES | Volume of position |
| availableVolume| Float64 | YES | Volume of position that can be closed |
| unrealisedPNL  | Float64 | YES | Unrealized  |
| realisedPNL    | Float64 | YES | Realised  |
| margin         | Float64 | YES | Margin |
| avgPrice       | Float64 | YES | Average open price |
| liquidatedPrice| Float64 | YES | Estimated Liquidation Price|
| leverage       | Float64 | YES | leverage |

  **Remark**
  
    For more return error codes, please see the error code description on the homepage
    
### 7. Get Server Time

    Get Server Time
 
**HTTP Request**

  ```http
        # Request
        POST api/v1/server/time
  ```
  
**Request Method**

    GET / POST

**Request Parameters**
  
    null   
             
   ```javascript
    # Response
        {
            "code": 0,
            "msg": "",
            "currentTime": 1534431933321
        }
   ```

**Return Value Description**

| Return Field | Type | Mandatory | Description |
| ------------- |----|----|----|
| currentTime |Int64  | YES | The current time of the system，unit:ms |


  **Remark**
  
    For more return error codes, please see the error code description on the homepage


### 8. Cancel Multiple Orders
  
  Cancel Multiple Orders

**HTTP Request**


```http
    # Request
    POST api/v1/user/batchCancelOrders
```

**Request Method**

    POST

**Request Parameters**  

| Name | Type | Mandatory | Description |
| ------------- |----|----|----|
| symbol | String | YES |  The symbol of contract(BTC-USDT) |
| apiKey | String | YES | Interface key | 


```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
        }
    }
```
**Return Value Description**
| Return Field | Type | Mandatory | Description |
| ---- |---- | ---- | ---- |
| orderId | String | YES | Order ID |

  **Remark**
  
    For more return error codes, please see the error code description on the homepage

    
[Bingbon]: https://bingbon.pro
[English Docs]: https://bingbon.pro
[Unix Epoch]: https://en.wikipedia.org/wiki/Unix_time
