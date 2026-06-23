package main

import (
	"context"
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
)

func main() {
	PlaceOrderMap()
}
func PlaceOrderMap() {
	client := bybit.NewBybitHttpClient("xxxx", "xxx", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{
		"category":    "linear",
		"symbol":      "BTCUSDT",
		"side":        "Buy",
		"positionIdx": 0,
		"orderType":   "Limit",
		"qty":         "0.001",
		"price":       "10000",
		"timeInForce": "GTC",
	}
	orderResult, err := client.NewUtaBybitServiceWithParams(params).PlaceOrder(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(orderResult))
}
