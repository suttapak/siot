package handler

import (
	"context"
	"encoding/json"

	socketio "github.com/googollee/go-socket.io"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
)

type websocketHandler struct {
	wsServ         service.WsService
	socket         *socketio.Server
	userOnlineServ service.UserOnlineService
}

func NewWsHandler(wsServ service.WsService, socket *socketio.Server, userOnlineServ service.UserOnlineService) *websocketHandler {
	return &websocketHandler{wsServ, socket, userOnlineServ}
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

func (h *websocketHandler) UserOnline(s socketio.Conn, msg interface{}) {
	s.Join("userOnline")
	ctx := context.Background()
	_, err := h.userOnlineServ.Increment(ctx)
	if err != nil {
		s.Close()
	}
	res, err := h.userOnlineServ.Decrement(ctx)
	if err != nil {
		s.Close()
		return
	}
	h.socket.BroadcastToRoom("", "userOnline", "userOnline", res)
}

func (h *websocketHandler) OnConnect(s socketio.Conn) error {
	logs.Info("User connect : " + s.ID())
	ctx := context.Background()
	res, err := h.userOnlineServ.Increment(ctx)
	if err != nil {
		s.Close()
		return errs.ErrInternalServerError
	}
	h.socket.BroadcastToRoom("", "userOnline", "userOnline", res)
	return nil
}

func (h *websocketHandler) OnDisconnect(s socketio.Conn, reason string) {
	logs.Info("Socket disconnect : " + reason)
	ctx := context.Background()
	res, err := h.userOnlineServ.Decrement(ctx)
	if err != nil {
		s.Close()
		return
	}
	h.socket.BroadcastToRoom("", "userOnline", "userOnline", res)
}
