package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suttapak/siot-backend/utils/errs"
)

var ResponseOk = gin.H{
	"status": "OK",
}

func handleError(err error) (code int, e *errs.Error) {
	switch e := err.(type) {
	case *errs.Error:
		return e.Code, e
	default:
		return http.StatusInternalServerError, errs.ErrInternalServerError
	}
}
