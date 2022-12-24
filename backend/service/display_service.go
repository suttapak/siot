package service

import (
	"context"
	"errors"
	"strings"

	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
	"gorm.io/gorm"
)

type displayService struct {
	boxRepo     repository.BoxRepository
	displayRepo repository.DisplayRepository
	layoutRepo  repository.LayoutRepository
	widgetRepo  repository.WidgetDisplayRepository
}

func NewDisplayService(boxRepo repository.BoxRepository,
	displayRepo repository.DisplayRepository, layoutRepo repository.LayoutRepository,
	widgetRepo repository.WidgetDisplayRepository) DisplayService {

	return &displayService{boxRepo, displayRepo, layoutRepo, widgetRepo}

}

func (s *displayService) Create(ctx context.Context, req *CreateDisplayRequest) (res *DisplayResponse, err error) {
	// check box
	_, err = s.boxRepo.FindIsMember(ctx, req.BoxId, req.UserId)
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
	layoutBody := repository.CreateLayoutRequest{
		I: req.Layout.I,
		X: req.Layout.X,
		Y: req.Layout.Y,
		W: widget.Width,
		H: widget.Height,
	}

	layout, err := s.layoutRepo.Create(ctx, &layoutBody)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}

	// create display
	displayBody := repository.CreateDisplayRequest{
		Name:        req.Name,
		Key:         strings.ToUpper(req.Key),
		Description: req.Description,
		BoxId:       req.BoxId,
		LayoutId:    layout.ID,
		WidgetId:    req.Widget.ID,
	}
	// create display
	display, err := s.displayRepo.Create(ctx, &displayBody)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}

	display, err = s.displayRepo.FindDisplay(ctx, &repository.FindDisplayRequest{DisplayId: display.ID})
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	// map response
	res, err = utils.Recast[*DisplayResponse](display)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}

func (s *displayService) FindDisplay(ctx context.Context, req *FindDisplaysRequest) (res []DisplayResponse, err error) {
	_, err = s.boxRepo.FindIsMember(ctx, req.BoxId, req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error(err)
			return nil, errs.ErrUnauthorized
		}
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}

	displays, err := s.displayRepo.FindDisplays(ctx, &repository.FindDisplaysRequest{
		BoxId: req.BoxId,
	})
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}

	res, err = utils.Recast[[]DisplayResponse](displays)

	return res, err
}

func (s *displayService) Update(ctx context.Context, dId uint, req *UpdateDisplayRequest) (res *DisplayResponse, err error) {
	body := repository.UpdateDisplayRequest(*req)
	d, err := s.displayRepo.Update(ctx, dId, body)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[*DisplayResponse](d)
	if err != nil {
		logs.Error(err)
		return nil, errs.ErrInternalServerError
	}
	return res, err
}
func (s *displayService) Delete(ctx context.Context, dId uint) error {
	err := s.displayRepo.Delete(ctx, dId)
	if err != nil {
		logs.Error(err)
		return errs.ErrInternalServerError
	}
	return nil
}
