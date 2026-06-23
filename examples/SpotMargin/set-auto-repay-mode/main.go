package main

import (
	"context"
	"fmt"

	bybit "github.com/bybit-exchange/bybit.go.api"
)

func main() {
	client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{
		"currency":      "BTC",
		"autoRepayMode": "0",
	}
	result, err := client.NewUtaBybitServiceWithParams(params).SetSpotMarginAutoRepayMode(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(result))
}
