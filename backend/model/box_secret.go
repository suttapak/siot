package model

import "github.com/google/uuid"

type BoxSecret struct {
	Model
	// attribute
	Secret string `json:"secret"`
	// fk
	BoxId uuid.UUID `json:"boxId"`
	// relation
}
