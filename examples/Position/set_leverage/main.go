package main

import (
	"context"
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
)

func main() {
	SetPositionLeverage()
}

func SetPositionLeverage() {
	client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "XRPUSDT", "buyLeverage": "20", "sellLeverage": "20"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).SetPositionLeverage(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(serverResult))
}
