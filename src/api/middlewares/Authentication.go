package middlewares

import (
	"errors"
	"fmt"
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

func Authorization(validRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Keys) == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helpers.GenerateBaseRes(
				nil, false, -301,
			))
			return
		}
		rolesVal := ctx.Keys[constants.RolesKey]
		fmt.Println(rolesVal)
		if rolesVal == nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helpers.GenerateBaseRes(nil, false, -302))
			return
		}
		roles := rolesVal.([]interface{})
		val := map[string]int{}
		for _, item := range roles {
			val[item.(string)] = 0
		}

		for _, item := range validRoles {
			if _, ok := val[item]; ok {
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatusJSON(http.StatusForbidden, helpers.GenerateBaseRes(nil, false, -303))

	}
}
