package main

import (
	"context"
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
)

func main() {
	GetMarketTicker()
}

func GetMarketTicker() {
	client := bybit.NewBybitHttpClient("", "", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"category": "spot", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetMarketTickers(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(serverResult))
}
