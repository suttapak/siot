package external

import (
	"fmt"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/suttapak/siot-backend/config"
	"github.com/suttapak/siot-backend/model"
	"github.com/suttapak/siot-backend/utils/logs"
)

func NewMqttClient(conf *config.Configs) mqtt.Client {
	opts := mqtt.NewClientOptions().
		AddBroker(fmt.Sprintf("tcp://%v:%v", conf.Mqtt.Broker, conf.Mqtt.Port)).
		SetUsername(conf.Mqtt.Username).
		SetPassword(conf.Mqtt.Password)
	opts.SetClientID(getRandStr())
	opts.SetCleanSession(true)
	opts.SetConnectionLostHandler(func(c mqtt.Client, err error) {
		panic(err)
	})
	client := mqtt.NewClient(opts)

	if err := client.Connect().Error(); err != nil {
		logs.Error(err)
		panic(err)
	}

	return client
}

func getRandStr() string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	length := 6
	randStr := make([]byte, length)
	rand.Read(randStr)
	for i, b := range randStr {
		randStr[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(randStr)
}

type MqttMessage struct {
	Value float64 `json:"value"`
}

type MqttMachineResponse struct {
	DisplayResponse []DisplayResponse `json:"displayData"`
	ControlResponse []ControlResponse `json:"controlData"`
}

type DisplayResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Data      float64   `json:"data"`
	Label     string    `json:"label"`
	DisplayId uint      `json:"displayId"`
}

type ControlResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Data      float64   `json:"data"`
	Label     string    `json:"label"`
	ControlId uint      `json:"controlId"`
}

type Data struct {
	Control []model.ControlData `json:"controlData"`
	Display []model.DisplayData `json:"displayData"`
}
