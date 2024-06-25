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
	client := bybit.NewBybitHttpClient("8wYkmpLsMg10eNQyPm", "Ouxc34myDnXvei54XsBZgoQzfGxO4bkr2Zsj", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"fromAccountType": "UNIFIED", "toAccountType": "CONTRACT"}
	response, err := client.NewAssetService(params).GetTransferableCoin(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(response))
}
