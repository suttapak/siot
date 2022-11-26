package model

import "github.com/google/uuid"

type CanPublish struct {
	// common
	Model
	// attribute
	CanPublish string `json:"canPublish"`
	// fk
	BoxId uuid.UUID `json:"boxId"`
	// relation
}
