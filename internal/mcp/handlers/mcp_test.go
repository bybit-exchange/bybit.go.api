package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestToolsHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/tools", nil)
	rr := httptest.NewRecorder()
	ToolsHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := []ToolSchema{}
	var got []ToolSchema
	json.NewDecoder(rr.Body).Decode(&got)
	if len(got) != 0 {
		t.Errorf("expected empty tools, got %v", got)
	}
}

func TestCallHandler_InvalidJSON(t *testing.T) {
	req, _ := http.NewRequest("POST", "/call", bytes.NewBufferString("invalid"))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	CallHandler(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", status)
	}
}

func TestCallHandler(t *testing.T) {
	payload := `{"name": "test", "args": {}}`
	req, _ := http.NewRequest("POST", "/call", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	CallHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected 200, got %d", status)
	}

	var res ToolResult
	json.NewDecoder(rr.Body).Decode(&res)
	if res.Content != `{"error": "No tools implemented yet. See PR #2 for market tools."}` {
		t.Errorf("unexpected content: %s", res.Content)
	}
}
