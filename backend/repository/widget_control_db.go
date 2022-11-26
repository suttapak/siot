package repository

import (
	"context"

	"github.com/suttapak/siot-backend/model"
	"gorm.io/gorm"
)

type widgetControlRepository struct {
	db *gorm.DB
}

func NewWidgetControlRepository(db *gorm.DB) WidgetControlRepository {
	widget := []model.WidgetControl{
		{Name: "CButton", Description: "A Button have two state are on and off.", DataType: "boolean", Width: 1, Height: 1},
		{Name: "CButtonNumber", Description: "A Button it's contain number", DataType: "number", Width: 1, Height: 2},
		{Name: "CSlider", Description: "A Slider to slide to set state number", DataType: "number", Width: 2, Height: 1},
		{Name: "CSwitch", Description: "A Switch have two state are on and off", DataType: "boolean", Width: 1, Height: 2},
	}
	db.Create(&widget)
	return &widgetControlRepository{db}
}

func (r *widgetControlRepository) FindWidget(ctx context.Context, widgetId uint) (widget *model.WidgetControl, err error) {
	err = r.db.WithContext(ctx).Where(model.WidgetControl{Model: model.Model{ID: widgetId}}).First(&widget).Error
	return widget, err
}

func (r *widgetControlRepository) FindWidgets(ctx context.Context) (widget []model.WidgetControl, err error) {
	err = r.db.WithContext(ctx).Find(&widget).Error
	return widget, err
}
func (r *widgetControlRepository) Create(ctx context.Context, name, description, dataType string) (widget *model.WidgetControl, err error) {
	widget = &model.WidgetControl{Name: name, Description: description, DataType: dataType}
	err = r.db.WithContext(ctx).Create(&widget).Error
	return widget, err
}
