package service

import (
	"context"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	socketio "github.com/googollee/go-socket.io"
	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils/logs"
)

type wsService struct {
	mqtt mqtt.Client

	boxRepo     repository.BoxRepository
	controlRepo repository.ControlRepository
	displayRepo repository.DisplayRepository
}

func NewWsService(mqtt mqtt.Client, boxRepo repository.BoxRepository,
	controlRepo repository.ControlRepository,
	displayRepo repository.DisplayRepository) WsService {

	return &wsService{mqtt, boxRepo, controlRepo, displayRepo}
}

func (s *wsService) RunSub(ctx context.Context, c socketio.Conn, sv *socketio.Server, req SubscriptMessageRequest) {
	_, err := s.boxRepo.FindBoxById(ctx, req.BoxId)
	if err != nil {
		logs.Error(err)
		return
	}
	c.Join(req.Key)
}

func (s *wsService) RunPub(ctx context.Context, req PublishMessageRequest) {

	b, err := s.boxRepo.FindBoxById(ctx, req.BoxId)
	if err != nil {
		logs.Error(err)
		return
	}

	str := fmt.Sprintf("%f", req.Data)
	data := []byte(str)

	t := s.mqtt.Publish(fmt.Sprintf("%v/%v", b.CanPub.CanPublish, req.Key), 0, false, data)
	if t.Error() != nil {
		logs.Error(t.Error())
		return
	}

}
