# bybit-go-api
[![GO 1.21.0](https://img.shields.io/badge/Go-1.21.0-brightgreen.svg)](https://github.com/VictorFrWu/bybit-go-api)   [![Contributor Victor](https://img.shields.io/badge/contributor-Victor-blue.svg)](https://github.com/bybit-exchange/bybit-go-api)   [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/VictorFrWu/bybit-go-api/blob/main/LICENSE)
## Table of Contents
- [About](#about)
- [Release](#release)
- [Development](#development)
- [Installation](#installation)
- [Usage](#usage)
- [Contact](#contact)
- [Contributors](#contributors)
## About
The Official Go Lang API connector for Bybit's HTTP and WebSocket APIs.

Dive into a plethora of functionalities:
- Market Data Retrieval
- Trade Execution
- Position Management
- Account and Asset Info Retrieval
- User and Upgrade Management
- Public Websocket Streaming
- Private Websocket Streaming
- Institution Loan 
- Crypto Loan
- Broker
- Earn
- Spread Trade

bybit-go-api provides an official, robust, and high-performance go connector to Bybit's trading APIs.

Initially conceptualized by go developer Victor, this module is now maintained by Bybit's in-house go experts.

Your contributions are most welcome!

## Development
bybit-go-api is under active development with the latest features and updates from Bybit's API implemented promptly. The module utilizes minimal external libraries to provide a lightweight and efficient experience. If you've made enhancements or fixed bugs, please submit a pull request.

## Installation
Ensure you have go 1.21.0 or higher. And use dependencies as below
```go
require (
	github.com/google/uuid v1.4.0
	github.com/gorilla/websocket v1.5.1
	github.com/stretchr/testify v1.8.4
)
```

To import my package you need just to put the link to your go mode file
**github.com/bybit-exchange/bybit.go.api**

## Usage
Note: Replace placeholders (like YOUR_API_KEY, links, or other details) with the actual information. You can also customize this template to better fit the actual state and details of your Java API.
### Rest API
- Place an order by Map
```go
client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT", "side": "Buy", "positionIdx": 0, "orderType": "Limit", "qty": "0.001", "price": "10000", "timeInForce": "GTC"}
orderResult, err := client.NewUtaBybitServiceWithParams(params).PlaceOrder(context.Background())
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(bybit.PrettyPrint(orderResult))
```

- Place an order by Trade Class
```go
client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
orderResult, err := client.NewPlaceOrderService("linear", "XRPUSDT", "Buy", "Market", "10").Do(context.Background())
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(bybit.PrettyPrint(orderResult))
```

- Place batch order
```go
client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
params := map[string]interface{}{"category": "option",
	"request": []map[string]interface{}{
		{
			"category":    "option",
			"symbol":      "BTC-10FEB23-24000-C",
			"orderType":   "Limit",
			"side":        "Buy",
			"qty":         "0.1",
			"price":       "5",
			"orderIv":     "0.1",
			"timeInForce": "GTC",
			"orderLinkId": "9b381bb1-401",
			"mmp":         false,
			"reduceOnly":  false,
		},
		{
			"category":    "option",
			"symbol":      "BTC-10FEB23-24000-C",
			"orderType":   "Limit",
			"side":        "Buy",
			"qty":         "0.1",
			"price":       "5",
			"orderIv":     "0.1",
			"timeInForce": "GTC",
			"orderLinkId": "82ee86dd-001",
			"mmp":         false,
			"reduceOnly":  false,
		},
	},
}
orderResult, err := client.NewUtaBybitServiceWithParams(params).PlaceBatchOrder(context.Background())
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(bybit.PrettyPrint(orderResult))
```

- Get Position 
```go
client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
params := map[string]interface{}{"category": "linear", "settleCoin": "USDT", "limit": 10}
orderResult, err := client.NewUtaBybitServiceWithParams(params).GetPositionList(context.Background())
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(bybit.PrettyPrint(orderResult))
```

- Get Transaction Log
```go
client := bybit.NewBybitHttpClient("YOUR_API_KEY", "YOUR_API_SECRET", bybit.WithBaseURL(bybit.TESTNET))
params := map[string]interface{}{"accountType": "UNIFIED", "category": "linear"}
accountResult, err := client.NewUtaBybitServiceWithParams(params).GetTransactionLog(context.Background())
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(bybit.PrettyPrint(accountResult))
```

### Websocket Trade channel
```go
ws := bybit.NewBybitPrivateWebSocket(bybit.WEBSOCKET_TRADE_TESTNET, "YOUR_API_KEY", "YOUR_API_SECRET", func(message string) error {
fmt.Println("Received:", message)
return nil
}, bybit.WithPingInterval(10))
_, _ = ws.Connect().SendTradeRequest(map[string]interface{}{
"reqId": "test-005",
"header": map[string]string{
"X-BAPI-TIMESTAMP":   fmt.Sprintf("%d", time.Now().UnixMilli()),
"X-BAPI-RECV-WINDOW": "8000",
"Referer":            "bot-001",
},
"op": "order.create",
"args": []interface{}{
map[string]interface{}{
"symbol":      "ETHUSDT",
"side":        "Buy",
"orderType":   "Limit",
"qty":         "0.2",
"price":       "2800",
"category":    "linear",
"timeInForce": "PostOnly",
},
},
})
select {}
```

### Websocket public channel
- Order book Subscribe
```go
ws := bybit.NewBybitPublicWebSocket("wss://stream.bybit.com/v5/public/spot", func(message string) error {
fmt.Println("Received:", message)
return nil
})
_ = ws.Connect([]string{"orderbook.1.BTCUSDT"})
select {}
```

### Websocket private channel
```go
ws := bybit.NewBybitPrivateWebSocket("wss://stream-testnet.bybit.com/v5/private", "YOUR_API_KEY", "YOUR_API_SECRET", func(message string) error {
	fmt.Println("Received:", message)
	return nil
})
_ = ws.Connect([]string{"order"})
select {}
```

## MCP Server (New!)

Run the MCP server to expose Bybit APIs as AI tools:

```bash
# Testnet (recommended)
export BYBIT_TESTNET=true
export BYBIT_API_KEY=your_testnet_key
export BYBIT_API_SECRET=your_testnet_secret
export MCP_PORT=30001  # optional
go run cmd/mcp-server/main.go
```

Endpoints:
- `GET /tools` → Tool schemas JSON
- `POST /call` → Execute tool `{name, args}`
- `GET /health` → OK

Example:
```bash
curl http://localhost:30001/tools  # []
curl -X POST http://localhost:30001/call \
  -H 'Content-Type: application/json' \
  -d '{"name":"get_tickers","args":{"symbol":"BTCUSDT"}}'
  # {"error": "No tools yet"}
```

Tools added in future PRs.

## Contact

For support, join our Bybit API community on [Telegram](https://t.me/Bybitapi).

## Contributors
List of other contributors
<table>
  <tr>
    <td align="center">
        <a href="https://github.com/VictorFrWu">
            <img src="https://avatars.githubusercontent.com/u/32245754?v=4" width="100px;" alt=""/>
            <br />
            <sub>   
                <b>Victor</b>
            </sub>
        </a>
        <br />
        <a href="https://github.com/VictorFrWu/bybit.go.api/commits?author=bybit-exchange" title="Code">💻</a>
        <a href="https://github.com/VictorFrWu/bybit.go.api/commits?author=bybit-exchange" title="Documentation">📖</a>
    </td>
  </tr>
</table>