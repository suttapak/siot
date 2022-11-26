package service

import (
	"context"
	"errors"

	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
	"gorm.io/gorm"
)

type widgetDisplayService struct {
	widget repository.WidgetDisplayRepository
}

func NewWidgetDisplayService(widget repository.WidgetDisplayRepository) WidgetDisplayService {
	return &widgetDisplayService{widget}
}

func (s *widgetDisplayService) Create(ctx context.Context, req *CreateWidgetDisplayRequest) (res *WidgetDisplayResponse, err error) {
	widget, err := s.widget.CreateWidget(ctx, req.Name, req.Description, req.DataType)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}

	res, err = utils.Recast[*WidgetDisplayResponse](widget)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err

}

func (s *widgetDisplayService) Widgets(ctx context.Context) (res []WidgetDisplayResponse, err error) {
	widgets, err := s.widget.FindWidgets(ctx)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[[]WidgetDisplayResponse](widgets)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}

func (s *widgetDisplayService) Widget(ctx context.Context, widgetId uint) (res *WidgetDisplayResponse, err error) {
	widget, err := s.widget.FindWidget(ctx, widgetId)
	if err != nil {
		logs.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrNotFound
		}
		return nil, errs.ErrInternalServerError
	}

	res, err = utils.Recast[*WidgetDisplayResponse](widget)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}

	return res, err
}
