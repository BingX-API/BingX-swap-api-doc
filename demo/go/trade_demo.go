package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	urlStr    = "https://api-swap-rest.bingbon.pro"
	apiKey    = "Set your api key here !!"
	secretKey = "Set your secret key here !!"
)

func computeHmac256(strMessage string, strSecret string) string {
	key := []byte(strSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(strMessage))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func createSign(method, path string, mapParams url.Values) string {
	strParams := method + path + mapParams.Encode()
	return url.QueryEscape(computeHmac256(strParams, secretKey))
}

func post(requestUrl string) (responseData []byte, err error) {
	response, err := http.Post(requestUrl, "", nil)
	if err != nil {
		return
	}
	defer response.Body.Close()

	responseData, err = ioutil.ReadAll(response.Body)
	return
}

func getBalance() {
	requestPath := "/api/v1/user/getBalance"
	requestUrl := urlStr + requestPath
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	mapParams := url.Values{}
	mapParams.Set("currency", "USDT")
	mapParams.Set("apiKey", apiKey)
	mapParams.Set("timestamp", fmt.Sprint(timestamp))
	mapParams.Set("sign", createSign(http.MethodPost, requestPath, mapParams))

	requestUrl += "?"
	requestUrl += mapParams.Encode()
	responseData, err := post(requestUrl)
	fmt.Println("\t", string(responseData), err)
}

func getPositions(symbol string) {
	requestPath := "/api/v1/user/getPositions"
	requestUrl := urlStr + requestPath
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	mapParams := url.Values{}
	mapParams.Set("symbol", symbol)
	mapParams.Set("apiKey", apiKey)
	mapParams.Set("timestamp", fmt.Sprint(timestamp))
	mapParams.Set("sign", createSign(http.MethodPost, requestPath, mapParams))

	requestUrl += "?"
	requestUrl += mapParams.Encode()
	responseData, err := post(requestUrl)
	fmt.Println("\t", string(responseData), err)
}

func placeOrder(symbol, side, price, volume, tradeType, action string) {
	requestPath := "/api/v1/user/trade"
	requestUrl := urlStr + requestPath
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	mapParams := url.Values{}
	mapParams.Set("symbol", symbol)
	mapParams.Set("apiKey", apiKey)
	mapParams.Set("side", side)
	mapParams.Set("entrustPrice", price)
	mapParams.Set("entrustVolume", volume)
	mapParams.Set("tradeType", tradeType)
	mapParams.Set("action", action)
	mapParams.Set("timestamp", fmt.Sprint(timestamp))
	mapParams.Set("sign", createSign(http.MethodPost, requestPath, mapParams))

	requestUrl += "?"
	requestUrl += mapParams.Encode()
	responseData, err := post(requestUrl)
	fmt.Println("\t", string(responseData), err)
}

func main() {
	fmt.Println("getBalance:")
	getBalance()

	fmt.Println("placeOpenOrder:")
	placeOrder("BTC-USDT", "Bid", "0", "0.0004", "Market", "Open")

	fmt.Println("getPositions:")
	getPositions("BTC-USDT")

	fmt.Println("placeCloseOrder:")
	placeOrder("BTC-USDT", "Ask", "0", "0.0004", "Market", "Close")
}
