package service

import (
	"context"
	"github.com/google/uuid"
)

type MqttAuthService interface {
	Auth(ctx context.Context, req *MqttAuthRequest) error
	Admin(ctx context.Context, req *MqttAdminRequest) error
	ACLCheckI(ctx context.Context, req *MqttACLRequest) error
}

type MqttAuthRequest struct {
	Id     uuid.UUID `json:"username" binding:"required"`
	Secret string    `json:"password"`
}

type MqttAdminRequest struct {
	UserId uuid.UUID `json:"username" binding:"required"`
}

type MqttACLRequest struct {
	Acc   int       `json:"acc"`
	Id    uuid.UUID `json:"username"`
	Topic string    `json:"topic"`
}
