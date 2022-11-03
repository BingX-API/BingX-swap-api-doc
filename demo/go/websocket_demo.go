package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "open-ws-swap.bingbon.pro", "service address")

// reqType request type
const (
	// ReqTypeSubscribe sub
	ReqTypeSubscribe string = "sub"

	// ReqTypeUnsubscribe unsub
	ReqTypeUnsubscribe string = "unsub"
)

// WsRequest websocket request
type WsRequest struct {
	ReqID    string `json:"id"`
	ReqType  string `json:"reqType"`
	DataType string `json:"dataType"`
}

// wsMessage Packing ws messages
type wsMessage struct {
	messageType int
	messageData []byte
}

// convertMessage Print the decompressed message
func convertMessage(messageData []byte) []byte {
	buffer := bytes.NewBuffer(messageData)
	reader, _ := gzip.NewReader(buffer)
	data, _ := ioutil.ReadAll(reader)
	return data
}

// getCli Get client request
func getCli() *websocket.Conn {
	url := url.URL{Scheme: "wss", Host: *addr, Path: "/ws"}
	header := http.Header{}
	cli, _, err := websocket.DefaultDialer.Dial(url.String(), header)
	if err != nil {
		panic(fmt.Sprint("getCli dial fail, err:", err))
		return nil
	}
	return cli
}

// receiveMessages Receive messages from the server and print
func receiveMessages(cli *websocket.Conn, writeCh chan wsMessage) {
	for {
		_, messageData, err := cli.ReadMessage()
		if err != nil {
			fmt.Println("testReq read message fail, err:", err)
			close(writeCh)
			break
		}
		message := string(convertMessage(messageData))
		if message == "Ping" {
			writeCh <- wsMessage{
				messageType: websocket.TextMessage,
				messageData: []byte("Pong"),
			}
		}
		fmt.Println(time.Now(), ": receiveMessages read text message success, message:", message)
	}
}

// testSub Test subscription logic
func testSub(dataTypes []string) {
	cli := getCli()
	writeCh := make(chan wsMessage)
	go receiveMessages(cli, writeCh)

	sendReq := func(reqType, dataType string) {
		request := &WsRequest{
			ReqID:    "test",
			ReqType:  reqType,
			DataType: dataType,
		}
		messageData, _ := json.Marshal(request)
		err := cli.WriteMessage(websocket.TextMessage, messageData)
		if err != nil {
			fmt.Println("sendReq write message fail, err:", err)
			return
		}
		fmt.Println("sendReq write message success, message:", string(messageData))
	}

	sub := func(dataType string) {
		sendReq(ReqTypeSubscribe, dataType)
	}
	unsub := func(dataType string) {
		sendReq(ReqTypeUnsubscribe, dataType)
	}

	time.Sleep(2 * time.Second)
	for _, dataType := range dataTypes {
		sub(dataType)
	}

	// Unsubscribe after one minute
	unSubTimout := time.NewTimer(600 * time.Second)
OUT:
	for {
		select {
		case message, ok := <-writeCh:
			if !ok {
				break OUT
			}
			err := cli.WriteMessage(message.messageType, message.messageData)
			if err != nil {
				fmt.Println("testSub write message fail, err:", err)
				break OUT
			}
			fmt.Println("testSub write pong message success")
		case <-unSubTimout.C:
			for _, dataType := range dataTypes {
				unsub(dataType)
			}
		}
	}
}

func main() {
	flag.Parse()

	// Test the case of subscribing to messages
	dataTypes := []string{
		// todo Fill in the type you want to subscribe
		"market.kline.BTC-USDT.1min",
		"market.kline.BTC-USDT.1hour",
	}
	go testSub(dataTypes)

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
