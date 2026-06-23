package main

import (
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
)

func main() {
	ws := bybit.NewBybitPublicWebSocket(bybit.SPREAD_MAINNET, func(message string) error {
		fmt.Println("Received:", message)
		return nil
	})
	_, _ = ws.Connect().SendSubscription([]string{"orderbook.25.SOLUSDT_SOL/USDT"})
	select {}
}
