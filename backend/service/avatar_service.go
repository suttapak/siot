package service

import (
	"context"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/suttapak/siot-backend/repository"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/errs"
)

const (
	FileDst = "./public/asset/images/"
	BaseUrl = "/asset/images/"
)

type avatarService struct {
	avatarRepo repository.AvatarRepository
}

func NewAvatarService(avatarRepo repository.AvatarRepository) AvatarService {
	return &avatarService{avatarRepo}
}

func (s *avatarService) Update(ctx context.Context, req UpdateAvatarRequest) (res *AvatarReponse, err error) {
	avatar, err := s.avatarRepo.Create(ctx, repository.CreateAvatarRequest{
		UId:   req.UId,
		Title: req.Titile,
		Url:   req.Url,
	})

	if err != nil {
		return nil, errs.ErrInternalServerError
	}
	res, err = utils.Recast[*AvatarReponse](avatar)
	if err != nil {
		return nil, errs.ErrInternalServerError
	}

	return res, err
}

func (s *avatarService) GenerateName(ctx context.Context, File *multipart.FileHeader) (name, url, dst string) {
	// Retrieve file information
	extension := filepath.Ext(File.Filename)
	name = File.Filename
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	nameFile := uuid.New().String() + extension
	url = BaseUrl + nameFile
	dst = FileDst + nameFile
	return name, url, dst
}
