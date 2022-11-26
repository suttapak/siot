package handler

import (
	"net/http"

	"github.com/suttapak/siot-backend/utils/errs"
)

func handleError(err error) (code int, e *errs.Error) {
	switch e := err.(type) {
	case *errs.Error:
		return e.Code, e
	default:
		return http.StatusInternalServerError, errs.ErrInternalServerError
	}
}
