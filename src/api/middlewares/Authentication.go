package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/FathiMohammadDev/car-selling/api/helpers"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/constants"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	service := services.NewTokenService(cfg)

	return func(ctx *gin.Context) {
		var claims map[string]interface{}
		var err error

		auth := ctx.GetHeader(constants.AuthorizationHeaderKey)
		if auth == "" {
			err = errors.New("token is required")
		} else {
			token := strings.Split(auth, " ")
			claims, err = service.GetClaims(token[1])
		}

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.GenerateBaseResWithErr(
				nil, false, -1, err,
			))
			return
		}

		ctx.Set(constants.UserIdKey, claims[constants.UserIdKey])
		ctx.Set(constants.FirstNameKey, claims[constants.FirstNameKey])
		ctx.Set(constants.LastNameKey, claims[constants.LastNameKey])
		ctx.Set(constants.UsernameKey, claims[constants.UsernameKey])
		ctx.Set(constants.EmailKey, claims[constants.EmailKey])
		ctx.Set(constants.MobileNumberKey, claims[constants.MobileNumberKey])
		ctx.Set(constants.RolesKey, claims[constants.RolesKey])
		ctx.Set(constants.ExpireTimeKey, claims[constants.ExpireTimeKey])

		ctx.Next()

	}
}
