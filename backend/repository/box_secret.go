package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/model"
)

type BoxSecretRepository interface {
	Create(ctx context.Context, secret string, boxId uuid.UUID) (boxSecret *model.BoxSecret, err error)
}
