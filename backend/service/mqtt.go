package service

import (
	"context"

	"github.com/google/uuid"
	socketio "github.com/googollee/go-socket.io"
	"github.com/suttapak/siot-backend/model"
)

type WsService interface {
	RunSub(ctx context.Context, s socketio.Conn, sv *socketio.Server, req SubscriptMessageRequest)
	RunPub(ctx context.Context, req PublishMessageRequest)
}

type SubscriptMessageResponse struct {
	CanSub string `json:"canSub"`
	Data   struct {
		Control []ControlResponse `json:"control"`
		Display []DisplayResponse `json:"display"`
	} `json:"data"`
}

type SubscriptMessageRequest struct {
	BoxId uuid.UUID `json:"boxId"`
	Key   string    `json:"key"`
}

type PublishMessageRequest struct {
	Data  float64   `json:"data"`
	Key   string    `json:"key"`
	BoxId uuid.UUID `json:"boxId"`
}

type MockSubMsgRes struct {
	CanSub string `json:"canSub"`
	Data   Data   `json:"data"`
}

type Data struct {
	Control []model.Control `json:"control"`
	Display []model.Display `json:"display"`
}
