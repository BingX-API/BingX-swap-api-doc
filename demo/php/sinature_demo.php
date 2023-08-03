<?php

$secretKey = "mheO6dR8ovSsxZQCOYEFCtelpuxcWGTfHw7te326y6jOwq5WpvFQ9JNljoTwBXZGv5It07m9RXSPpDQEK2w";

function getSignature(array $params) {
    global $secretKey;
    $query = buildQuery($params);
    $signature = hash_hmac('sha256', $query, $secretKey);
    return $signature;
}

function buildQuery(array $params)
    {
        $query_array = array();
        foreach ($params as $key => $value) {
            if (is_array($value)) {
                $query_array = array_merge($query_array, array_map(function ($v) use ($key) {
                    if (is_bool($v) === true) {
                        return $v ? urlencode($key) . '=true' : urlencode($key) . '=false';
                    }
                    return urlencode($key) . '=' . urlencode($v);
                }, $value));
            } else {
                if (is_bool($value) === true) {
                    $query_array[] =  $value ? urlencode($key) . '=true' : urlencode($key) . '=false';
                } else {
                    $query_array[] = urlencode($key) . '=' . urlencode($value);
                }
            }
        }
        return implode('&', $query_array);
    }

function testSign() {
    // interface params
    $params = array();
    $params['symbol'] = 'BTC-USDT';
    $params['timestamp'] = 1667872120843;
    $params['side'] = 'LONG';
    $params['leverage'] = 6;

    // generate signature
    $signature = getSignature($params);
    echo "\t";
    echo $signature;
    echo "\n";
}

echo "testSign:\n";
testSign();

?>