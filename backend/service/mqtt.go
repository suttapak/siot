package service

import (
	"context"
	"time"

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
		Control []struct {
			ID          uint      `json:"id" `
			CreatedAt   time.Time `json:"createdAt"`
			UpdatedAt   time.Time `json:"updatedAt" `
			Name        string    `json:"name" `
			Key         string    `json:"key" `
			Description string    `json:"description" `
			BoxId       uuid.UUID `json:"boxId" `
			ControlData []struct {
				ID        uint      `json:"id" `
				CreatedAt time.Time `json:"createdAt" `
				UpdatedAt time.Time `json:"updatedAt" `
				Data      float64   `json:"data"`
				Label     string    `json:"label"`
				ControlId uint      `json:"controlId" `
			} `json:"controlData"`
		} `json:"control"`
		Display []struct {
			ID          uint      `json:"id"`
			CreatedAt   time.Time `json:"createdAt"`
			UpdatedAt   time.Time `json:"updatedAt"`
			Name        string    `json:"name"`
			Key         string    `json:"key"`
			Description string    `json:"description"`
			BoxId       uuid.UUID `json:"boxId"`
			DisplayData []struct {
				ID        uint      `json:"id"`
				CreatedAt time.Time `json:"createdAt"`
				UpdatedAt time.Time `json:"updatedAt"`
				Data      float64   `json:"data"`
				Label     string    `json:"label"`
				ControlId uint      `json:"controlId"`
			} `json:"controlData"`
		} `json:"display"`
	} `json:"data"`
}

type SubscriptMessageRequest struct {
	BoxId uuid.UUID `json:"boxId"`
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
