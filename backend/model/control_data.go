package model

type ControlData struct {
	Model
	// attribute
	XNumber int    `json:"XNumber"`
	YNumber int    `json:"YNumber"`
	XString string `json:"XString"`
	YString string `json:"YString"`
	// fk
	ControlId uint `json:"controlId"`
	// relation
}
