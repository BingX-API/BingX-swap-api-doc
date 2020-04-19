package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"time"
)

var flagAPIHTTPTradeURL = flag.String("api_http_trade_url",
	"https://api-swap-rest.bingbon.pro/api/v1/user", "")
var flagAPIHTTPMarketURL = flag.String("api_http_market_url",
	"https://api-swap-rest.bingbon.pro/api/v1/market", "")
var flagAPIKey = flag.String("api_key", "testAPIKey", "")
var flagScrectKey = flag.String("screct_key", "testScrectKey", "")

// HTTPParam 参数
type HTTPParam struct {
	Key   string
	Value string
}

// mapToURLQuery 构建参数
func mapToURLQuery(mapParams map[string]string) string {
	var strParams string
	for key, value := range mapParams {
		strParams += (key + "=" + value + "&")
	}

	if 0 < len(strParams) {
		strParams = string([]rune(strParams)[:len(strParams)-1])
	}

	return strParams
}

// createSign 签名加密
func createSign(mapParams map[string]string) string {
	sortedParams := mapSortByKey(mapParams)
	strParams := listURLQuery(sortedParams)

	fmt.Println("createSign strParams:", strParams)
	return url.QueryEscape(computeHmac256(strParams, *flagScrectKey))
}

// mapSortByKey 参数排序
func mapSortByKey(mapValue map[string]string) []HTTPParam {
	var keys []string
	for key := range mapValue {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	mapReturn := make([]HTTPParam, 0)
	for _, key := range keys {
		httpParam := HTTPParam{Key: key, Value: mapValue[key]}
		mapReturn = append(mapReturn, httpParam)
	}

	return mapReturn
}

// mapValueEncodeURI URI Encode
func mapValueEncodeURI(mapValue map[string]string) map[string]string {
	for key, value := range mapValue {
		valueEncodeURI := url.QueryEscape(value)
		mapValue[key] = valueEncodeURI
	}

	return mapValue
}

// listURLQuery 参数拼接
func listURLQuery(listParams []HTTPParam) string {
	var strParams string
	for _, httpParam := range listParams {
		strParams += (httpParam.Key + "=" + httpParam.Value + "&")
	}

	if 0 < len(strParams) {
		strParams = string([]rune(strParams)[:len(strParams)-1])
	}

	return strParams
}

// computeHmac256 计算哈希值
func computeHmac256(strMessage string, strSecret string) string {
	key := []byte(strSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(strMessage))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// HTTPGetRequest HTTP get 请求
func HTTPGetRequest(strURL string, mapParams map[string]string) string {
	httpClient := &http.Client{}

	var strRequestURL string
	if nil == mapParams {
		strRequestURL = strURL
	} else {
		strParams := mapToURLQuery(mapParams)
		strRequestURL = strURL + "?" + strParams
	}

	fmt.Println("url=", strRequestURL)
	request, err := http.NewRequest("GET", strRequestURL, nil)
	if nil != err {
		return err.Error()
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")

	response, err := httpClient.Do(request)
	if nil != err {
		return err.Error()
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err.Error()
	}

	return string(body)
}

// HTTPPostRequest HTTP Post
func HTTPPostRequest(strURL string, mapParams map[string]string) string {
	httpClient := &http.Client{}

	var strRequestURL string
	if nil == mapParams {
		strRequestURL = strURL
	} else {
		strParams := mapToURLQuery(mapParams)
		strRequestURL = strURL + "?" + strParams
	}

	fmt.Println("url=", strRequestURL)
	request, err := http.NewRequest("POST", strRequestURL, nil)
	if nil != err {
		return err.Error()
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")

	response, err := httpClient.Do(request)
	if nil != err {
		return err.Error()
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err.Error()
	}

	return string(body)
}

func queryOrderInfo(symbol, orderID string) {
	strURL := *flagAPIHTTPTradeURL + "/queryOrderStatus"
	timestamp := time.Now().Unix()
	mapParams := make(map[string]string, 0)
	mapParams["symbol"] = symbol
	mapParams["orderId"] = orderID
	mapParams["apiKey"] = *flagAPIKey
	mapParams["timestamp"] = fmt.Sprintf("%d", timestamp)
	mapParams["sign"] = createSign(mapParams)
	fmt.Println("queryOrderInfo sign:", mapParams["sign"])
	resp := HTTPPostRequest(strURL, mapParams)
	fmt.Println("queryOrderInfo resp:", resp)
	return
}

func queryBalance(currency string) {
	strURL := *flagAPIHTTPTradeURL + "/getBalance"
	timestamp := time.Now().Unix()
	mapParams := make(map[string]string, 0)
	mapParams["currency"] = currency
	mapParams["apiKey"] = *flagAPIKey
	mapParams["timestamp"] = fmt.Sprintf("%d", timestamp)
	mapParams["sign"] = createSign(mapParams)
	fmt.Println("queryBalance sign:", mapParams["sign"])
	resp := HTTPPostRequest(strURL, mapParams)
	fmt.Println("queryBalance resp:", resp)
	return
}

func doTrade(symbol, side, tradeType, action, price, volume string) {
	strURL := *flagAPIHTTPTradeURL + "/trade"
	timestamp := time.Now().Unix()
	mapParams := make(map[string]string, 0)
	mapParams["symbol"] = symbol
	mapParams["side"] = side
	mapParams["tradeType"] = tradeType
	mapParams["action"] = action
	mapParams["entrustPrice"] = price
	mapParams["entrustVolume"] = volume
	mapParams["apiKey"] = *flagAPIKey
	mapParams["timestamp"] = fmt.Sprintf("%d", timestamp)
	mapParams["sign"] = createSign(mapParams)
	fmt.Println("doTrade sign:", mapParams["sign"])
	resp := HTTPPostRequest(strURL, mapParams)
	fmt.Println("doTrade resp:", resp)
	return
}

func cancelOrder(symbol, orderID string) {
	strURL := *flagAPIHTTPTradeURL + "/cancelOrder"
	timestamp := time.Now().Unix()
	mapParams := make(map[string]string, 0)
	mapParams["symbol"] = symbol
	mapParams["orderId"] = orderID
	mapParams["apiKey"] = *flagAPIKey
	mapParams["timestamp"] = fmt.Sprintf("%d", timestamp)
	mapParams["sign"] = createSign(mapParams)
	fmt.Println("cancelOrder sign:", mapParams["sign"])
	resp := HTTPPostRequest(strURL, mapParams)
	fmt.Println("cancelOrder resp:", resp)
	return
}

func getPendingOrders(symbol string) {
	strURL := *flagAPIHTTPTradeURL + "/pendingOrders"
	timestamp := time.Now().Unix()
	mapParams := make(map[string]string, 0)
	mapParams["symbol"] = symbol
	mapParams["apiKey"] = *flagAPIKey
	mapParams["timestamp"] = fmt.Sprintf("%d", timestamp)
	mapParams["sign"] = createSign(mapParams)
	fmt.Println("getPendingOrders sign:", mapParams["sign"])
	resp := HTTPPostRequest(strURL, mapParams)
	fmt.Println("getPendingOrders resp:", resp)
	return
}

func getPositions(symbol string) {
	strURL := *flagAPIHTTPTradeURL + "/getPositions"
	timestamp := time.Now().Unix()
	mapParams := make(map[string]string, 0)
	mapParams["symbol"] = symbol
	mapParams["apiKey"] = *flagAPIKey
	mapParams["timestamp"] = fmt.Sprintf("%d", timestamp)
	mapParams["sign"] = createSign(mapParams)
	fmt.Println("getPositions sign:", mapParams["sign"])
	resp := HTTPPostRequest(strURL, mapParams)
	fmt.Println("getPositions resp:", resp)
	return
}

func getAllContracts() {
	strURL := *flagAPIHTTPMarketURL + "/getAllContracts"
	mapParams := make(map[string]string, 0)
	resp := HTTPGetRequest(strURL, mapParams)
	fmt.Println("getAllContracts resp:", resp)
	return
}

func getOrderBook(symbol string) {
	strURL := *flagAPIHTTPMarketURL + "/getMarketDepth"
	mapParams := make(map[string]string, 0)
	mapParams["symbol"] = symbol
	resp := HTTPGetRequest(strURL, mapParams)
	fmt.Println("getOrderBook resp:", resp)
	return
}

func getMarketTrades(symbol string) {
	strURL := *flagAPIHTTPMarketURL + "/getMarketTrades"
	mapParams := make(map[string]string, 0)
	mapParams["symbol"] = symbol
	resp := HTTPGetRequest(strURL, mapParams)
	fmt.Println("getMarketTrades resp:", resp)
	return
}

func getHistoryFunding(symbol string) {
	strURL := *flagAPIHTTPMarketURL + "/getHistoryFunding"
	mapParams := make(map[string]string, 0)
	mapParams["symbol"] = symbol
	resp := HTTPGetRequest(strURL, mapParams)
	fmt.Println("getHistoryFunding resp:", resp)
	return
}

func getLatestKline(symbol, klineType string) {
	strURL := *flagAPIHTTPMarketURL + "/getLatestKline"
	mapParams := make(map[string]string, 0)
	mapParams["symbol"] = symbol
	mapParams["klineType"] = klineType
	resp := HTTPGetRequest(strURL, mapParams)
	fmt.Println("getLatestKline resp:", resp)
	return
}

func getHistoryKlines(symbol, klineType, startTs, endTs string) {
	strURL := *flagAPIHTTPMarketURL + "/getHistoryKlines"
	mapParams := make(map[string]string, 0)
	mapParams["symbol"] = symbol
	mapParams["klineType"] = klineType
	mapParams["startTs"] = startTs
	mapParams["endTs"] = endTs
	resp := HTTPGetRequest(strURL, mapParams)
	fmt.Println("getHistoryKlines resp:", resp)
	return
}

func getServerTime() {
	strURL := *flagAPIHTTPTradeURL + "/server/time"
	mapParams := make(map[string]string, 0)
	resp := HTTPPostRequest(strURL, mapParams)
	fmt.Println("getServerTime resp:", resp)
	return
}

var usage = func() {
	fmt.Println("USAGE: command [arguments] ...")
	fmt.Println("\nThe commands are:\n\torderBook\tsymbol\n\tqueryOrder\tsymbol\torderId\n\tqueryBalance\tcurrency\n\ttrade\tsymbol\tside\ttradeType\taction\tprice\tvolume\n\tpendingOrders\tsymbol\n\tcancelOrder\tsymbol\torderId")
	fmt.Println("\tpositions\tsymbol\n\tgetAllContracts\n\tmarketTrades\tsymbol\n\thistoryFunding\tsymbol\n\tlatestKline\tsymbol\tklineType")
	fmt.Println("\thistoryKlines\tsymbol\tklineType\tstartTs\tendTs\n\tserverTime")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		switch args[0] {
		case "orderBook":
			if len(args) != 2 {
				usage()
				return
			}
			getOrderBook(args[1])
		case "marketTrades":
			if len(args) != 2 {
				usage()
				return
			}
			getMarketTrades(args[1])
		case "getAllContracts":
			getAllContracts()
		case "historyFunding":
			if len(args) != 2 {
				usage()
				return
			}
			getHistoryFunding(args[1])
		case "latestKline":
			if len(args) != 3 {
				usage()
				return
			}
			getLatestKline(args[1], args[2])
		case "historyKlines":
			if len(args) != 5 {
				usage()
				return
			}
			getHistoryKlines(args[1], args[2], args[3], args[4])
		case "queryOrder":
			if len(args) != 3 {
				usage()
				return
			}
			queryOrderInfo(args[1], args[2])
		case "queryBalance":
			if len(args) != 2 {
				usage()
				return
			}
			queryBalance(args[1])
		case "pendingOrders":
			if len(args) != 2 {
				usage()
				return
			}
			getPendingOrders(args[1])
		case "positions":
			if len(args) != 2 {
				usage()
				return
			}
			getPositions(args[1])
		case "cancelOrder":
			if len(args) != 3 {
				usage()
				return
			}
			cancelOrder(args[1], args[2])
		case "serverTime":
			getServerTime()
		case "trade":
			if len(args) != 7 {
				usage()
				return
			}
			doTrade(args[1], args[2], args[3], args[4], args[5], args[6])
		default:
			usage()
		}
		return
	}

	usage()
}
