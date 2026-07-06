package main

import (
	"context"
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
)

func main() {
	GetMarketInsurance()
}

func GetMarketInsurance() {
	client := bybit.NewBybitHttpClient("", "", bybit.WithBaseURL(bybit.TESTNET), bybit.WithDebug(true))
	params := map[string]interface{}{"category": "spot", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetMarketInsurance(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(serverResult))
}
