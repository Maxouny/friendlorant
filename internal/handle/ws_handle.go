package handle

import (
	"context"
	"net/http"
	"sync"

	"friendlorant/internal/models"

	"github.com/gorilla/websocket"
	"github.com/lxzan/gws"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type client struct {
	conn *gws.Conn
	send chan []byte
	user *models.User
}

type wsHandler struct {
	clients map[*client]*client
	mutex   sync.Mutex
}

func (w *wsHandler) HandleConnect(ctx context.Context, conn *gws.Conn) {
	client := &client{
		conn: conn,
		send: make(chan []byte, 256),
		user: &models.User{},
	}
	w.mutex.Lock()
	w.clients[client] = client
	w.mutex.Unlock()
}

func (w *wsHandler) HandleDisconnect(ctx context.Context, conn *gws.Conn) {
	w.mutex.Lock()
	delete(w.clients, &client{conn: conn})
	w.mutex.Unlock()
}
