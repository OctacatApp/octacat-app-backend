package wss

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
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

type ChatRequest struct {
	FromUserID string `json:"from_user_id,omitempty" validate:"required"`
	ToUserID   string `json:"to_user_id,omitempty"   validate:"required"`
	Message    string `json:"message,omitempty"      validate:"required"`
}

type ChatResponse struct {
	FromUserID string `json:"from_user_id,omitempty" validate:"required"`
	ToUserID   string `json:"to_user_id,omitempty"   validate:"required"`
	Message    string `json:"message,omitempty"      validate:"required"`
}

type WSSResponse[T any] struct {
	ErrorsMSG []string `json:"errors_msg,omitempty"`
	Data      T        `json:"data,omitempty"`
}

func (wsc *WebSocketConnection) Chat(c *websocket.Conn) {
	ws := websocket.JSON
	for {
		// receive json from client
		request := new(ChatRequest)
		if err := ws.Receive(c, request); err != nil {
			result := WSSResponse[ChatResponse]{ErrorsMSG: []string{err.Error()}}
			if err := ws.Send(c, result); err != nil {
				log.Println(err)
				continue
			}

			// if connection is done, break the loop
			if err == io.EOF {
				break
			}
		}

		// validate json body
		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(request); err != nil {
			result := WSSResponse[ChatResponse]{}
			if vErr, isOk := err.(validator.ValidationErrors); isOk {
				for _, err := range vErr {
					result.ErrorsMSG = append(result.ErrorsMSG, fmt.Sprintf("field '%v' have tag '%v'", err.Field(), err.Tag()))
				}

				if err := ws.Send(c, result); err != nil {
					log.Println(err)
					continue
				}
			} else {
				result.ErrorsMSG = append(result.ErrorsMSG, err.Error())
				if err := ws.Send(c, result); err != nil {
					log.Println(err)
					continue
				}
			}
		}

		// create connection client if not exist
		from, isExist := wsc.clients[request.FromUserID]
		if !isExist {
			client := WebSocketClient{
				ID:         request.FromUserID,
				Connection: c,
			}
			wsc.clients[request.FromUserID] = client
			from = wsc.clients[request.FromUserID]
			log.Printf("client with user_id '%v' is connected", request.FromUserID)

			// cleanup connection if done
			defer func() {
				log.Printf("client with user_id '%v' disconnected", request.FromUserID)
				c.Close()
				delete(wsc.clients, request.FromUserID)
			}()
		}

		// send message
		result := WSSResponse[ChatResponse]{
			ErrorsMSG: []string{},
			Data: ChatResponse{
				FromUserID: request.FromUserID,
				ToUserID:   request.ToUserID,
				Message:    request.Message,
			},
		}

		to, isOnline := wsc.clients[request.ToUserID]
		if !isOnline {
			result := WSSResponse[ChatResponse]{
				ErrorsMSG: []string{fmt.Sprintf("user with id '%v' is not online", request.ToUserID)},
			}
			if err := ws.Send(from.Connection, result); err != nil {
				log.Println(err)
			}
			continue
		}

		if err := ws.Send(to.Connection, result); err != nil {
			log.Println(err)
			continue
		}
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
