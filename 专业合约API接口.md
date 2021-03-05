Bingbon交易所官方API文档
==================================================
[Bingbon][]交易所开发者文档([English Docs][])。

<!-- TOC -->

- [介绍](#介绍)
- [API接口加密验证](#api接口加密验证)
    - [生成API Key](#生成api-key)
    - [发起请求](#发起请求)
    - [签名](#签名)
    - [选择时间戳](#选择时间戳)
    - [请求交互](#请求交互)
        - [请求](#请求)
        - [分页](#分页)
    - [标准规范](#标准规范)
        - [时间戳](#时间戳)
        - [例子](#例子)
        - [数字](#数字)
        - [限流](#限流)
            - [REST API](#rest-api)
- [永续合约业务API参考](#永续合约业务api参考)
    - [永续合约行情API](#永续合约行情api)
        - [1. 获取所有合约交易对信息](#1-获取所有合约交易对信息)
        - [2. 获取合约交易对交易深度](#2-获取合约交易对交易深度)
        - [3. 获取单个合约的最新成交记录](#3-获取单个合约的最新成交记录)
        - [4. 获取单个资金费率历史记录](#4-获取单个资金费率历史记录)
        - [5. 获取最新一条K线数据](#5-获取最新一条K线数据)
        - [6. 获取K线历史数据](#6-获取K线历史数据)
        - [7. 获取合约的最新价格信息](#7-获取合约的最新价格信息)
        - [8. 获取合约当前最新的资金费率](#8-获取合约当前最新的资金费率)
        - [9. 查询合约未平仓数量](#9-查询合约未平仓数量)
    - [永续合约账户API](#永续合约账户api)
        - [1. 下单交易](#1-下单交易)
        - [2. 撤销订单](#2-撤销订单)
        - [3. 查询用户委托中的订单](#3-查询用户委托中的订单)
        - [4. 查询单个订单的详情信息](#4-查询单个订单的详情信息)
        - [5. 获取用户账户资产信息](#5-获取用户账户资产信息)
        - [6. 查询用户持仓信息](#6-查询用户持仓信息)
        - [7. 批量撤销订单](#7-批量撤销订单)
        - [8. 一键平仓下单](#8-一键平仓下单)
        - [9. 全部一键平仓下单](#9-全部一键平仓下单)
        - [10. 撤销全部订单](#10-撤销全部订单)
        - [11. 修改账户保证金模式](#11-修改账户保证金模式)
        - [12. 修改杠杆](#12-修改杠杆)
    - [永续合约基础信息API](#永续合约基础信息api)
        - [1. 获取系统时间](#1-获取系统时间)

<!-- /TOC -->

# 介绍

欢迎使用[Bingbon][]开发者文档。

本文档提供了永续合约交易业务的账户管理、行情查询、交易功能等相关API的使用方法介绍。
行情API提供市场的公开的行情数据接口，账户和交易API需要身份验证，提供下单、撤单，查询订单和帐户信息等功能。

# API接口加密验证
## 生成API Key

在对任何请求进行签名之前，您必须通过 Bingbon 网站【用户中心】-【API】创建一个API key。 创建key后，您将获得2个必须记住的信息：
* API Key
* Secret Key


API Key 和 Secret Key将由随机生成和提供。

## 发起请求

所有REST请求都必须包含以下参数：

* API Key作为一个字符串。
* sign 使用一定算法得出的签名（请参阅签名信息）。
* timestamp 作为您的请求的时间戳。
* 所有请求都应该含有application/json类型内容，并且是有效的JSON。

## 签名
sign 参数是对 **所有参数(含timestamp)按照字典排序之后，按照key1=value1 + key2=value2 ... + Secret Key** 字符串(+表示字符串连接)使用 **HMAC SHA256** 方法加密而得到的。

* method 是请求方法(POST/GET/PUT/DELETE)，字母全部大写。


**例如：对于如下的请求参数进行签名**

```bash
curl "https://api-swap-rest.bingbon.pro/api/v1/user/getBalance"
      
```
* 获取获取用户某资产余额信息，以 apiKey=bingbonOneUser123, secretKey=bingbonSecondUser456 为例
```
timestamp = 1540286290170
apiKey = bingbonOneUser123
currency = BTC
```

按字典排序之后，为
```
apiKey = bingbonOneUser123
currency = BTC
timestamp = 1540286290170
```

生成待签名的字符串

```
originString = 'apiKey=bingbonOneUser123&currency=BTC&timestamp=1540286290170'
  
```

然后，将待签名字符串添加私钥参数生成最终待签名字符串。


例如：
```
Signature = HmacSHA256(secretkey, originString)
即：
Signature = HmacSHA256("bingbonSecondUser456", "apiKey=bingbonOneUser123&currency=BTC&timestamp=1540286290170")

```
假如Signature的结果为bingbonHashxxxxyyyyzzzz，则签名之后的url query参数为
```
apiKey = ABC
currency = BTC
timestamp = 1540286290170
sign = aabbbbccccffffeeeeffff

即最终发送给服务器的API请求应该为：
"https://api-swap-rest.bingbon.pro/api/v1/user/getBalance?apiKey=bingbonOneUser123&currency=BTC&timestamp=1540286290170&sign=bingbonHashxxxxyyyyzzzz"
```

## 请求交互  

REST访问的根URL：`https://api-swap-rest.bingbon.pro`

### 请求

所有请求基于Https协议，请求头信息中Content-Type 需要统一设置为:'application/json’。

**请求交互说明**

1、请求参数：根据接口请求参数规定进行参数封装。

2、提交请求参数：将封装好的请求参数通过POST/GET/DELETE等方式提交至服务器。

3、服务器响应：服务器首先对用户请求数据进行参数安全校验，通过校验后根据业务逻辑将响应数据以JSON格式返回给用户。

4、数据处理：对服务器响应数据进行处理。

**成功**

HTTP状态码200表示成功响应，并可能包含内容。如果响应含有内容，则将显示在相应的返回内容里面。

**常见HTTP错误码**

* 4XX 错误码用于指示错误的请求内容、行为、格式

* 5XX 错误码用于指示Bingbon服务侧的问题

* 400 Bad Request – Invalid request format 请求格式无效

* 401 Unauthorized – Invalid API Key 无效的API Key

* 403 Forbidden – You do not have access to the requested resource 请求无权限

* 404 Not Found 没有找到请求

* 429 Too Many Requests 请求太频繁被系统限流

* 418 表示收到429后继续访问，于是被封了

* 500 Internal Server Error – We had a problem with our server 服务器内部错误

* 504 表示API服务端已经向业务核心提交了请求但未能获取响应，特别需要注意的是504代码不代表请求失败，而是未知。很可能已经得到了执行，也有可能执行失败，需要做进一步确认

* 如果失败，response body 带有错误描述信息

* 每个接口都有可能抛出异常


## 标准规范

### 时间戳

除非另外指定，API中的所有时间戳均以微秒为单位返回。

请求的时间戳必须在API服务时间的30秒内，否则请求将被视为过期并被拒绝。如果本地服务器时间和API服务器时间之间存在较大的偏差，那么我们建议您使用通过查询API服务器时间来更新http header。

### 例子

1587091154123

### 数字

为了保持跨平台时精度的完整性，十进制数字作为字符串返回。建议您在发起请求时也将数字转换为字符串以避免截断和精度错误。 

整数（如交易编号和顺序）不加引号。

### 限流

如果请求过于频繁系统将自动限制请求。

##### REST API

* 行情接口：我们通过IP限制公共接口的调用：每1秒最多10个请求。

* 账户和交易接口：我们通过用户ID限制私人接口的调用：每1秒最多10个请求。

* 某些接口的特殊限制在具体的接口上注明

# 永续合约业务API参考

## 永续合约行情API

### 1. 获取所有交易对信息

**HTTP请求**

```http
    # Request
    GET api/v1/market/getAllContracts
    
    example： https://api-swap-rest.bingbon.pro/api/v1/market/getAllContracts
```

**返回值说明**


|返回字段 | 字段说明|
| ---------- |:-------:|
| code       | 是否有错误信息，0为正常，1为有错误|
| msg        | 错误信息描述
| contractId | 合约ID
| symbol     | 合约产品符号，以A_B的形式返回  |
| name       | 合约产品名字 |
| size       | 合约大小，例如0.0001 BTC |
| volumePrecision  | 交易数量精度 |
| pricePrecision   | 价格精度 |
| feeRate          | 交易手续费 |
| tradeMinLimit    | 交易最小单位，单位为张 |
| currency   | 结算和保证金货币资产 |
| asset      | 合约交易资产 |

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

### 2. 获取合约交易对的交易深度

    获取合约对盘口深度的请求列表。

**HTTP请求**

```http
    # Request
    GET api/v1/market/getMarketDepth
```

**请求参数**


| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| symbol | String | 是 | 币对, 如 BTC-USDT |
| level | String | 是 | 层数，没指定则默认返回5层 |

**返回值说明**

|返回字段|字段说明|
| ------------- |----|
| code   | 是否有错误信息，0为正常，1为有错误|
| msg    | 错误信息描述
| asks   | 卖方深度 |
| bids   | 买方深度 |
| p      | price价格  | float64
| v      | volume数量 | float64


```javascript
    # Response
    {
        "code": 0,
        "msg": "",
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
                "p": 5320.94,
                "v": 0.05483456
              },{
                "p": 5330.19,
                "v": 1.05734545
              },{
                "p": 5330.39,
                "v": 1.16307999
              },
            ],
            "bids": [
              {
                "p": 5319.93,
                "v": 0.05483456
              },{
                "p": 5318.19,
                "v": 1.05734545
              },{
                "p": 5317.39,
                "v": 1.16307999
              },{
                "p": 5316.94,
                "v": 0.05483456
              },{
                "p": 5315.19,
                "v": 1.05734545
              },{
                "p": 5314.39,
                "v": 1.16307999
              },
            ],
        }
    }
```

### 3. 获取单个合约的最新成交记录

    获取单个合约的最新成交记录

**HTTP请求**

  ```http
    # Request
    GET api/v1/market/getMarketTrades
```

**请求参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 是 |合约名称| 合约名称后面需要加下划线(BTC-USDT) |

**返回值说明**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| time      | data   |    | 成交时间 |
| makerSide | String |    | 吃单方向(Buy / Sell 买/卖) |
| price     | String |    | 成交价格 |
| volume    | String |    | 成交数量 |

```javascript
    # Response
    {
        "code": 0,
        "msg": "",
        "data": {
            "trades": [
                {
                    "time": "2018-04-25T15:00:51.999Z",
                    "makerSide": "Buy",
                    "price": 0.279563,
                    "volume": 100,
                },
                {
                    "time": "2018-04-25T15:00:51.000Z",
                    "makerSide": "Sell",
                    "price": 0.279563,
                    "volume": 300,
                }
            ]
        }
    }
```


### 4. 获取单个资金费率历史记录

    获取单个资金费率历史记录

**HTTP请求**

  ```http
    # Request
    GET api/v1/market/getHistoryFunding
```

**请求参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 是 |合约名称| 合约名称后面需要加下划线(BTC-USDT) |

**返回值说明**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| historyId     | String |    | 历史ID号 |
| fundingRate   | String |    | 资金费率 |
| fairPrice     | String |    | 标记价格 |
| interval      | String |    | 资金费率结算周期，单位：小时 |
| time          | data   |    | 结算时间 |

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


### 5. 获取最新一条K线数据

    获取最新一条K线数据

**HTTP请求**

  ```http
    # Request
    GET api/v1/market/getLatestKline
```

**请求参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 是 |合约名称| 合约名称后面需要加下划线(BTC-USDT) |
| klineType | String | 是 |k线类型| k线类型(分钟，小时，周等等) |

**备注**

| klineType 字段说明  | |
| ----------|----|
| 1	        | 1m一分钟K线 |
| 3         | 3m三分钟K线 |
| 5         | 5m五分钟K线 |
| 15        | 15m十五分钟K线 |
| 30        | 30m三十分钟K线 |
| 60        | 1h一小时K线 |
| 120       | 2h两小时K线 |
| 240       | 4h四小时K线 |
| 360       | 6h六小时K线 |
| 720       | 12h十二小时K线 |
| 1D        | 1D日K线 |
| 1W        | 1W周K线 |
| 1M        | 1M月K线 |

**返回值说明**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| open     | float64 |    | 开盘价 |
| close    | float64 |    | 收盘价 |
| high     | float64 |    | 最高价 |
| low      | float64 |    | 最低价 |
| volume   | float64 |    | 交易数量 |
| ts       | int64  |    | k线时间戳，单位毫秒 |

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


### 6. 获取K线历史数据

    获取K线历史数据

**HTTP请求**

  ```http
    # Request
    GET api/v1/market/getHistoryKlines
```

**请求参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 是 |合约名称| 合约名称后面需要加下划线(BTC-USDT) |
| klineType | String | 是 |k线类型| k线类型(分钟，小时，周等等) |
| startTs       | int64  |    | 起始时间戳，单位毫秒 |
| endTs       | int64  |    | 结束时间戳，单位毫秒 |

**备注**

| klineType 字段说明  | |
| ----------|----|
| 1	        | 1m一分钟K线 |
| 3         | 3m三分钟K线 |
| 5         | 5m五分钟K线 |
| 15        | 15m十五分钟K线 |
| 30        | 30m三十分钟K线 |
| 60        | 1h一小时K线 |
| 120       | 2h两小时K线 |
| 240       | 4h四小时K线 |
| 360       | 6h六小时K线 |
| 720       | 12h十二小时K线 |
| 1D        | 1D日K线 |
| 1W        | 1W周K线 |
| 1M        | 1M月K线 |

**返回值说明**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| klines   | 数组     |    | K线数据 |
| open     | float64 |    | 开盘价 |
| close    | float64 |    | 收盘价 |
| high     | float64 |    | 最高价 |
| low      | float64 |    | 最低价 |
| volume   | float64 |    | 交易数量 |
| ts       | int64  |    | k线时间戳，单位毫秒 |

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


### 7. 获取合约的最新价格信息

    获取合约的最新价格信息

**HTTP请求**

  ```http
    # Request
    GET api/v1/market/getLatestPrice
```

**请求参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 是 |合约名称| 合约名称后面需要加下划线(BTC-USDT) |

**返回值说明**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| tradePrice    | float64 |    | 成交价格 |
| indexPrice    | float64 |    | 指数价格 |
| fairPrice     | float64 |    | 标记价格 |

```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
          "tradePrice": "50000.18",
          "indexPrice": "50000.18",
          "fairPrice": "50000.18"
        }
    }
```


### 8. 获取合约当前最新的资金费率

    获取合约当前最新的资金费率

**HTTP请求**

  ```http
    # Request
    GET api/v1/market/getLatestFunding
```

**请求参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 是 |合约名称| 合约名称后面需要加下划线(BTC-USDT) |

**返回值说明**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| fundingRate   | float64 |    | 当前资金费率 |
| fairPrice     | float64 |    | 当前的标记价格 |
| leftSeconds   | float64 |    | 下次结算剩余时间，单位为秒 |

```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
          "fundingRate": "0.3000",
          "fairPrice": "182.90",
          "leftSeconds": "1024",
        }
    }
```



### 9. 查询合约未平仓数量

    查询合约未平仓数量

**HTTP请求**

  ```http
    # Request
    GET api/v1/market/getOpenPositions
```

**请求参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 是 |合约名称| 合约名称后面需要加下划线(BTC-USDT) |

**返回值说明**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| volume  | float64 |    | 持仓数量 |
| unit    | string  |    | 持仓数量对应的单位，CONT(张), BTC, ETH, LINK, BCH等等 |

```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
          "volume": "10.00",
          "unit": "BTC",
        }
    }
```


## 交易API


### 1. 下单交易接口

     下单交易

**HTTP请求**

     
           
```http
    # Request
    POST api/v1/user/trade
```
**请求方式**

    POST

**请求参数**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| symbol | String | 是 | 合约符号(BTC-USDT) |
| apiKey | String | 是 | 接口密钥 |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |
| side | String | 是 | (Bid/Ask 买/卖) |
| entrustPrice | float64 | 是 | 价格  |
| entrustVolume | float64 | 是 | 数量 |
| tradeType | String | 是 | Market/Limit  市价/限价 |
| action | String | 是 | Open/Close 开仓/平仓  |

**返回值说明**

| 参数名 | 参数类型  | 必填 | 描述 |
| ---- |---- | ---- | ---- |
| orderId | String | 是 | 订单ID |

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



### 2. 撤销订单
 
       撤销订单  

**HTTP请求**
 
   
   
```http
    # Request
    POST api/v1/user/cancelOrder
```

**请求方式**

    POST

**请求参数**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| orderId   | String | 是 | 订单ID |
| symbol    | String | 是 | 合约符号(BTC-USDT) |
| apiKey | String | 是 | 接口密钥 |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**
| 参数名 | 参数类型  | 必填 | 描述 |
| ---- |---- | ---- | ---- |
| orderId | String | 是 | 订单ID |


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


### 3. 查询委托中的订单

    查询委托中的订单

**HTTP请求**

```http
    # Request
    POST api/v1/user/pendingOrders
```

**请求方式**

    POST

**请求参数**  

| 参数名 | 参数类型  | 必填 | 字段描述
| ------------- |----|----|----|
| symbol | String | 是 |  合约产品(BTC-USDT)，为空则返回全部 |
| apiKey | String | 是 |  |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |

 **返回值说明**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| entrustTm     | String  | 是 | 订单委托时间 |
| side          | String  | 是 | 交易方向(Bid/Ask 买/卖) |
| tradeType     | String  | 是 | 委托类型(Market/Limit 市价/限价) |
| action        | String  | 是 | Open/Close 开仓/平仓 |
| entrustPrice  | Float64 | 是 | 委托价格 |
| entrustVolume | Float64 | 是 | 委托数量 |
| avgFilledPrice| Float64 | 是 | 成交均价 |
| filledVolume  | Float64 | 是 | 成交数量 |
| orderId       | String  | 是 | 订单号 |

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

  
### 4. 查询单个订单的详情信息
   
    查询单个订单的详情信息

**HTTP请求**
   ```http
    # Request
    POST api/v1/user/queryOrderStatus
```
**请求方式**

    POST

**请求参数**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| apiKey | String | 是 | 接口秘钥 |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |
| symbol | String | 是 | 合约符号(BTC-USDT) |
| orderId | String | 是 | 订单ID |

**返回值说明**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| entrustTm     | String  | 是 | 订单委托时间 |
| side          | String  | 是 | 交易方向(Bid/Ask 买/卖) |
| tradeType     | String  | 是 | 委托类型(Market/Limit 市价/限价) |
| action        | String  | 是 | Open/Close 开仓/平仓 |
| entrustPrice  | Float64 | 是 | 委托价格 |
| entrustVolume | Float64 | 是 | 委托数量 |
| avgFilledPrice| Float64 | 是 | 成交均价 |
| filledVolume  | Float64 | 是 | 成交数量 |
| orderId       | String  | 是 | 订单号 |
| status        | String  | 是 | 订单状态(Filled or PartiallyFilled, Pending, Cancelled, Failed) |
  
**备注**

| Status 字段说明  | |
| ----------|----|
| Pending           | 尚未成交 |
| PartiallyFilled   | 部分成交 |
| Cancelled         | 已撤销 |
| Filled            | 已完成 |
| Failed            | 失败 |

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

### 5. 获取用户账户资产信息

    获取用户账户资产信息
          
**HTTP请求**
             
    ```http
        # Request
        POST api/v1/user/getBalance
    ```
**请求方式**
        POST

**请求参数**

| 参数名 | 参数类型  | 必填 |字段描述 |描述 |
| ------------- |----|----|---|---- |
| apiKey | String | 是 | 接口秘钥 | |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 | |
| currency  | String | 是  | 合约资产 | |

**返回值说明**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| code           | Int64   | 是 | 错误码，0表示成功，不为0表示异常失败 |
| msg            | String  | 是 | 错误信息提示 |
| userId	       | String | 是 | 用户ID |
| currency       | String | 是 | 用户资产 |
| balance        | Float64 | 是 | 资产余额 |
| equity         | Float64 | 是 | 资产净值 |
| unrealisedPNL  | Float64 | 是 | 未实现盈亏 |
| realisedPNL    | Float64 | 是 | 已实现盈亏 |
| availableMargin| Float64 | 是 | 可用保证金 |
| usedMargin     | Float64 | 是 | 已用保证金 |
| freezedMargin  | Float64 | 是 | 冻结保证金 |
| longLeverage   | Float64 | 是 | 做多杠杆 |
| shortLeverage  | Float64 | 是 | 做空杠杆 |

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

### 6. 查询用户持仓信息

    查询用户持仓信息

**HTTP请求**

```http
    # Request
    POST api/v1/user/getPositions
```

**请求方式**

    POST

**请求参数**  

| 参数名 | 参数类型  | 必填 | 字段描述
| ------------- |----|----|----|
| symbol | String | 是 |  合约产品(BTC-USDT)，为空则表示全部都返回 |
| apiKey | String | 是 |  |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| code           | Int64   | 是 | 错误码，0表示成功，不为0表示异常失败 |
| msg            | String  | 是 | 错误信息提示 |
| symbol         | String  | 是 | 合约品种 |
| currency       | String  | 是 | 用户资产 |
| positionSide   | String  | 是 | 仓位方向 Long/Short 多/空 |
| marginMode     | String  | 是 | 保证金模式 Cross/Isolated 全仓/逐仓  |
| volume         | Float64 | 是 | 持仓数量 |
| availableVolume| Float64 | 是 | 可平仓数量 |
| unrealisedPNL  | Float64 | 是 | 未实现盈亏 |
| realisedPNL    | Float64 | 是 | 已实现盈亏 |
| margin         | Float64 | 是 | 保证金 |
| avgPrice       | Float64 | 是 | 开仓均价 |
| liquidatedPrice| Float64 | 是 | 预估强平价 |
| leverage       | Float64 | 是 | 杠杆 |

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

### 7. 批量撤销订单

  批量撤销订单

**HTTP请求**


```http
    # Request
    POST api/v1/user/batchCancelOrders
```

**请求方式**

    POST

**请求参数**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| symbol    | String | 是 | 合约符号(BTC-USDT) |
| apiKey | String | 是 | 接口密钥 |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**
| 参数名 | 参数类型  | 必填 | 描述 |
| ---- |---- | ---- | ---- |
| code    | Int64  | 是 | 错误码，0表示成功，不为0表示异常失败 |
| msg     | String | 是 | 错误信息提示 |


```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
        }
    }
```

### 8. 一键平仓下单

  一键平仓下单

**HTTP请求**


```http
    # Request
    POST api/v1/user/oneClickClosePosition
```

**请求方式**

    POST

**请求参数**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| symbol     | String | 是 | 合约符号(BTC-USDT) |
| positionId | Int64  | 是 | 一键平仓对应的仓位ID |
| apiKey     | String | 是 | 接口密钥 |
| timestamp  | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**
| 参数名 | 参数类型  | 必填 | 描述 |
| ---- |---- | ---- | ---- |
| code    | Int64  | 是 | 错误码，0表示成功，不为0表示异常失败 |
| msg     | String | 是 | 错误信息提示 |
| orderId | String | 是 | 一键平仓产生的委托订单ID |

```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
        }
    }
```


### 9. 全部一键平仓下单

  全部一键平仓下单

**HTTP请求**


```http
    # Request
    POST api/v1/user/oneClickCloseAllPositions
```

**请求方式**

    POST

**请求参数**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| symbol    | String | 是 | 合约符号(BTC-USDT) |
| apiKey | String | 是 | 接口密钥 |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**
| 参数名 | 参数类型  | 必填 | 描述 |
| ---- |---- | ---- | ---- |
| code    | Int64      | 是 | 错误码，0表示成功，不为0表示异常失败 |
| msg     | String     | 是 | 错误信息提示 |
| orders  | String数组  | 是 | 全部一键平仓产生的多个委托订单ID |

```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
          "orders": ["123", "456", "789"]
        }
    }
```

### 10. 撤销全部订单

  撤销全部订单

**HTTP请求**


```http
    # Request
    POST api/v1/user/cancelAll
```

**请求方式**

    POST

**请求参数**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| apiKey | String | 是 | 接口密钥 |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**
| 参数名 | 参数类型  | 必填 | 描述 |
| ---- |---- | ---- | ---- |
| code | Int64  | 是 | 错误码，0表示成功，不为0表示异常失败 |
| msg  | String | 是 | 错误信息提示 |

```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
        }
    }
```


### 11. 修改账户保证金模式

  修改账户保证金模式

**HTTP请求**


```http
    # Request
    POST api/v1/user/setMarginMode
```

**请求方式**

    POST

**请求参数**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| symbol      | String | 是 | 合约符号(BTC-USDT) |
| marginMode  | String | 是 | Isolated or Cross, 账户保证金模式，逐仓或者全仓 |
| apiKey      | String | 是 | 接口密钥 |
| timestamp   | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**
| 参数名 | 参数类型  | 必填 | 描述 |
| ---- |---- | ---- | ---- |
| code | Int64  | 是 | 错误码，0表示成功，不为0表示异常失败 |
| msg  | String | 是 | 错误信息提示 |


```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
        }
    }
```

### 12. 修改杠杆

  修改杠杆

**HTTP请求**


```http
    # Request
    POST api/v1/user/setLeverage
```

**请求方式**

    POST

**请求参数**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| symbol    | String | 是 | 合约符号(BTC-USDT) |
| side      | String | 是 | 多仓或者空仓的杠杆，Long表示多仓，Short表示空仓 |
| leverage  | String | 是 | 杠杆倍数 |
| apiKey    | String | 是 | 接口密钥 |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**
| 参数名 | 参数类型  | 必填 | 描述 |
| ---- |---- | ---- | ---- |
| code | Int64  | 是 | 错误码，0表示成功，不为0表示异常失败 |
| msg  | String | 是 | 错误信息提示 |


```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
        }
    }
```

## 基础信息API

### 1. 获取系统时间

    获取系统时间

**HTTP请求**

```http
    # Request
    POST api/v1/common/server/time
```

**请求方式**

    GET / POST


**请求参数**

    无

**返回值说明**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| code        | Int64  | 是 | 错误码，0表示成功，不为0表示异常失败 |
| msg         | String | 是 | 错误信息提示 |
| currentTime | Int64  | 是 | 系统当前时间，单位毫秒 |

```javascript
    # Response
    {
        "code": 0,
        "msg": "",
        "currentTime": 1534431933321
    }
```

**备注**
    
[Bingbon]: https://bingbon.pro
[English Docs]: https://bingbon.pro
[Unix Epoch]: https://en.wikipedia.org/wiki/Unix_time
