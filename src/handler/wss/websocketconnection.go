package wss

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/websocket"
)

type WebSocketConnection struct {
	clients map[string]WebSocketClient
}

type WebSocketClient struct {
	ID         string
	Name       string
	Connection *websocket.Conn `json:"-"`
}

func (wsc *WebSocketConnection) Chat(c *websocket.Conn) {
	query := c.Request().URL.Query()
	toUserID := query.Get("to_user_id")
	fromUserID := query.Get("from_user_id")

	log.Printf("client with user_id '%v' is connected", fromUserID)

	defer func() {
		log.Printf("client with user_id '%v' disconnected", fromUserID)
		// c.Close()
		// delete(wsc.clients, fromUserID)
	}()

	if strings.Trim(toUserID, " ") == "" {
		c.WriteClose(websocket.CloseFrame)
		log.Println("user_id is empty")
		return
	}

	if strings.Trim(fromUserID, " ") == "" {
		c.WriteClose(websocket.CloseFrame)
		log.Println("from_id is empty")
		return
	}

	if _, isOk := wsc.clients[toUserID]; !isOk {
		wsc.clients[toUserID] = WebSocketClient{
			ID:         toUserID,
			Name:       "haha",
			Connection: c,
		}
	}

	if _, isOk := wsc.clients[fromUserID]; !isOk {
		wsc.clients[fromUserID] = WebSocketClient{
			ID:         fromUserID,
			Name:       "hihi",
			Connection: c,
		}
	}

	client := wsc.clients[toUserID]
	from := wsc.clients[fromUserID]

	for {
		msg := ""
		if err := websocket.Message.Receive(c, &msg); err != nil {
			log.Println(err)
			break
		}

		// for _, client := range wsc.clients {
		if err := websocket.Message.Send(client.Connection, msg); err != nil {
			log.Println(err)
			break
		}
		if err := websocket.Message.Send(from.Connection, msg); err != nil {
			log.Println(err)
			break
		}
		// }
	}
}

func (wsc *WebSocketConnection) GetListClient(w http.ResponseWriter, r *http.Request) {
	listClient := []WebSocketClient{}
	for _, v := range wsc.clients {
		client := WebSocketClient{
			ID:   v.ID,
			Name: v.Name,
		}
		listClient = append(listClient, client)
	}

	results, err := json.Marshal(listClient)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(results)
	return
}
