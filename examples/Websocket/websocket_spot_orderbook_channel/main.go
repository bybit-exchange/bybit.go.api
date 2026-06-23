package main

import (
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
)

func main() {
	ws := bybit.NewBybitPublicWebSocket("wss://stream.bybit.com/v5/public/spot", func(message string) error {
		fmt.Println("Received:", message)
		return nil
	})
	_, _ = ws.Connect().SendSubscription([]string{"orderbook.1.BTCUSDT", "orderbook.1.ETHUSDT"})
	select {}
}
