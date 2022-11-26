package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/suttapak/siot-backend/config"
	"github.com/suttapak/siot-backend/utils/errs"
	"github.com/suttapak/siot-backend/utils/logs"
)

type _JWTmiddleware struct {
	conf *config.Configs
}

func NewJWTWare(conf *config.Configs) *_JWTmiddleware {
	return &_JWTmiddleware{conf}
}

func (m *_JWTmiddleware) JWTWare(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		logs.Error("header is not exist")
		ctx.AbortWithStatusJSON(handleError(errs.ErrUnauthorized))
		return
	}
	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		logs.Error("bad token")
		ctx.AbortWithStatusJSON(handleError(errs.ErrUnauthorized))
		return
	}
	token, err := parseToken(jwtToken[1], m.conf)
	if err != nil {
		logs.Error(err)
		ctx.AbortWithStatusJSON(handleError(err))
		return

	}
	user, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logs.Error("bad claims")
		ctx.AbortWithStatusJSON(handleError(errs.ErrInternalServerError))
		return
	}
	userId, ok := user["userId"]
	if !ok {
		logs.Error("can not convert user id")
		ctx.AbortWithStatusJSON(handleError(errs.ErrInternalServerError))
		return
	}
	email, ok := user["email"].(string)
	if !ok {
		logs.Error("can not convert email")
		ctx.AbortWithStatusJSON(handleError(errs.ErrInternalServerError))
		return
	}
	ctx.Set("userId", userId)
	ctx.Set("email", email)
	ctx.Next()
}

func parseToken(jwtToken string, conf *config.Configs) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.ErrBadRequest
		}
		return []byte(conf.JWT.Secret), nil
	})

	if err != nil {
		return nil, errs.NewError(http.StatusBadRequest, "bad jwt token")
	}

	return token, nil
}

func handleError(err error) (code int, e *errs.Error) {
	switch e := err.(type) {
	case *errs.Error:
		return e.Code, e
	default:
		return http.StatusInternalServerError, errs.ErrInternalServerError
	}
}
