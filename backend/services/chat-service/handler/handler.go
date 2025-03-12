package chat_handler

import (
	"encoding/json"
	"fmt"
	"github.com/ProjectSMAA/commons/logging"
	"github.com/ProjectSMAA/commons/models"
	"github.com/SwanHtetAungPhyo/chat/middleware"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
)

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	logging.AppLogger.Info("TokenHandler", zap.String("user_id", userId))
	token := middleware.GenerateJWT(userId)
	tokenStruct := map[string]string{
		"token": token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(tokenStruct)
	if err != nil {
		return
	}
}

// HandleConnection manages WebSocket client connections.
func HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgradeConnection(w, r)
	if err != nil {
		return
	}
	defer closeConnection(conn)

	clientID, err := registerClient(conn)
	if err != nil {
		return
	}

	handleClientMessages(conn, clientID, r)
}

// upgradeConnection upgrades an HTTP connection to a WebSocket connection.
func upgradeConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := models.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logging.AppLogger.Error("websocket upgrader.Upgrade", zap.Error(err))
		return nil, err
	}
	return conn, nil
}

// closeConnection safely closes the WebSocket connection.
func closeConnection(conn *websocket.Conn) {
	if err := conn.Close(); err != nil {
		logging.AppLogger.Error("websocket conn.Close", zap.Error(err))
	}
}

// registerClient handles client registration.
func registerClient(conn *websocket.Conn) (string, error) {
	var req struct {
		Type   string `json:"type"`
		Sender string `json:"sender"`
	}
	if err := conn.ReadJSON(&req); err != nil {
		logging.AppLogger.Error("ReadJSON", zap.Error(err))
		return "", err
	}

	if req.Type != "new_client" || req.Sender == "" {
		logging.AppLogger.Info("Invalid Registration")
		_ = conn.WriteMessage(websocket.TextMessage, []byte("Invalid Registration"))
		return "", fmt.Errorf("invalid Registration")
	}

	models.Mu.Lock()
	models.Clients[req.Sender] = conn
	models.Mu.Unlock()

	token := middleware.GenerateJWT(req.Sender)
	if err := conn.WriteJSON(map[string]string{"token": token}); err != nil {
		return "", err
	}

	return req.Sender, nil
}

// handleClientMessages processes incoming messages from a WebSocket client.
func handleClientMessages(conn *websocket.Conn, clientID string, r *http.Request) {
	for {
		var msg models.Message
		if err := conn.ReadJSON(&msg); err != nil {
			logging.AppLogger.Error("ReadJSON", zap.Error(err))
			return
		}

		if r.Header == nil {
			_ = conn.WriteMessage(websocket.TextMessage, []byte("Auth Header Required"))
			continue
		}

		authHeader := r.Header.Get("Authorization")
		if !middleware.WebSockMiddleWare(authHeader) {
			_ = conn.WriteMessage(websocket.TextMessage, []byte("Invalid Token"))
			return
		}

		sendMessageToClient(msg)
	}
}

// sendMessageToClient forwards a message to the intended recipient.
func sendMessageToClient(message models.Message) {
	models.Mu.Lock()
	defer models.Mu.Unlock()

	conn, exists := models.Clients[message.Receiver]
	if !exists {
		logging.AppLogger.Error("Client not found")
		return
	}

	if err := conn.WriteJSON(message); err != nil {
		logging.AppLogger.Error("WriteJSON", zap.Error(err))
		_ = conn.Close()
		delete(models.Clients, message.Sender)
		return
	}

	logging.AppLogger.Info("Message sent to client", zap.String("sender", message.Sender))
}
