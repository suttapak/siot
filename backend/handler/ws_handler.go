package handler

import (
	"context"
	"encoding/json"

	socketio "github.com/googollee/go-socket.io"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils/logs"
)

type websocketHandler struct {
	// chan service <-
	wsServ service.WsService
	socket *socketio.Server
}

func NewWsHandler(wsServ service.WsService, socket *socketio.Server) *websocketHandler {
	return &websocketHandler{wsServ, socket}
}

func (h *websocketHandler) Subscript(s socketio.Conn, msg any) {
	ctx := context.Background()
	jsonByte, err := json.Marshal(msg)
	if err != nil {
		logs.Error(err)
		s.Close()
		return
	}
	body := service.SubscriptMessageRequest{}
	err = json.Unmarshal(jsonByte, &body)
	if err != nil {
		logs.Error(err)
		s.Close()
		return
	}
	h.wsServ.RunSub(ctx, s, h.socket, body)
}
func (h *websocketHandler) Publish(s socketio.Conn, msg interface{}) {
	ctx := context.Background()
	jsonByte, err := json.Marshal(msg)
	if err != nil {
		logs.Error(err)
		s.Close()
		return
	}
	body := service.PublishMessageRequest{}
	err = json.Unmarshal(jsonByte, &body)
	if err != nil {
		logs.Error(err)
		s.Close()
		return
	}
	h.wsServ.RunPub(ctx, body)

}
