package model

import "github.com/google/uuid"

type CanSubscribe struct {
	Model
	// attribute
	CanSubscribe string `json:"canSubscribe"`
	// fk
	BoxId uuid.UUID `json:"boxId"`
	// relation
}
