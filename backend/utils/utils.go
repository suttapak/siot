package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserId(ctx *gin.Context) (uid uuid.UUID, err error) {
	uidStr := ctx.GetString("userId")

	uid, err = uuid.Parse(uidStr)
	return uid, err
}

func Recast[R any](data any) (R, error) {
	var result R
	b, err := json.Marshal(data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return result, err
	}
	return result, err
}

func Recasts[R any](data ...any) (R, error) {
	var result R
	b, err := json.Marshal(data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return result, err
	}
	return result, err
}
