package main

import (
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
)

func main() {
	ws := bybit.NewBybitPrivateWebSocket(bybit.SPREAD_TESTNET, "YOUR_API_KEY", "YOUR_API_SECRET", func(message string) error {
		fmt.Println("Received:", message)
		return nil
	}, bybit.WithPingInterval(2))
	_, _ = ws.Connect().SendSubscription([]string{"order", "position", "wallet"})
	select {}
}
