package main

import (
	"context"
	"fmt"
	bybit "github.com/0xArchiviste/bybit.go.api"
)

func main() {
	GetOrderbookInfo()
}

func GetOrderbookInfo() {
	client := bybit.NewBybitHttpClient("", "", bybit.WithBaseURL(bybit.TESTNET), bybit.WithDebug(true))
	params := map[string]interface{}{"category": "spot", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetOrderBookInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(serverResult))
}
