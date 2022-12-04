package service

import (
	"context"
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	socketio "github.com/googollee/go-socket.io"
	"github.com/suttapak/siot-backend/external"
	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils"
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
	b, err := s.boxRepo.FindBoxById(ctx, req.BoxId)
	if err != nil {
		logs.Error(err)
		return
	}
	res := s.getData(ctx, req.BoxId, b.CanSub.CanSubscribe)
	c.Join(b.CanSub.CanSubscribe)
	sv.BroadcastToRoom("", b.CanSub.CanSubscribe, b.CanSub.CanSubscribe, res.Data)

	t := s.mqtt.Subscribe(fmt.Sprintf("%v/#", b.CanSub.CanSubscribe), 1, func(client mqtt.Client, message mqtt.Message) {
		res := s.getData(ctx, req.BoxId, b.CanSub.CanSubscribe)
		sv.BroadcastToRoom("", b.CanSub.CanSubscribe, b.CanSub.CanSubscribe, res.Data)
	})

	if t.Error() != nil {
		logs.Error(t.Error())
	}
}

func (s *wsService) RunPub(ctx context.Context, req PublishMessageRequest) {

	b, err := s.boxRepo.FindBoxById(ctx, req.BoxId)
	if err != nil {
		logs.Error(err)
		return
	}
	msg := external.MqttMessage{
		Value: req.Data,
	}
	msgByte, err := json.Marshal(msg)
	if err != nil {
		logs.Error(err)
		return
	}
	t := s.mqtt.Publish(fmt.Sprintf("%v/%v", b.CanPub.CanPublish, req.Key), 1, false, msgByte)
	if t.Error() != nil {
		logs.Error(t.Error())
		return
	}

}

func (s *wsService) getData(ctx context.Context, boxId uuid.UUID, canSub string) *SubscriptMessageResponse {
	control, err := s.controlRepo.FindControls(ctx, &repository.FindControlsRequest{BoxId: boxId})
	if err != nil {
		logs.Error(err)
		return nil
	}
	display, err := s.displayRepo.FindDisplays(ctx, &repository.FindDisplaysRequest{BoxId: boxId})
	if err != nil {
		logs.Error(err)
		return nil
	}
	temp := MockSubMsgRes{CanSub: canSub, Data: Data{control, display}}
	res, err := utils.Recast[SubscriptMessageResponse](temp)
	if err != nil {
		logs.Error(err)
		return nil
	}
	return &res
}
