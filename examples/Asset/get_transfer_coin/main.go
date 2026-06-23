package main

import (
	"context"
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
)

func main() {
	PlaceOrder()
}

func PlaceOrder() {
	client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"fromAccountType": "UNIFIED", "toAccountType": "CONTRACT"}
	response, err := client.NewUtaBybitServiceWithParams(params).GetTransferableCoin(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(response))
}
