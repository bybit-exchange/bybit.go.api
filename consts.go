package bybit_connector

const (
	Name    = "bybit.api.go"
	Version = "1.1.0"
	// Https
	MAINNET        = "https://api.bybit.com"
	MAINNET_BACKT  = "https://api.bytick.com"
	TESTNET        = "https://api-testnet.bybit.com"
	DEMO_ENV       = "https://api-demo.bybit.com"
	NETHERLAND_ENV = "https://api.bybit.nl"
	HONGKONG_ENV   = "https://api.byhkbit.com"
	TURKEY_ENV     = "https://api.bybit-tr.com"
	KAZAKHSTAN_ENV = "https://api.bybit.kz"

	// WebSocket public channel - Mainnet
	SPOT_MAINNET    = "wss://stream.bybit.com/v5/public/spot"
	LINEAR_MAINNET  = "wss://stream.bybit.com/v5/public/linear"
	INVERSE_MAINNET = "wss://stream.bybit.com/v5/public/inverse"
	SPREAD_MAINNET  = "wss://stream.bybit.com/v5/public/spread"

	// WebSocket public channel - Testnet
	SPOT_TESTNET    = "wss://stream-testnet.bybit.com/v5/public/spot"
	LINEAR_TESTNET  = "wss://stream-testnet.bybit.com/v5/public/linear"
	INVERSE_TESTNET = "wss://stream-testnet.bybit.com/v5/public/inverse"
	OPTION_TESTNET  = "wss://stream-testnet.bybit.com/v5/public/option"
	SPREAD_TESTNET  = "wss://stream-testnet.bybit.com/v5/public/spread"

	// WebSocket private channel
	WEBSOCKET_PRIVATE_MAINNET = "wss://stream.bybit.com/v5/private"
	WEBSOCKET_TRADE_MAINNET   = "wss://stream.bybit.com/v5/trade"
	WEBSOCKET_PRIVATE_TESTNET = "wss://stream-testnet.bybit.com/v5/private"
	WEBSOCKET_TRADE_TESTNET   = "wss://stream-testnet.bybit.com/v5/trade"
	WEBSOCKET_PRIVATE_DEMO    = "wss://stream-demo.bybit.com/v5/private"
	WEBSOCKET_TRADE_DEMO      = "wss://stream-demo.bybit.com/v5/trade"

	// Deprecated: V3 is deprecated and replaced by v5
	V3_CONTRACT_PRIVATE = "wss://stream.bybit.com/contract/private/v3"
	V3_UNIFIED_PRIVATE  = "wss://stream.bybit.com/unified/private/v3"
	V3_SPOT_PRIVATE     = "wss://stream.bybit.com/spot/private/v3"

	// Globals
	timestampKey  = "X-BAPI-TIMESTAMP"
	signatureKey  = "X-BAPI-SIGN"
	apiRequestKey = "X-BAPI-API-KEY"
	recvWindowKey = "X-BAPI-RECV-WINDOW"
	signTypeKey   = "X-BAPI-SIGN-TYPE"
)
