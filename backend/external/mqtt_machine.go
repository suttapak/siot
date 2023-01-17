package external

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	socketio "github.com/googollee/go-socket.io"
	"github.com/suttapak/siot-backend/model"
	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/logs"
)

type mqttMachine struct {
	client mqtt.Client

	sv *socketio.Server

	canSubRepo  repository.CanSubRepository
	conRepo     repository.ControlRepository
	conDataRepo repository.ControlDataRepository
	disRepo     repository.DisplayRepository
	disDataRepo repository.DisplayDataRepository
}

func NewMQTTMachine(client mqtt.Client,
	sv *socketio.Server,
	canSubRepo repository.CanSubRepository,
	conRepo repository.ControlRepository,
	conDataRepo repository.ControlDataRepository,
	disRepo repository.DisplayRepository,
	disDataRepo repository.DisplayDataRepository,
) *mqttMachine {
	return &mqttMachine{
		client,
		sv,
		canSubRepo,
		conRepo,
		conDataRepo,
		disRepo,
		disDataRepo,
	}
}

func (m *mqttMachine) MQTTMachine() {
	time.Sleep(5 * time.Second)
	ctx := context.Background()
	t := m.client.Subscribe("#", 0, func(c mqtt.Client, msg mqtt.Message) {

		s := string(msg.Payload())
		s = strings.TrimSpace(s)
		if s == "" {
			fmt.Println("Error: message payload is empty or only whitespace")
			return
		}
		// Try to convert the string to a float64
		data, err := strconv.ParseFloat(s, 64)

		if err != nil || math.IsNaN(data) {
			// Handle error if the message payload is not a number
			fmt.Println("Error: message payload is not a number")
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

		now := time.Now()

		if compareTime(now, control, data) {
			return
		}

		for _, con := range control {
			m.conDataRepo.Crate(ctx, con.ID, data, label)
			continue
		}

		display, err := m.disRepo.FindDisplaysByKey(ctx, canSub.BoxId, key)
		if err != nil {
			logs.Error(err)
			return
		}

		// TODO redis
		for _, dis := range display {
			m.disDataRepo.Crate(ctx, dis.ID, data, label)
			continue
		}
		var cData []model.ControlData
		if len(control) != 0 {
			cData, err = m.conDataRepo.FindByCId(ctx, control[0].ID)
			if err != nil {
				logs.Error(err)
				return
			}
		}
		var dData []model.DisplayData
		if len(display) != 0 {
			dData, err = m.disDataRepo.FindByDisplayId(ctx, display[0].ID)
			if err != nil {
				logs.Error(err)
				return
			}
		}

		temp := Data{Display: dData, Control: cData}
		res, err := utils.Recast[MqttMachineResponse](temp)
		if err != nil {
			logs.Error(err)
			return
		}
		defer m.sv.BroadcastToRoom("", canSub.CanSubscribe+"/"+key, canSub.CanSubscribe+"/"+key, res)
	})

	go func() {
		if t.Error() != nil {
			logs.Error(t.Error())
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

func compareTime(now time.Time, c []model.Control, data float64) bool {
	if len(c) <= 0 {
		return false
	}

	if len(c[0].ControlData) <= 0 {
		return false
	}
	return !now.After(c[0].ControlData[len(c[0].ControlData)-1].CreatedAt.Add(1*time.Second)) && c[0].ControlData[len(c[0].ControlData)-1].Data == data
}
