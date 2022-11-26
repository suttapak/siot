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

type widgetControlService struct {
	widget repository.WidgetControlRepository
}

func NewWidgetControlService(widget repository.WidgetControlRepository) WidgetControlService {
	return &widgetControlService{widget}
}

func (s *widgetControlService) Create(ctx context.Context, req *CreateWidgetControlRequest) (res *WidgetControlResponse, err error) {
	widget, err := s.widget.Create(ctx, req.Name, req.Description, req.DataType)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[*WidgetControlResponse](widget)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}
func (s *widgetControlService) Widgets(ctx context.Context) (res []WidgetControlResponse, err error) {
	widgets, err := s.widget.FindWidgets(ctx)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}

	res, err = utils.Recast[[]WidgetControlResponse](widgets)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}

func (s *widgetControlService) Widget(ctx context.Context, widgetId uint) (res *WidgetControlResponse, err error) {
	widget, err := s.widget.FindWidget(ctx, widgetId)
	if err != nil {
		logs.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrNotFound
		}
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[*WidgetControlResponse](widget)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}
