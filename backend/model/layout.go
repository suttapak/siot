package model

type Layout struct {
	Model
	// attribute
	I string `json:"i"`
	X int    `json:"x"`
	Y int    `json:"y"`
	W int    `json:"w"`
	H int    `json:"h"`
	// fk
	// relation
}
