package main

import (
	"context"
	"fmt"
	bybit "github.com/0xArchiviste/bybit.go.api"
)

func main() {
	GetRecentTrade()
}

func GetRecentTrade() {
	client := bybit.NewBybitHttpClient("", "", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetPublicRecentTrades(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(serverResult))
}
