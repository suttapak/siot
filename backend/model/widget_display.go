package model

type WidgetDisplay struct {
	Model
	// attribute
	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description"`
    DataType    string `json:"dataType"`
    Width       int    `json:"width"`
    Height      int    `json:"height"`
}
