package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bybit-exchange/bybit.go.api" // SDK root
	"github.com/bybit-exchange/bybit.go.api/internal/mcp/config"
	"github.com/bybit-exchange/bybit.go.api/internal/mcp/handlers"
)

func main() {
	cfg := config.Load()

	if cfg.APIKey == "" || cfg.APISecret == "" {
		log.Fatal("BYBIT_API_KEY and BYBIT_API_SECRET required")
	}

// Init Bybit client - MAINNET ALWAYS for prod server\n// Tests use BYBIT_TESTNET=true\nbaseURL := bybit.PROD // Default mainnet\nif cfg.Testnet {\n\tbaseURL = bybit.TESTNET\n}\nclient := bybit.NewBybitHttpClient(cfg.APIKey, cfg.APISecret, bybit.WithBaseURL(baseURL))\n\n// Log network\nlog.Printf("Bybit network: %s", baseURL)

	// Register handlers
	http.HandleFunc("/tools", handlers.ToolsHandler)
	http.HandleFunc("/call", handlers.CallHandler)
	http.HandleFunc("/health", healthHandler)

	port := cfg.Port
	log.Printf("MCP Server starting on :%s (testnet=%t)", port, cfg.Testnet)
	log.Printf("Bybit client initialized")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"status": "ok"}`)
}
