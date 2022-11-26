package repository

import (
	"context"
	"github.com/suttapak/siot-backend/model"
)

type WidgetControlRepository interface {
	FindWidget(ctx context.Context, widgetId uint) (widget *model.WidgetControl, err error)
	FindWidgets(ctx context.Context) (widget []model.WidgetControl, err error)
	Create(ctx context.Context, name, description, dataType string) (widget *model.WidgetControl, err error)
}
