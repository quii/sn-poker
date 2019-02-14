package poker

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type playerServerWS struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

func newPlayerServerWS(w http.ResponseWriter, r *http.Request) *playerServerWS {
	conn, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("problem upgrading connection to websockets %v\n", err)
	}

	return &playerServerWS{conn: conn}
}

func (w *playerServerWS) Write(p []byte) (n int, err error) {
	err = w.conn.WriteMessage(1, p)

	if err != nil {
		return 0, err
	}

	return len(p), nil
}

func (w *playerServerWS) WaitForMsg() string {
	w.mu.Lock()
	defer w.mu.Unlock()

	_, msg, err := w.conn.ReadMessage()
	if err != nil {
		log.Printf("error reading from websocket %v\n", err)
	}
	return string(msg)
}
