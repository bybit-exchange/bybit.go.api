package main

import (
	"context"
	"fmt"

	bybit "github.com/bybit-exchange/bybit.go.api"
)

func main() {
	GetCoinInfo()
}

func GetCoinInfo() {
	client := bybit.NewBybitHttpClient("8wYkmpLsMg10eNQyPm", "Ouxc34myDnXvei54XsBZgoQzfGxO4bkr2Zsj", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"coin": "USDT"}
	assetResult, err := client.NewAssetService(params).GetCoinInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(assetResult))
}
