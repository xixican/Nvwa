package nvwa_socket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrade = websocket.Upgrader{
	HandshakeTimeout:  0,
	ReadBufferSize:    1024, // nvwa
	WriteBufferSize:   1024, // nvwa
	WriteBufferPool:   nil,
	Subprotocols:      nil,
	Error:             nil,
	CheckOrigin:       nil,
	EnableCompression: false,
}

type WebSocketServer struct {
}

func (w *WebSocketServer) StartServer(address string) {
	log.Fatal(http.ListenAndServe(address, &WebSocketHandler{}))
}

type WebSocketHandler struct {
}

func (h *WebSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	websocketConnection, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("websocket upgrade err:%v", err)
	}
	connectionChan <- &WebSocketConnection{conn: websocketConnection}
}

type WebSocketConnection struct {
	conn *websocket.Conn
}

func (w *WebSocketConnection) GetNextMessage() (b []byte, err error) {
	_, msgBytes, err := w.conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	return msgBytes, err
}

func (w *WebSocketConnection) Close() {
	w.conn.Close()
}
