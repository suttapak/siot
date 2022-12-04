package external

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/suttapak/siot-backend/model"
	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils/logs"
)

type mqttMachine struct {
	client mqtt.Client

	canSubRepo  repository.CanSubRepository
	conRepo     repository.ControlRepository
	conDataRepo repository.ControlDataRepository
	disRepo     repository.DisplayRepository
	disDataRepo repository.DisplayDataRepository
}

func NewMQTTMachine(client mqtt.Client,
	canSubRepo repository.CanSubRepository,
	conRepo repository.ControlRepository,
	conDataRepo repository.ControlDataRepository,
	disRepo repository.DisplayRepository,
	disDataRepo repository.DisplayDataRepository) *mqttMachine {
	return &mqttMachine{
		client,
		canSubRepo,
		conRepo,
		conDataRepo,
		disRepo,
		disDataRepo,
	}
}

func (m *mqttMachine) MQTTMachine() {
	time.Sleep(3 * time.Second)
	ctx := context.Background()
	t := m.client.Subscribe("#", 0, func(c mqtt.Client, msg mqtt.Message) {

		body := MqttMessage{}
		err := json.Unmarshal(msg.Payload(), &body)
		if err != nil {
			logs.Error(err)
			return
		}
		topic, key := m.xTopic(msg)
		canSub, err := m.canSubRepo.GetCanSubByTopic(ctx, topic)
		if err != nil {
			logs.Error(err)
			return
		}
		control, err := m.conRepo.FindControlsByKey(ctx, canSub.BoxId, key)
		if err != nil {
			logs.Error(err)
			return
		}
		// TODO redis
		year, month, day := time.Now().Date()
		label := fmt.Sprintf("%v %v %v", day, month, year)
		go func(contorl []model.Control) {
			for _, con := range control {
				if len(con.ControlData) < 1 {
					m.conDataRepo.Crate(ctx, con.ID, body.Value, label)
					continue
				}
				if con.ControlData[len(con.ControlData)-1].Data != body.Value {
					m.conDataRepo.Crate(ctx, con.ID, body.Value, label)
				}

			}

		}(control)

		display, err := m.disRepo.FindDisplaysByKey(ctx, canSub.BoxId, key)
		if err != nil {
			logs.Error(err)
			return
		}

		go func(display []model.Display) {
			// TODO redis
			for _, dis := range display {
				if len(dis.DisplayData) < 1 {
					m.disDataRepo.Crate(ctx, dis.ID, body.Value, label)
					continue
				}
				if dis.DisplayData[len(dis.DisplayData)-1].Data != body.Value {
					m.disDataRepo.Crate(ctx, dis.ID, body.Value, label)
				}

			}
		}(display)

	})

	go func() {
		if t.Error() != nil {
			logs.Error(t.Error())
		}

		if t.Wait() {
			logs.Info("wait")

		}
	}()
}

func (m *mqttMachine) xTopic(msg mqtt.Message) (topic, key string) {
	slices := strings.Split(msg.Topic(), "/")
	if len(slices) < 2 {
		return "", ""
	}
	return slices[0], slices[1]
}
