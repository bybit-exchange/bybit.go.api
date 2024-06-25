package main

import (
	"context"
	"fmt"

	bybit "github.com/bybit-exchange/bybit.go.api"
)

func main() {
	PlaceTrade()
}

func PlaceTrade() {
	client := bybit.NewBybitHttpClient("8wYkmpLsMg10eNQyPm", "Ouxc34myDnXvei54XsBZgoQzfGxO4bkr2Zsj", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"accountType": "UNIFIED", "category": "linear"}
	accountResult, err := client.NewAccountService(params).GetTransactionLog(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(accountResult))
}
