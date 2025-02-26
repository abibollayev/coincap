package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	green = "\033[32m"
	red   = "\033[31m"
	reset = "\033[0m"
)

type SubscribeMessage struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
}

type CryptoData struct {
	D struct {
		ID    int     `json:"id"`
		Price float64 `json:"p"`
	} `json:"d"`
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Error: You must provide exactly one argument (btc or eth).")
	}

	coin := strings.ToLower(os.Args[1])
	if coin != "btc" && coin != "eth" {
		log.Fatal("Error: Invalid argument. Use 'btc' or 'eth' only.")
	}

	u := url.URL{
		Scheme:   "wss",
		Host:     "push.coinmarketcap.com",
		Path:     "/ws",
		RawQuery: "device=web&client_source=coin_detail_page",
	}

	headers := http.Header{}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), headers)
	if err != nil {
		log.Fatalf("Error: Connection failed: %v", err)
	}
	defer conn.Close()

	cryptoCoin := map[string]string{
		"btc": "1",
		"eth": "1027",
	}

	subMsg := SubscribeMessage{
		Method: "RSUBSCRIPTION",
		Params: []string{
			"main-site@crypto_price_5s@{}@normal",
			cryptoCoin[coin],
		},
	}

	msg, _ := json.Marshal(subMsg)
	err = conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Fatalf("Error: Sending subscribe message: %v", err)
	}

	var lastPrice float64

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error: Read message", err)
			break
		}

		var crypto CryptoData
		err = json.Unmarshal(message, &crypto)
		if err != nil {
			log.Println("Error: Decode JSON:", err)
			continue
		}

		if crypto.D.Price == 0 {
			continue
		}

		price := crypto.D.Price
		var arrow string
		var color string

		if lastPrice == 0 {
			arrow = "-"
			color = reset
		} else if price > lastPrice {
			arrow = "↑"
			color = green
		} else if price < lastPrice {
			arrow = "↓"
			color = red
		}

		lastPrice = price

		fmt.Printf("\r%s: %.2f$ %s%s%s ", strings.ToUpper(coin), price, color, arrow, reset)
		time.Sleep(500 * time.Microsecond)
	}
}
