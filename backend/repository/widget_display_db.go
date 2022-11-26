package repository

import (
	"context"

	"github.com/suttapak/siot-backend/model"
	"gorm.io/gorm"
)

type widgetDisplayRepository struct {
	db *gorm.DB
}

func NewWidgetDisplayRepository(db *gorm.DB) WidgetDisplayRepository {
	widgets := []model.WidgetDisplay{
		{
			Name:        "DLineChart",
			Description: "A line chart have axis X and axis Y.",
			DataType:    "Axis X allow all, Axis Y allow number.",
			Width:       4,
			Height:      2,
		},
		{
			Name:        "DNumber",
			Description: "Box of number can show integer and floating number.",
			DataType:    "Number",
			Width:       1,
			Height:      1,
		},
		{
			Name:        "DCircularPercent",
			Description: "Box of Percent number.",
			DataType:    "Axis X allow all, Axis Y allow number.",
			Width:       2,
			Height:      2,
		},
	}
	db.Create(widgets)
	return &widgetDisplayRepository{db}
}

func (r *widgetDisplayRepository) FindWidget(ctx context.Context, widgetId uint) (widget *model.WidgetDisplay, err error) {
	err = r.db.WithContext(ctx).Where(model.WidgetDisplay{Model: model.Model{ID: widgetId}}).First(&widget).Error
	return widget, err
}

func (r *widgetDisplayRepository) FindWidgets(ctx context.Context) (widget []model.WidgetDisplay, err error) {
	err = r.db.WithContext(ctx).Find(&widget).Error

	return widget, err
}
func (r *widgetDisplayRepository) CreateWidget(ctx context.Context, name, description, dataType string) (widget *model.WidgetDisplay, err error) {
	widget = &model.WidgetDisplay{Name: name, Description: description, DataType: dataType}
	err = r.db.WithContext(ctx).Create(&widget).Error
	return widget, err
}
