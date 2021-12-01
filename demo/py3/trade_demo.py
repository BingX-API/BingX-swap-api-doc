#coding: utf-8

import urllib.request
import json
import base64
import hmac
import time

APIURL = "https://api-swap-rest.bingbon.pro"
APIKEY = "Set your api key here !!"
SECRETKEY = "Set your secret key here!!"

def genSignature(path, method, paramsMap):
    sortedKeys = sorted(paramsMap)
    paramsStr = "&".join(["%s=%s" % (x, paramsMap[x]) for x in sortedKeys])
    paramsStr = method + path + paramsStr
    return hmac.new(SECRETKEY.encode("utf-8"), paramsStr.encode("utf-8"), digestmod="sha256").digest()

def post(url, body):
    req = urllib.request.Request(url, data=body.encode("utf-8"), headers={'User-Agent': 'Mozilla/5.0'})
    return urllib.request.urlopen(req).read()

def getBalance():
    paramsMap = {
        "apiKey": APIKEY,
        "timestamp": int(time.time()*1000),
        "currency": "USDT",
    }
    sortedKeys = sorted(paramsMap)
    paramsStr = "&".join(["%s=%s" % (x, paramsMap[x]) for x in sortedKeys])
    paramsStr += "&sign=" + urllib.parse.quote(base64.b64encode(genSignature("/api/v1/user/getBalance", "POST", paramsMap)))
    url = "%s/api/v1/user/getBalance" % APIURL
    return post(url, paramsStr)

def getPositions(symbol):
    paramsMap = {
        "symbol": symbol,
        "apiKey": APIKEY,
        "timestamp": int(time.time()*1000),
    }
    sortedKeys = sorted(paramsMap)
    paramsStr = "&".join(["%s=%s" % (x, paramsMap[x]) for x in sortedKeys])
    paramsStr += "&sign=" + urllib.parse.quote(base64.b64encode(genSignature("/api/v1/user/getPositions", "POST", paramsMap)))
    url = "%s/api/v1/user/getPositions" % APIURL
    return post(url, paramsStr)

def placeOrder(symbol, side, price, volume, tradeType, action):
    paramsMap = {
        "symbol": symbol,
        "apiKey": APIKEY,
        "side": side,
        "entrustPrice": price,
        "entrustVolume": volume,
        "tradeType": tradeType,
        "action": action,
        "timestamp": int(time.time()*1000),
    }
    sortedKeys = sorted(paramsMap)
    paramsStr = "&".join(["%s=%s" % (x, paramsMap[x]) for x in sortedKeys])
    paramsStr += "&sign=" + urllib.parse.quote(base64.b64encode(genSignature("/api/v1/user/trade", "POST", paramsMap)))
    url = "%s/api/v1/user/trade" % APIURL
    return post(url, paramsStr)

def main():
    print("getBalance:", getBalance())

    print("placeOpenOrder:", placeOrder("BTC-USDT", "Bid", 0, 0.0004, "Market", "Open"))

    print("getPositions:", getPositions("BTC-USDT"))

    print("placeCloseOrder:", placeOrder("BTC-USDT", "Ask", 0, 0.0004, "Market", "Close"))

if __name__ == "__main__":
    main()
