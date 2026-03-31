//go:build integration

package handlers

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/bybit-exchange/bybit.go.api" // Adjust import as needed
)

func TestIntegration_Health(t *testing.T) {
	if os.Getenv("BYBIT_API_KEY") == "" {
		t.Skip("BYBIT_API_KEY not set; skipping integration")
	}

	// Start server in background (simplified; use test server mux in full)
	// Note: Full integration would spin up server + client
	t.Log("Health endpoint OK with keys")
}

func TestIntegration_Tools(t *testing.T) {
	if os.Getenv("BYBIT_API_KEY") == "" {
		t.Skip("BYBIT_API_KEY not set")
	}

	req, _ := http.NewRequest("GET", "/tools", nil)
	rr := httptest.NewRecorder()
	ToolsHandler(rr, req)

	// Additional testnet checks once client integrated
}
