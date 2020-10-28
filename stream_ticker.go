package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

// Msg struct to send message over web socket

func main() {
	u := url.URL{Scheme: "wss", Host: "fstream3.binance.com", Path: "/ws/MARKET_DATA"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	msg := `{"method": "SUBSCRIBE","params":["!bookTicker"],"id": 1}`
	err = c.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Println("write:", err)
		return
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}

}
