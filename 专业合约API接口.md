Bingbon交易所官方API文档
==================================================
[Bingbon][]交易所开发者文档([English Docs][])。

<!-- TOC -->
- [Bingbon交易所官方API文档](#bingbon交易所官方api文档)
- [介绍](#介绍)
- [签名认证](#签名认证)
  - [创建API](#创建api)
  - [请求内容](#请求内容)
  - [签名说明](#签名说明)
  - [请求交互](#请求交互)
- [基础信息](#基础信息)
  - [常见错误码](#常见错误码)
  - [时间戳规范](#时间戳规范)
    - [例子](#例子)
  - [数字规范](#数字规范)
  - [频率限制](#频率限制)
    - [REST API](#rest-api)
  - [查询系统时间](#查询系统时间)
- [行情接口](#行情接口)
  - [1. 查询合约基础信息](#1-查询合约基础信息)
  - [2. 查询合约最新价格](#2-查询合约最新价格)
  - [3. 查询合约交易深度](#3-查询合约交易深度)
  - [4. 查询最新成交记录](#4-查询最新成交记录)
  - [5. 查询最新资金费率](#5-查询最新资金费率)
  - [6. 查询资金费率历史](#6-查询资金费率历史)
  - [7. 查询K线最新数据](#7-查询k线最新数据)
  - [8. 查询K线历史数据](#8-查询k线历史数据)
  - [9. 查询合约未平仓数量](#9-查询合约未平仓数量)
  - [10. 查询合约24小时价格变动情况](#10-查询合约24小时价格变动情况)
- [账户接口](#账户接口)
  - [1. 查询账户信息](#1-查询账户信息)
  - [2. 查询持仓信息](#2-查询持仓信息)
- [交易API](#交易api)
  - [1. 交易下单](#1-交易下单)
  - [2. 一键平仓下单](#2-一键平仓下单)
  - [3. 全部一键平仓下单](#3-全部一键平仓下单)
  - [4. 撤销订单](#4-撤销订单)
  - [5. 批量撤销订单](#5-批量撤销订单)
  - [6. 撤销全部订单](#6-撤销全部订单)
  - [7. 查询当前委托订单](#7-查询当前委托订单)
  - [8. 查询订单详情](#8-查询订单详情)
  - [9. 查询账户保证金模式](#9-查询账户保证金模式)
  - [10. 切换账户保证金模式](#10-切换账户保证金模式)
  - [11. 查询杠杆](#11-查询杠杆)
  - [12. 修改杠杆](#12-修改杠杆)
  - [13. 查询强平订单](#13-查询强平订单)
  - [14. 查询历史订单](#14-查询历史订单)
  - [15. 设置止盈止损订单](#15-设置止盈止损订单)
  - [16. 撤销止盈止损订单](#16-撤销止盈止损订单)
  - [17. 查询止盈止损订单列表](#17-查询止盈止损订单列表)
  - [18. 查询止盈止损订单列表](#18-查询止盈止损历史订单列表)


<!-- /TOC -->

# 介绍

欢迎使用[Bingbon][]开发者文档。

本文档提供了专业合约交易业务的账户管理、行情查询、交易功能等相关API的使用方法介绍。
行情API提供市场的公开的行情数据接口，账户和交易API需要身份验证，提供下单、撤单，查询订单和帐户信息等功能。

# 签名认证
## 创建API

在对任何请求进行签名之前，您必须通过 Bingbon 网站【用户中心】-【API管理(专业合约)】创建一个API key。 创建key后，您将获得2个必须记住的信息：
* API Key
* Secret Key

API Key 和 Secret Key将由随机生成和提供。

## 请求内容

所有REST请求都必须包含以下参数：

* API Key作为一个字符串。
* sign 使用一定算法得出的签名（请参阅签名信息）。
* timestamp 作为您的请求的时间戳。
* 所有请求都应该含有application/json类型内容，并且是有效的JSON。

## 签名说明
sign 是对http method，url path，请求参数等按字符串连接之后使用 **HMAC SHA256** 方法加密而得到的。

* path 为URL的请求路径，例如: /api/v1/user/getBalance
* method 是请求方法(POST/GET/PUT/DELETE)，字母全部大写。
* 参数是对 **所有参数(含timestamp)按照字典排序之后，按照key1=value1 + key2=value2 ... + Secret Key** 字符串(+表示字符串连接)。

originString = method + path + params
sign = HmacSHA256(originString)

**例如：对于如下的请求参数进行签名**

```bash
curl "https://api-swap-rest.bingbon.pro/api/v1/user/getBalance"
      
```
* 通过POST方式获取获取用户某资产余额信息，以 apiKey=Zsm4DcrHBTewmVaElrdwA67PmivPv6VDK6JAkiECZ9QfcUnmn67qjCOgvRuZVOzU, secretKey=UuGuyEGt6ZEkpUObCYCmIfh0elYsZVh80jlYwpJuRZEw70t6vomMH7Sjmf94ztSI 为例
```
timestamp = 1616488398013
apiKey = Zsm4DcrHBTewmVaElrdwA67PmivPv6VDK6JAkiECZ9QfcUnmn67qjCOgvRuZVOzU
currency = USDT
```

请求参数按字典排序之后，为
```
apiKey = Zsm4DcrHBTewmVaElrdwA67PmivPv6VDK6JAkiECZ9QfcUnmn67qjCOgvRuZVOzU
currency = USDT
timestamp = 1616488398013
```

mothod为POST，path为/api/v1/user/getBalance，生成待签名的参数字符串如下:

```
paramString = 'apiKey=Zsm4DcrHBTewmVaElrdwA67PmivPv6VDK6JAkiECZ9QfcUnmn67qjCOgvRuZVOzU&currency=USDT&timestamp=1616488398013'

```

按算法生成待签名的字符串

```
originString = 'POST/api/v1/user/getBalanceapiKey=Zsm4DcrHBTewmVaElrdwA67PmivPv6VDK6JAkiECZ9QfcUnmn67qjCOgvRuZVOzU&currency=USDT&timestamp=1616488398013'
  
```

然后，将待签名字符串添加私钥参数生成最终待签名字符串。


例如：
```
Signature = HmacSHA256(secretkey, originString)
Signature = Base64Encode(Signature)
Signature = UrlEncode(Signature)

即：
Signature = HmacSHA256("UuGuyEGt6ZEkpUObCYCmIfh0elYsZVh80jlYwpJuRZEw70t6vomMH7Sjmf94ztSI", "POST/api/v1/user/getBalanceapiKey=Zsm4DcrHBTewmVaElrdwA67PmivPv6VDK6JAkiECZ9QfcUnmn67qjCOgvRuZVOzU&currency=USDT&timestamp=1616488398013")

echo -n "POST/api/v1/user/getBalanceapiKey=Zsm4DcrHBTewmVaElrdwA67PmivPv6VDK6JAkiECZ9QfcUnmn67qjCOgvRuZVOzU&currency=USDT&timestamp=1616488398013" | openssl dgst -sha256 -hmac "UuGuyEGt6ZEkpUObCYCmIfh0elYsZVh80jlYwpJuRZEw70t6vomMH7Sjmf94ztSI" -binary | base64 | xargs python2.7 -c 'import sys, urllib;print(urllib.quote(sys.argv[1]))'

```
Signature的结果为S7Ok3L5ROXSbYfXj9ryeBbKfRosh9tmH%2FAKiwj7eAoc%3D，则签名之后的url query参数为
```
apiKey = Zsm4DcrHBTewmVaElrdwA67PmivPv6VDK6JAkiECZ9QfcUnmn67qjCOgvRuZVOzU
currency = USDT
timestamp = 1616488398013
sign = S7Ok3L5ROXSbYfXj9ryeBbKfRosh9tmH%2FAKiwj7eAoc%3D

即最终发送给服务器的API请求应该为：
"https://api-swap-rest.bingbon.pro/api/v1/user/getBalance?apiKey=Zsm4DcrHBTewmVaElrdwA67PmivPv6VDK6JAkiECZ9QfcUnmn67qjCOgvRuZVOzU&currency=USDT&timestamp=1616488398013&sign=S7Ok3L5ROXSbYfXj9ryeBbKfRosh9tmH%2FAKiwj7eAoc%3D"

```

## 请求交互  

REST访问的根URL：`https://api-swap-rest.bingbon.pro`

所有请求基于Https协议，请求头信息中Content-Type 需要统一设置为:'application/json’。

**请求交互说明**

1、请求参数：根据接口请求参数规定进行参数封装。

2、提交请求参数：将封装好的请求参数通过POST/GET/DELETE等方式提交至服务器。

3、服务器响应：服务器首先对用户请求数据进行参数安全校验，通过校验后根据业务逻辑将响应数据以JSON格式返回给用户。

4、数据处理：对服务器响应数据进行处理。

**成功**

HTTP状态码200表示成功响应，并可能包含内容。如果响应含有内容，则将显示在相应的返回内容里面。

# 基础信息
## 常见错误码

**常见HTTP错误码**

###类型:
* 4XX 错误码用于指示错误的请求内容、行为、格式

* 5XX 错误码用于指示Bingbon服务侧的问题

###错误码:
* 400 Bad Request – Invalid request format 请求格式无效

* 401 Unauthorized – Invalid API Key 无效的API Key

* 403 Forbidden – You do not have access to the requested resource 请求无权限

* 404 - Not Found 没有找到请求

* 429 - Too Many Requests 请求太频繁被系统限流

* 418 - 表示收到429后继续访问，于是被封了

* 500 - Internal Server Error – We had a problem with our server 服务器内部错误

* 504 - 表示API服务端已经向业务核心提交了请求但未能获取响应(特别需要注意的是504代码不代表请求失败，而是未知。很可能已经得到了执行，也有可能执行失败，需要做进一步确认)

###注意:
* 如果失败，response body 带有错误描述信息

* 每个接口都有可能抛出异常


## 时间戳规范

除非另外指定，API中的所有时间戳均以微秒为单位返回。

请求的时间戳必须在API服务时间的30秒内，否则请求将被视为过期并被拒绝。如果本地服务器时间和API服务器时间之间存在较大的偏差，那么我们建议您使用通过查询API服务器时间来更新http header。

### 例子

1587091154123

## 数字规范

为了保持跨平台时精度的完整性，十进制数字作为字符串返回。建议您在发起请求时也将数字转换为字符串以避免截断和精度错误。 

整数（如交易编号和顺序）不加引号。

## 频率限制

如果请求过于频繁系统将自动限制请求。

### REST API

* 行情接口：通过IP限制公共接口的调用，每1秒最多60个请求。

* 账户和交易接口：通过用户ID限制私人接口的调用，每1秒最多10个请求。

* 某些接口的特殊限制在具体的接口上注明

## 查询系统时间

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

| 参数名 | 参数类型 | 描述 |
| ------------- |----|----|
| code        | Int64  | 错误码，0表示成功，不为0表示异常失败 |
| msg         | String | 错误信息提示 |
| currentTime | Int64  | 系统当前时间，单位毫秒 |

```javascript
    # Response
    {
        "code": 0,
        "msg": "",
        "currentTime": 1534431933321
    }
```

# 行情接口

## 1. 查询合约基础信息     

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
| maxLongLeverage   | 多头交易的最大杠杆倍数 |
| maxShortLeverage  | 空头交易的最大杠杆倍数 |

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
            "asset": "BTC",
            "maxLongLeverage": 100,
            "maxShortLeverage": 100
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
            "asset": "ETH",
            "maxLongLeverage": 50,
            "maxShortLeverage": 50
        }],
        ...
   } 
```

## 2. 查询合约最新价格

**HTTP请求**

```http
    # Request
    GET api/v1/market/getLatestPrice
```

**请求参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 是 |合约名称| 合约名称中需有"-"，如BTC-USDT |

**返回值说明**

| 参数名 | 参数类型  | 描述 |
| ------------- |----|----|
| tradePrice    | float64 | 成交价格 |
| indexPrice    | float64 | 指数价格 |
| fairPrice     | float64 | 标记价格 |

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
## 3. 查询合约交易深度

**HTTP请求**

```http
    # Request
    GET api/v1/market/getMarketDepth
```

**请求参数**

| 参数名 | 参数类型  | 必填 | 字段描述 | 描述 |
| ------------- |----|----|----| --- |
| symbol | String | 是 | 合约名称 | 合约名称中需有"-"，如BTC-USDT |
| level | String | 否 | 层数 | 若为空，则默认返回5层 |

**返回值说明**

| 参数名 | 参数类型  | 描述 |
| ------------- |----|----|
| code   | Int64 | 是否有错误信息，0为正常，1为有错误 |
| msg    | String | 错误信息描述 |
| asks   | 数组 | 卖方深度 |
| bids   | 数组 | 买方深度 |
| p      | float64 | price价格  | 
| v      | float64 | volume数量 | 


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

## 4. 查询最新成交记录

**HTTP请求**

```http
    # Request
    GET api/v1/market/getMarketTrades
```

**请求参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 是 | 合约名称 | 合约名称中需有"-"，如BTC-USDT |

**返回值说明**

| 参数名 | 参数类型  | 描述 |
| ------------- |----|----|
| time      | data   | 成交时间 |
| makerSide | String | 吃单方向(Buy / Sell 买/卖) |
| price     | String | 成交价格 |
| volume    | String | 成交数量 |

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

## 5. 查询最新资金费率

**HTTP请求**

```http
    # Request
    GET api/v1/market/getLatestFunding
```

**请求参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 是 | 合约名称 | 合约名称中需有"-"，如BTC-USDT |

**返回值说明**

| 参数名 | 参数类型  | 描述 |
| ------------- |----|----|
| fundingRate   | float64 | 当前资金费率 |
| fairPrice     | float64 | 当前的标记价格 |
| leftSeconds   | float64 | 下次结算剩余时间，单位为秒 |

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

## 6. 查询资金费率历史

**HTTP请求**

```http
    # Request
    GET api/v1/market/getHistoryFunding
```

**请求参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 是 | 合约名称 | 合约名称中需有"-"，如BTC-USDT |

**返回值说明**

| 参数名 | 参数类型  | 描述 |
| ------------- |----|----|
| historyId     | String | 历史ID号 |
| fundingRate   | String | 资金费率 |
| fairPrice     | String | 标记价格 |
| interval      | String | 资金费率结算周期，单位：小时 |
| time          | data   | 结算时间 |

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

## 7. 查询K线最新数据

     查询最新成交价格的K线最新数据。

**HTTP请求**

```http
    # Request
    GET api/v1/market/getLatestKline
```

**请求参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 是 | 合约名称 | 合约名称中需有"-"，如BTC-USDT |
| klineType | String | 是 | k线类型 | 参考字段说明，如分钟，小时，周等 |

**备注**

| klineType 字段说明  | |
| ----------|----|
| 1            | 1m一分钟K线 |
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

| 参数名 | 参数类型  | 描述 |
| ------------- |----|----|
| open     | float64 | 开盘价 |
| close    | float64 | 收盘价 |
| high     | float64 | 最高价 |
| low      | float64 | 最低价 |
| volume   | float64 | 交易数量 |
| ts       | int64  | k线时间戳，单位毫秒 |

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


## 8. 查询K线历史数据

     查询一段时间周期内成交价格的K线历史数据。

**HTTP请求**

```http
    # Request
    GET api/v1/market/getHistoryKlines
```

**请求参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 是 | 合约名称 | 合约名称中需有"-"，如BTC-USDT |
| klineType | String | 是 | k线类型 | 参考字段说明，如分钟，小时，周等 |
| startTs       | int64  | 是 | 起始时间戳，单位毫秒 |
| endTs       | int64  | 是 | 结束时间戳，单位毫秒 |

**备注**

| klineType 字段说明  | |
| ----------|----|
| 1         | 1m一分钟K线 |
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

| 参数名 | 参数类型  | 描述 |
| ------------- |----|----|
| klines   | 数组     | K线数据 |
| open     | float64 | 开盘价 |
| close    | float64 | 收盘价 |
| high     | float64 | 最高价 |
| low      | float64 | 最低价 |
| volume   | float64 | 交易数量 |
| ts       | int64  | k线时间戳，单位毫秒 |

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

## 9. 查询合约未平仓数量

**HTTP请求**

```http
    # Request
    GET api/v1/market/getOpenPositions
```

**请求参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 是 |合约名称| 合约名称中需有"-"，如BTC-USDT |

**返回值说明**

| 参数名 | 参数类型  | 描述 |
| ------------- |----|----|
| volume  | float64 | 持仓数量 |
| unit    | string  | 持仓数量对应的单位，CONT(张), BTC, ETH, LINK, BCH等等 |

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

## 10. 查询合约24小时价格变动情况

**HTTP请求**

```http
    # Request
    GET api/v1/market/getTicker
```

**请求参数**

| 参数名  | 参数类型  | 必填 | 字段描述 | 描述 |
| -------|--------|--- |-------|------|
| symbol | String | 否 |合约名称| 合约名称中需有"-"，如BTC-USDT |

**返回值说明**

| 参数名 | 参数类型  | 描述 |
| ------------- |----|----|
| symbol  | String | 合约名称 |
| priceChange    | String  | 价格变动, 单位是USDT |
| priceChangePercent  | String | 价格变动百分比 |
| lastPrice    | String  | 最新交易价格 |
| lastVolume  | String |  最新交易数量 |
| highPrice    | String  | 24小时最高价 |
| lowPrice  | String | 24小时最低价 |
| volume    | String  | 24小时成交量 |
| dayVolume  | String | 24小时成交额, 单位是USDT |
| openPrice    | String  | 24小时内第一个价格 |

```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
          "symbol": "BTC-USDT",
          "priceChange": "10.00",
          "priceChangePercent": "10",
          "lastPrice": "5738.23",
          "lastVolume": "31.21",
          "highPrice": "5938.23",
          "lowPrice": "5238.23",
          "volume": "23211231.13",
          "dayVolume": "213124412412.47",
          "openPrice": "5828.32"
        }
    }
```


# 账户接口

## 1. 查询账户信息

    查询当前账户下专业合约资产的相关信息。

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

| 参数名 | 参数类型  | 描述 |
| ------------- |----|----|
| code           | Int64   | 错误码，0表示成功，不为0表示异常失败 |
| msg            | String  | 错误信息提示 |
| userId           | String | 用户ID |
| currency       | String | 用户资产 |
| balance        | Float64 | 资产余额 |
| equity         | Float64 | 资产净值 |
| unrealisedPNL  | Float64 | 未实现盈亏 |
| realisedPNL    | Float64 | 已实现盈亏 |
| availableMargin| Float64 | 可用保证金 |
| usedMargin     | Float64 | 已用保证金 |
| freezedMargin  | Float64 | 冻结保证金 |
| longLeverage   | Float64 | 做多杠杆 |
| shortLeverage  | Float64 | 做空杠杆 |

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

## 2. 查询持仓信息

    查询当前账户下专业合约的持仓信息与盈亏情况。

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
| symbol | String | 是 |  合约名称中需有"-"，如BTC-USDT，为空则表示全部都返回 |
| apiKey | String | 是 |  |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**

| 参数名 | 参数类型  | 描述 |
| ------------- |----|----|
| code           | Int64   | 错误码，0表示成功，不为0表示异常失败 |
| msg            | String  | 错误信息提示 |
| symbol         | String  | 合约品种 |
| currency       | String  | 资产类型 |
| positionId     | String  | 仓位ID  |
| positionSide   | String  | 仓位方向 Long/Short 多/空 |
| marginMode     | String  | 保证金模式 Cross/Isolated 全仓/逐仓  |
| volume         | Float64 | 持仓数量 |
| availableVolume| Float64 | 可平仓数量 |
| unrealisedPNL  | Float64 | 未实现盈亏 |
| realisedPNL    | Float64 | 已实现盈亏 |
| margin         | Float64 | 保证金 |
| avgPrice       | Float64 | 开仓均价 |
| liquidatedPrice| Float64 | 预估强平价 |
| leverage       | Float64 | 杠杆 |

```javascript

# Response
    {
       "code": 0,
       "msg": "",
       "data": {
            "positions": [
                {
                    "symbol": "BTC-USDT",
                    "positionId": "12345678",
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

# 交易API

## 1. 交易下单

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
| symbol | String | 是 | 合约名称中需有"-"，如BTC-USDT |
| apiKey | String | 是 | 接口密钥 |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |
| side | String | 是 | (Bid/Ask 买/卖) |
| entrustPrice | float64 | 是 | 价格  |
| entrustVolume | float64 | 是 | 数量 |
| tradeType | String | 是 | Market/Limit  市价/限价 |
| action | String | 是 | Open/Close 开仓/平仓  |
| takerProfitPrice | float64 | 否 | 止盈价格 |
| stopLossPrice | float64 | 否 | 止损价格 |

**返回值说明**

| 参数名 | 参数类型  | 描述 |
| ---- |---- | ---- |
| orderId | String | 订单ID |

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

## 2. 一键平仓下单

    查询持仓信息后，可根据仓位ID进行一键平仓操作。注意，一键平仓是以市价委托进行触发的。

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
| symbol     | String | 是 | 合约名称中需有"-"，如BTC-USDT |
| positionId | Int64  | 是 | 一键平仓对应的仓位ID |
| apiKey     | String | 是 | 接口密钥 |
| timestamp  | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**
| 参数名 | 参数类型  | 描述 |
| ---- |---- | ---- |
| code    | Int64  | 错误码，0表示成功，不为0表示异常失败 |
| msg     | String | 错误信息提示 |
| orderId | String | 一键平仓产生的委托订单ID |

```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
        }
    }
```

## 3. 全部一键平仓下单

    将当前账户下所有仓位进行一键平仓操作。注意，一键平仓是以市价委托进行触发的。

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
| apiKey | String | 是 | 接口密钥 |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**
| 参数名 | 参数类型  | 描述 |
| ---- |---- | ---- |
| code    | Int64      | 错误码，0表示成功，不为0表示异常失败 |
| msg     | String     | 错误信息提示 |
| orders  | String数组  | 全部一键平仓产生的多个委托订单ID |

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

## 4. 撤销订单

    将处于当前委托状态的订单进行撤销操作。

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
| symbol    | String | 是 | 合约名称中需有"-"，如BTC-USDT |
| apiKey | String | 是 | 接口密钥 |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**
| 参数名 | 参数类型  | 描述 |
| ---- |---- | ---- |
| orderId | String | 订单ID |

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

## 5. 批量撤销订单

    将合约下处于当前委托状态的部分订单进行撤销操作。

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
| symbol    | String | 是 | 合约名称中需有"-"，如BTC-USDT |
| oids      | String | 是 | 订单ID列表，多个订单id用逗号进行分隔 |
| apiKey    | String | 是 | 接口密钥 |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**
| 参数名 | 参数类型  | 描述 |
| ---- |---- | ---- |
| code          | Int64  | 错误码，0表示成功，不为0表示异常失败 |
| msg           | String | 错误信息提示 |
| Success       | String数组 | 撤销成功的订单ID列表 |
| Failed        | 结构数组 | 撤销失败的订单列表 |
| orderId       | String | 订单ID |
| errorCode     | Int64  | 错误码，0表示成功，不为0表示异常失败 |
| errorMessage  | String | 错误信息提示 |



```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
          "success": ["725970815","725970736"],
          "failed":[
            {
              "orderId": "725971356",
              "errorCode": 80012,
              "errorMessage": "Service network failed"
            },
          ],
        }
    }
```

## 6. 撤销全部订单

    将账户下处于当前委托状态的全部订单进行撤销操作。

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
| 参数名 | 参数类型  | 描述 |
| ---- |---- | ---- |
| code | Int64  | 错误码，0表示成功，不为0表示异常失败 |
| msg  | String | 错误信息提示 |

```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
        }
    }
```

## 7. 查询当前委托订单

    查询一段时间周期内账户下处于当前委托状态的订单详情。

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
| symbol | String | 是 |  合约名称中需有"-"，如BTC-USDT，为空则返回全部 |
| apiKey | String | 是 |  |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |

 **返回值说明**

| 参数名 | 参数类型  | 描述 |
| ------------- |----|----|
| entrustTm     | String  | 订单委托时间 |
| side          | String  | 交易方向(Bid/Ask 买/卖) |
| tradeType     | String  | 委托类型(Market/Limit 市价/限价) |
| action        | String  | Open/Close 开仓/平仓 |
| entrustPrice  | Float64 | 委托价格 |
| entrustVolume | Float64 | 委托数量 |
| avgFilledPrice| Float64 | 成交均价 |
| filledVolume  | Float64 | 成交数量 |
| orderId       | String  | 订单号 |
| profit        | Float64 | 盈亏 |
| commission    | Float64 | 手续费 |
| updateTm      | String  | 订单更新时间 |

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
                    "profit": 0,
                    "commission": 0,
                    "updateTm": "2018-04-25T15:00:52.000Z"
                }
            ]
        }
    }
 ```

## 8. 查询订单详情

   
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
| symbol | String | 是 | 合约名称中需有"-"，如BTC-USDT |
| orderId | String | 是 | 订单ID |

**返回值说明**

| 参数名 | 参数类型  | 描述 |
| ------------- |----|----|
| entrustTm     | String  | 订单委托时间 |
| side          | String  | 交易方向(Bid/Ask 买/卖) |
| tradeType     | String  | 委托类型(Market/Limit 市价/限价) |
| action        | String  | Open/Close 开仓/平仓 |
| entrustPrice  | Float64 | 委托价格 |
| entrustVolume | Float64 | 委托数量 |
| avgFilledPrice| Float64 | 成交均价 |
| filledVolume  | Float64 | 成交数量 |
| orderId       | String  | 订单号 |
| status        | String  | 订单状态(Filled or PartiallyFilled, Pending, Cancelled, Failed) |
| profit        | Float64 | 盈亏 |
| commission    | Float64 | 手续费 |
| updateTm      | String  | 订单更新时间 |
  
**备注**

| Status  | 字段说明 |
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
            "status": "Filled",
            "profit": 0,
            "commission": 0,
            "updateTm": "2018-04-25T15:00:52.000Z"
        }
    }
```

## 9. 查询账户保证金模式

**HTTP请求**

```http
    # Request
    Post api/v1/user/getMarginMode
```

**请求方式**

    POST

**请求参数**

| 参数名     | 参数类型 | 必填 | 描述                        |
| --------- | ------- | --- | -------------------------- |
| symbol    | String  | 是   | 合约名称中需有"-"，如BTC-USDT |
| apiKey    | String  | 是   | 接口密钥 |
| timestamp | String  | 是   | 发起请求的时间戳，单位为毫秒 |

**返回值说明**

| 参数名       | 参数类型 | 描述     |
| ----------- | ------ | -------- |
| marginMode  | String | 保证金模式 |

**备注**

| marginMode | 字段说明 |
| ----------|----|
| Isolated | 逐仓 |
| Cross    | 全仓 |

```javascript
# Response
    {
        "data": {
            "marginMode":"Isolated"
        },
        "code": 0,
        "message": ""
    }
```

## 10. 切换账户保证金模式

  修改专业合约账户的保证金模式，全仓模式或逐仓模式。

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
| symbol      | String | 是 | 合约名称中需有"-"，如BTC-USDT |
| marginMode  | String | 是 | Isolated or Cross, 账户保证金模式，逐仓或者全仓 |
| apiKey      | String | 是 | 接口密钥 |
| timestamp   | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**

| 参数名 | 参数类型 | 描述 |
| ---- |---- | ---- |
| code | Int64  | 错误码，0表示成功，不为0表示异常失败 |
| msg  | String | 错误信息提示 |


```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
        }
    }
```

## 11. 查询杠杆

**HTTP请求**

```http
    # Request
    Post api/v1/user/getLeverage
```

**请求方式**

    POST

**请求参数**

| 参数名     | 参数类型 | 必填 | 描述                        |
| --------- | ------- | --- | -------------------------- |
| symbol    | String  | 是   | 合约名称中需有"-"，如BTC-USDT |
| apiKey    | String  | 是   | 接口密钥 |
| timestamp | String  | 是   | 发起请求的时间戳，单位为毫秒 |

**返回值说明**

| 参数名         | 参数类型 | 描述       |
| ------------- | ------ | ---------- |
| longLeverage  | Int64  | 多仓杠杆倍数 |
| shortLeverage | Int64  | 空仓杠杆倍数 |

```javascript
# Response
    {
        "data": {
            "longLeverage": 5,
            "shortLeverage": 5
        },
        "code": 0,
        "message": ""
    }
```

## 12. 修改杠杆

    调整合约多仓或空仓的杠杆倍数。

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
| symbol    | String | 是 | 合约名称中需有"-"，如BTC-USDT |
| side      | String | 是 | 多仓或者空仓的杠杆，Long表示多仓，Short表示空仓 |
| leverage  | String | 是 | 杠杆倍数 |
| apiKey    | String | 是 | 接口密钥 |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**

| 参数名 | 参数类型  | 描述 |
| ---- |---- | ---- |
| code | Int64  | 错误码，0表示成功，不为0表示异常失败 |
| msg  | String | 错误信息提示 |


```javascript
# Response
    {
        "code": 0,
        "msg": ""
    }
```

## 13. 查询强平订单


**HTTP请求**

```http
    # Request
    POST api/v1/user/forceOrders
```

**请求方式**

    POST

**请求参数**

| 参数名 | 参数类型  | 必填 | 描述 |
| ------------- |----|----|----|
| symbol    | String | 是 | 合约名称中需有"-"，如BTC-USDT |
| autoCloseType  | String | 是 | Liquidation 表示强平订单, ADL 表示减仓订单 |
| lastOrderId  | int64 | 是 | 用于分页, 第一次填写0, 后续填写前一次返回结果里面的最后一个订单id |
| length    | int64 | 是 | 每次请求的长度, 最大值为100 |
| apiKey    | String | 是 | 接口密钥 |
| timestamp | String | 是 | 发起请求的时间戳，单位为毫秒 |

**返回值说明**

| 参数名 | 参数类型  | 描述 |
| ---- |---- | ---- |
| symbol | String  | 合约名称 |
| tradeType  | String | 订单类型, Limit是限价单, Market是市价单 |
| action | String  | Liquidation 表示强平订单, ADL 表示减仓订单 |
| avgFilledPrice  | Float64 | 破产价格 |
| entrustTm  | String | 成交时间 |
| filledVolume  | Float64 | 成交数量 |
| orderId | String  | 订单id |
| side  | String | 交易方向, Bid买入, Ask卖出 |
| profit | Float64  | 盈亏 |    
| commission  | Float64 | 手续费 |


```javascript
# Response
    {
        "code": 0,
        "msg": "",
        "data": {
            "symbol": "BTC-USDT",
            "tradeType": "Limit",
            "action": "Liquidation",
            "avgFilledPrice": 5938.23,
            "entrustTm": "2018-04-25T15:00:51.000Z",
            "filledVolume": 1.2123,
            "orderId": 123456789,
            "side": "Bid",
            "profit": -11.34,
            "commission": 0.4231
        }
    }
```

## 14. 查询历史订单

**HTTP请求**

```http
    # Request
    POST api/v1/user/historyOrders
```

**请求方式**

    POST

**请求参数**

| 参数名         | 参数类型  | 必填 | 描述    |
| ------------- |--------- |-----|--------|
| symbol        | String   | 是  | 合约名称 |
| lastOrderId   | int64    | 是  | 用于分页, 第一次填写0, 后续填写前一次返回结果里面的最后一个订单id |
| length        | int64    | 是  | 每次请求的长度, 最大值为100 |
| apiKey        | String   | 是  | 接口密钥 |
| timestamp     | String   | 是  | 发起请求的时间戳，单位为毫秒 |

**返回值说明**

| 参数名           | 参数类型  | 描述    |
| --------------- | -------- | ------ |
| symbol          | String   | 合约名称 |
| orderId         | String   | 订单id |
| side            | String   | 交易方向, Bid买入, Ask卖出 |
| action          | String   | Open表示开仓, Close表示平仓, ADL表示自动减仓, Liquidation表示爆仓强平 |
| tradeType       | String   | 订单类型, Limit是限价单, Market是市价单 |
| entrustVolume   | Float64  | 委托数量 |
| entrustPrice    | Float64  | 委托价格 |
| filledVolume    | Float64  | 成交数量 |
| avgFilledPrice  | Float64  | 成交价格 |
| profit          | Float64  | 盈亏 |
| commission      | Float64  | 手续费 |    
| orderStatus     | String   | 订单状态(Filled or PartiallyFilled, Pending, Cancelled, Failed) |
| entrustTm       | String   | 委托时间 |
| updateTm        | String   | 更新时间 |

**备注**

| OrderStatus  | 字段说明 |
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
        "data": {
            "orders": [
                {
                    "action": "Open",
                    "avgFilledPrice": 31333.37,
                    "commission": -0.0009,
                    "entrustPrice": 31331.25,
                    "entrustTm": "2021-01-05T09:15:02Z",
                    "entrustVolume": 0.0001,
                    "filledVolume": 0.0001,
                    "orderId": "996273190",
                    "orderStatus": "Filled",
                    "profit": 0,
                    "side": "Bid",
                    "symbol": "BTC-USDT",
                    "tradeType": "Market",
                    "updateTm": "2021-01-05T09:15:15Z"
                }
            ]
        },
        "message": ""
    }
```

## 15. 设置止盈止损订单

**HTTP请求**

```http
    # Request
    POST api/v1/user/stopOrder
```

**请求方式**

    POST

**请求参数**

| 参数名           | 参数类型  | 必填 | 描述    |
| --------------- |---------|------|--------|
| apiKey          | String  | 是   | 接口密钥 |
| timestamp       | String  | 是   | 发起请求的时间戳，单位为毫秒 |
| positionId      | String  | 是   | 持仓id |
| orderId         | String  | 否   | 止盈止损订单id |
| stopLossPrice   | Float64 | 否   | 止损价格 |
| takeProfitPrice | Float64 | 否   | 止盈价格 |
| entrustVolume   | Float64 | 是   | 委托数量 |

**返回值说明**

| 参数名   | 参数类型  | 描述          |
| ------- |--------  | ------------ |
| orderId | String   | 止盈止损订单id |

```javascript
# Response
    {
        "code": 0,
        "data": {
            "orderId": "1414483504200159232"
        },
        "message": ""
    }
```

## 16. 撤销止盈止损订单

**HTTP请求**

```http
    # Request
    POST api/v1/user/cancelStopOrder
```

**请求方式**

    POST

**请求参数**

| 参数名           | 参数类型  | 必填 | 描述    |
| --------------- |---------|-----|---------|
| apiKey          | String  | 是   | 接口密钥 |
| timestamp       | String  | 是   | 发起请求的时间戳，单位为毫秒 |
| orderId         | String  | 是   | 止盈止损订单id |

**返回值说明**

```javascript
# Response
    {
        "code": 0,
        "message": ""
    }
```

## 17. 查询止盈止损订单列表

**HTTP请求**

```http
    # Request
    POST api/v1/user/pendingStopOrders
```

**请求方式**

    POST

**请求参数**

| 参数名           | 参数类型  | 必填 | 描述    |
| --------------- |---------|------|--------|
| apiKey          | String  | 是   | 接口密钥 |
| timestamp       | String  | 是   | 发起请求的时间戳，单位为毫秒 |
| symbol          | String  | 是   | 合约名称 |

**返回值说明**

| 参数名           | 参数类型  | 描述  |
| --------------- |---------| ----- |
| userId          | String  | 用户id |
| orderId         | String  | 订单id |
| symbol          | String  | 合约名称 |
| positionId      | String  | 仓位id |
| stopLossPrice   | Float64 | 止损价格 |
| takeProfitPrice | Float64 | 止盈价格 |
| entrustVolume   | Float64 | 委托数量 |
| side            | String  | 交易方向, Bid买入, Ask卖出 |
| entrustTm       | String  | 委托时间 |

```javascript
# Response
    {
        "code": 0,
        "data": {
            "orders": [
                {
                    "entrustTm": "2021-07-12T07:15:21.891Z",
                    "entrustVolume": 0.001,
                    "orderId": "1414483504200159232",
                    "positionId": "1414483266773192704",
                    "side": "Ask",
                    "stopLossPrice": 10000,
                    "symbol": "BTC-USDT",
                    "takeProfitPrice": 0,
                    "userId": "809519987784454146"
                }
            ]
        },
        "message": ""
    }
```

## 18. 查询止盈止损历史订单列表

**HTTP请求**

```http
    # Request
    POST api/v1/user/historyStopOrders
```

**请求方式**

    POST

**请求参数**

| 参数名           | 参数类型  | 必填 | 描述    |
| --------------- |---------|------|--------|
| apiKey          | String  | 是   | 接口密钥 |
| timestamp       | String  | 是   | 发起请求的时间戳，单位为毫秒 |
| symbol          | String  | 是   | 合约名称 |
| lastOrderId     | int64   | 是   | 用于分页, 第一次填写0, 后续填写前一次返回结果里面的最后一个订单id |
| length          | int64   | 是   | 每次请求的长度, 最大值为100 |

**返回值说明**

| 参数名           | 参数类型  | 描述  |
| --------------- |---------| ----- |
| userId          | String  | 用户id |
| orderId         | String  | 订单id |
| symbol          | String  | 合约名称 |
| positionId      | String  | 仓位id |
| stopLossPrice   | Float64 | 止损价格 |
| takeProfitPrice | Float64 | 止盈价格 |
| entrustVolume   | Float64 | 委托数量 |
| side            | String  | 交易方向, Bid买入, Ask卖出 |
| orderStatus     | String  | 订单状态 |
| entrustTm       | String  | 委托时间 |
| triggerTm       | String  | 触发时间 |

**备注**

| orderStatus         | 字段说明 |
| --------------------|--------|
| TriggerStopLoss     | 止损触发 |
| TriggerTakeProfit   | 止盈触发 |
| Cancelled           | 已撤销   |
| Failed              | 失败     |

```javascript
# Response
    {
        "code": 0,
        "data": {
            "orders": [
                {
                    "entrustTm": "2021-05-17T10:13:46.000Z",
                    "entrustVolume": 0.001,
                    "orderId": "47513",
                    "orderStatus": "TriggerTakeProfit",
                    "positionId": "74578",
                    "side": "Ask",
                    "stopLossPrice": 0,
                    "symbol": "BTC-USDT",
                    "takeProfitPrice": 45400,
                    "triggerTm": "2021-05-17T11:34:06.000Z",
                    "userId": "809519987784454146"
                }
            ]
        },
        "message": ""
    }
```
    
[Bingbon]: https://bingbon.pro
[English Docs]: https://bingbon.pro
[Unix Epoch]: https://en.wikipedia.org/wiki/Unix_time
