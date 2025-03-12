package models

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

type WebSocketCtx struct {
	w http.ResponseWriter
	r *http.Request
}

const ClientRegistration = "client_registration"

var (
	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	Clients   = make(map[string]*websocket.Conn)
	Messages  = make([]Message, 0) // In-memory message history
	Mu        sync.Mutex
	JwtSecret = []byte("wafdslfadslkjfdaskljfads")
)

type Message struct {
	ChatRoomID string    `json:"chat_room_id"`
	ChatType   string    `json:"chat_type"`
	Sender     string    `json:"sender"`
	Receiver   string    `json:"receiver"`
	Content    string    `json:"content"`
	Timestamp  time.Time `json:"timestamp"`
}
