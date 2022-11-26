package repository

import (
	"context"

	"github.com/suttapak/siot-backend/model"
)

type WidgetDisplayRepository interface {
	FindWidget(ctx context.Context, widgetId uint) (widget *model.WidgetDisplay, err error)
	FindWidgets(ctx context.Context) (widget []model.WidgetDisplay, err error)
	CreateWidget(ctx context.Context, name, description, dataType string) (widget *model.WidgetDisplay, err error)
}
