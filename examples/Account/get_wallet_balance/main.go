package main

import (
	"context"
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
)

func main() {
	GetAccountWallet()
}

func GetAccountWallet() {
	client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"accountType": "UNIFIED"}
	accountResult, err := client.NewUtaBybitServiceWithParams(params).GetFeeRates(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(accountResult))
}
