<?php

$url = "https://api-swap-rest.bingbon.pro";
$apiKey = "Set your api key here!!";
$secretKey = "Set your secret key here!!";

function getOriginString(string $method, string $path, array $params) {
    // combine origin string
    $originString = $method.$path;
    $first = true;
    foreach($params as $n => $v) {
          if (!$first) {
              $originString .= "&";
          }
          $first = false;
          $originString .= $n . "=" . $v;
    }
    return $originString;
}

function getSignature(string $originString) {
    global $secretKey;
    $signature = hash_hmac('sha256', $originString, $secretKey, true);
    $signature = base64_encode($signature);
    $signature = urlencode($signature);
    return $signature;
}

function getRequestUrl(string $path, array $params) {
    global $url;
    $requestUrl = $url.$path."?";
    $first = true;
    foreach($params as $n => $v) {
          if (!$first) {
              $requestUrl .= "&";
          }
          $first = false;
          $requestUrl .= $n . "=" . $v;
    }
    return $requestUrl;
}

function httpPost($url)
{
    $curl = curl_init($url);
    curl_setopt($curl, CURLOPT_POST, true);
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($curl, CURLOPT_USERAGENT, "curl/7.80.0");
    $response = curl_exec($curl);
    curl_close($curl);
    return $response;
}

function getBalance() {
    global $apiKey;

    // interface info
    $path = "/api/v1/user/getBalance";
    $method = "POST";

    // interface params
    $params = array();
    $params['currency'] = 'USDT';
    $params['apiKey'] = $apiKey;
    $date = new DateTime();
    $params['timestamp'] = $date->getTimestamp()*1000;

    // sort params
    ksort($params);

    // generate signature
    $originString = getOriginString($method, $path, $params);
    $signature = getSignature($originString);
    $params["sign"] = $signature;

    // send http request
    $requestUrl = getRequestUrl($path, $params);
    $result = httpPost($requestUrl);
    echo "\t";
    echo $result;
    echo "\n";
}

function getPositions(string $symbol) {
    global $apiKey;

    // interface info
    $path = "/api/v1/user/getPositions";
    $method = "POST";

    // interface params
    $params = array();
    $params['symbol'] = $symbol;
    $params['apiKey'] = $apiKey;
    $date = new DateTime();
    $params['timestamp'] = $date->getTimestamp()*1000;

    // sort params
    ksort($params);

    // generate signature
    $originString = getOriginString($method, $path, $params);
    $signature = getSignature($originString);
    $params["sign"] = $signature;

    // send http request
    $requestUrl = getRequestUrl($path, $params);
    $result = httpPost($requestUrl);
    echo "\t";
    echo $result;
    echo "\n";
}

function placeOrder(string $symbol, string $side, string $price, string $volume,
    string $tradeType, string $action) {
    global $apiKey;

    // interface info
    $path = "/api/v1/user/trade";
    $method = "POST";

    // interface params
    $params = array();
    $params['symbol'] = $symbol;
    $params['apiKey'] = $apiKey;
    $params['side'] = $side;
    $params['entrustPrice'] = $price;
    $params['entrustVolume'] = $volume;
    $params['tradeType'] = $tradeType;
    $params['action'] = $action;
    $date = new DateTime();
    $params['timestamp'] = $date->getTimestamp()*1000;

    // sort params
    ksort($params);

    // generate signature
    $originString = getOriginString($method, $path, $params);
    $signature = getSignature($originString);
    $params["sign"] = $signature;

    // send http request
    $requestUrl = getRequestUrl($path, $params);
    $result = httpPost($requestUrl);
    echo "\t";
    echo $result;
    echo "\n";
}

echo "getBalance:\n";
getBalance();

echo "placeOpenOrder:\n";
placeOrder("BTC-USDT", "Bid", "0", "0.0004", "Market", "Open");

echo "getPositions:\n";
getPositions("BTC-USDT");

echo "placeCloseOrder:\n";
placeOrder("BTC-USDT", "Ask", "0", "0.0004", "Market", "Close");

?>
