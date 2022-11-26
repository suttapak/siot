package model

import "github.com/google/uuid"

type Control struct {
	Model
	// attribute
	Name        string `json:"name"`
	Key         string `json:"key"`
	Description string `json:"description"`
	// fk
	BoxId    uuid.UUID     `json:"boxId"`
	LayoutId uint          `json:"layoutId"`
	WidgetId uint          `json:"widgetId"`
	Widget   WidgetControl `json:"widget"`
	// relation
	ControlData []ControlData `gorm:"foreignKey:ControlId" json:"controlData"`
	Layout      Layout        `json:"layout"`
}
