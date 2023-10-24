package wss

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/octacat-app-backend/src/helper/config"
	"golang.org/x/net/websocket"
)

func InitAndRun(_ *config.AppConfig, server *http.ServeMux) http.Handler {
	listClients := map[string]*websocket.Conn{}
	userID := uuid.NewString()
	server.Handle("/ws/receive", websocket.Handler(func(ws *websocket.Conn) {
		// TODO: jwt here to get user id

		defer func() {
			ws.Close()
			delete(listClients, userID)
		}()

		// added client connection to list
		listClients[userID] = ws

		log.Println("client connected,", ws.Request().RemoteAddr)
		for {
			message := ""
			if err := websocket.Message.Receive(ws, &message); err != nil {
				log.Println("error while recheive message,", err)
				break
			}
		}
	}))

	server.Handle("/ws/send", websocket.Handler(func(c *websocket.Conn) {
		defer func() {
			c.Close()
		}()

		log.Println("client connected")

		// TODO: get user id from request
		client, isExist := listClients[userID]
		if !isExist {
			log.Println("client not online")
		}

		fmt.Printf("c.Request().Body: %v\n", func() string {
			b, _ := io.ReadAll(c.Request().Body)
			return string(b)
		}())

		// for {
		if err := websocket.Message.Send(client, "haha"); err != nil {
			log.Println("error while send message,", err)
			// break
		}
		// }
	}))

	wsc := &WebSocketConnection{
		clients: map[string]WebSocketClient{},
	}
	server.Handle("/wss/chat", websocket.Handler(wsc.Chat))
	server.Handle("/wss/chat/list", http.HandlerFunc(wsc.GetListClient))

	return (http.Handler)(server)
}
