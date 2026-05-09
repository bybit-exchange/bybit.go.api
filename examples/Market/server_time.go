package main

import (
	"context"
	"fmt"
	bybit "github.com/0xArchiviste/bybit.go.api"
)

func main() {
	GetServerTime()
}

func GetServerTime() {
	client := bybit.NewBybitHttpClient("", "", bybit.WithBaseURL(bybit.TESTNET))
	serverResult, err := client.NewUtaBybitServiceNoParams().GetServerTime(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(serverResult))
}
