package main

import (
	"context"
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
)

func main() {
	SpreadTradeOrderMap()
}
func SpreadTradeOrderMap() {
	client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{
		"symbol":      "SOLUSDT_SOL/USDT",
		"side":        "Buy",
		"orderType":   "Limit",
		"qty":         "0.1",
		"price":       "21",
		"orderLinkId": "1744072052193428479",
		"timeInForce": "PostOnly",
	}
	orderResult, err := client.NewUtaBybitServiceWithParams(params).PlaceSpreadTradeOrder(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(orderResult))
}
