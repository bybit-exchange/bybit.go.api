package main

import (
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
	"time"
)

func main() {
	ws := bybit.NewBybitPrivateWebSocket(bybit.WEBSOCKET_TRADE_TESTNET, "YOUR_API_KEY", "YOUR_API_SECRET", func(message string) error {
		fmt.Println("Received:", message)
		return nil
	}, bybit.WithPingInterval(10))
	_, _ = ws.Connect().SendRequest("order.create", map[string]interface{}{
		"symbol":      "ETHUSDT",
		"side":        "Buy",
		"orderType":   "Limit",
		"qty":         "0.2",
		"price":       "2800",
		"category":    "linear",
		"timeInForce": "PostOnly",
	},
		map[string]string{
			"X-BAPI-TIMESTAMP":   fmt.Sprintf("%d", time.Now().UnixMilli()),
			"X-BAPI-RECV-WINDOW": "8000",
			"Referer":            "bot-001",
		}, "my-custom-req-id")
	select {}
}
