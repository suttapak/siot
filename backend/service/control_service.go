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

type controlService struct {
	boxRepo     repository.BoxRepository
	controlRepo repository.ControlRepository
	layoutRepo  repository.LayoutRepository
	widgetRepo  repository.WidgetControlRepository
}

func NewControlService(boxRepo repository.BoxRepository, controlRepo repository.ControlRepository, layoutRepo repository.LayoutRepository, widgetRepo repository.WidgetControlRepository) ControlService {
	return &controlService{boxRepo, controlRepo, layoutRepo, widgetRepo}
}

func (s *controlService) Create(ctx context.Context, req *CreateControlRequest) (res *ControlResponse, err error) {
	_, err = s.boxRepo.FindBox(ctx, req.BoxId, req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error(err)
			return nil, errs.ErrUnauthorized
		}

		logs.Error(err)
		return nil, errs.ErrBadRequest
	}
	// check widget
	widget, err := s.widgetRepo.FindWidget(ctx, req.Widget.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error(err)
			return nil, errs.ErrBadRequest
		}
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	// TODO check key in box
	// create layout
    layoutBodyTemp := repository.CreateLayoutRequest{
        I: req.Layout.I,
        X: req.Layout.X,
        Y:req.Layout.Y,
        W: widget.Width,
        H:widget.Height,
    }

    layout, err := s.layoutRepo.Create(ctx, &layoutBodyTemp)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	controlBody := repository.CreateControlRequest{
		Name:        req.Name,
		Key:         req.Key,
		Description: req.Description,
		BoxId:       req.BoxId,
		LayoutId:    layout.ID,
		WidgetId:    req.Widget.ID,
	}
	control, err := s.controlRepo.Create(ctx, &controlBody)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}

	control, err = s.controlRepo.FindControl(ctx, &repository.FindControlRequest{
		ControlId: control.ID,
	})
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[*ControlResponse](control)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	// create control
	return res, err
}
func (s *controlService) FindControls(ctx context.Context, req *FindControlsRequest) (res []ControlResponse, err error) {
	_, err = s.boxRepo.FindBox(ctx, req.BoxId, req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error(err)
			return nil, errs.ErrUnauthorized
		}

		logs.Error(err)
		return nil, errs.ErrBadRequest
	}
	controls, err := s.controlRepo.FindControls(ctx, &repository.FindControlsRequest{BoxId: req.BoxId})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error(err)
			return nil, errs.ErrUnauthorized
		}
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[[]ControlResponse](controls)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}
