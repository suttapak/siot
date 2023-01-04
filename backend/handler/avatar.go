package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/suttapak/siot-backend/service"
	"github.com/suttapak/siot-backend/utils"
	"github.com/suttapak/siot-backend/utils/logs"
)

type avatarHandler struct {
	avatarServ service.AvatarService
}

func NewAvatarHandler(avatarServ service.AvatarService) *avatarHandler {
	return &avatarHandler{avatarServ}
}

func (h *avatarHandler) Update(c *gin.Context) {
	userId, err := utils.UserId(c)
	if err != nil {
		logs.Error(err)
		c.AbortWithStatusJSON(handleError(err))
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		logs.Error(err)
		c.AbortWithStatusJSON(handleError(err))
		return
	}
	name, url, dst := h.avatarServ.GenerateName(c, file)

	if _, err := os.Stat("./public/asset/images"); os.IsNotExist(err) {
		err := os.MkdirAll("./public/asset/images", 0775)
		if err != nil {
			logs.Error(err)
			c.AbortWithStatusJSON(handleError(err))
			return
		}
	}
	if err := c.SaveUploadedFile(file, dst); err != nil {
		logs.Error(err)
		c.AbortWithStatusJSON(handleError(err))
		return
	}

	body := service.UpdateAvatarRequest{
		UId:    userId,
		Titile: name,
		Url:    url,
	}
	res, err := h.avatarServ.Update(c, body)
	if err != nil {
		c.AbortWithStatusJSON(handleError(err))
		return
	}
	c.JSON(http.StatusCreated, res)
}
