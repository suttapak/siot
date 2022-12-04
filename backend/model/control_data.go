package model

type ControlData struct {
	Model
	Data      float64 `json:"data"`
	Label     string  `json:"label"`
	ControlId uint    `json:"controlId"`
}
