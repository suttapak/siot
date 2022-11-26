package model

import "github.com/google/uuid"

type Display struct {
	Model
	// attribute
	Name        string `json:"name"`
	Key         string `json:"key"`
	Description string `json:"description"`
	// fk
	BoxId    uuid.UUID     `json:"boxId"`
	LayoutId uint          `json:"layoutId"`
	WidgetId uint          `json:"widgetId"`
	Widget   WidgetDisplay `json:"widget"`
	// relation
	DisplayData []DisplayData `gorm:"foreignKey:DisplayId" json:"displayData"`
	Layout      Layout
}
