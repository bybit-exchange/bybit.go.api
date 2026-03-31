package handlers

import (
	"encoding/json"
	"net/http"
)

// MCP Tool Schema (standard MCP format)
type ToolSchema struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Parameters  interface{} `json:"parameters"` // JSONSchema
}

// ToolCallRequest for /call
type ToolCallRequest struct {
	Name string                 `json:"name"`
	Args map[string]interface{} `json:"args"`
}

// ToolResult for /call response
type ToolResult struct {
	Content string `json:"content"`
}

var tools []ToolSchema // Will be populated in later PRs

// ToolsHandler GET /tools
func ToolsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tools)
}

// CallHandler POST /call
func CallHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ToolCallRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Stub: No tools yet
	result := ToolResult{Content: `{"error": "No tools implemented yet. See PR #2 for market tools."}`}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
