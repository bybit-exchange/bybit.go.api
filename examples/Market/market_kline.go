package main

import (
	"context"
	"fmt"
	bybit "github.com/0xArchiviste/bybit.go.api"
)

func main() {
	GetMarketKline()
}

func GetMarketKline() {
	client := bybit.NewBybitHttpClient("", "", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"category": "spot", "symbol": "BTCUSDT", "interval": "1"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetMarketKline(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(serverResult))
}
