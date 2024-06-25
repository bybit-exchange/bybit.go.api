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
	orderResult, err := client.NewPlaceOrderService("linear", "XRPUSDT", "Buy", "Market", "10").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(orderResult))
}
