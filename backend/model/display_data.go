package model

type DisplayData struct {
	Model
	// attribute
	XNumber int    `json:"XNumber"`
	YNumber int    `json:"YNumber"`
	XString string `json:"XString"`
	YString string `json:"YString"`
	// fk
	DisplayId uint `json:"displayId"`
	// relation
}
