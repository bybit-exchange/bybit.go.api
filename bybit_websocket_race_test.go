package bybit_connector

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// TestWebSocketIsConnectedRace reproduces the data race on WebSocket.isConnected.
// Connect() starts handleIncomingMessages() and monitorConnection() as separate
// goroutines; the former writes b.isConnected = false when the read loop ends and
// the latter reads and writes b.isConnected every tick — all on a plain bool with
// no synchronization. Run with `go test -race`.
func TestWebSocketIsConnectedRace(t *testing.T) {
	upgrader := websocket.Upgrader{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		// Close almost immediately so handleIncomingMessages() hits a read error
		// and writes b.isConnected = false, concurrently with monitorConnection().
		_ = c.Close()
	}))
	defer srv.Close()

	wsURL := "ws" + strings.TrimPrefix(srv.URL, "http")

	ws := NewBybitPublicWebSocket(wsURL, func(message string) error { return nil })
	if ws.Connect() == nil {
		t.Fatal("connect failed")
	}
	// Let the goroutines run and race on isConnected.
	time.Sleep(50 * time.Millisecond)
	_ = ws.Disconnect()
}
