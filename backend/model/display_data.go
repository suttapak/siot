package model

type DisplayData struct {
	Model
	Data      float64 `json:"data"`
	Label     string  `json:"label"`
	DisplayId uint    `json:"displayId"`
	// relation
}
