package wss

import (
	"net/http"

	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
	"golang.org/x/net/websocket"
)

func InitAndRun(_ *config.AppConfig, server *http.ServeMux) *http.ServeMux {
	wsc := &WebSocketConnection{
		clients: map[string]WebSocketClient{},
	}
	server.Handle("/wss/chat", websocket.Handler(wsc.Chat))
	server.Handle("/wss/chat/list", http.HandlerFunc(wsc.GetListClient))

	return server
}
