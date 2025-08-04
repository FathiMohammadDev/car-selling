package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/FathiMohammadDev/car-selling/api/helpers"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
	"github.com/gin-gonic/gin"
)

var logger = logging.NewLogger(config.GetConfig())

func Create[Ti, To any](ctx *gin.Context, caller func(ctx context.Context, req *Ti) (*To, error)) {
	req := new(Ti)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResWithValidationError(
			nil, false, helpers.ValidationError, err,
		))
		return
	}
	res, err := caller(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResWithErr(
			nil, false, helpers.InternalError, err,
		))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.GenerateBaseRes(res, true, helpers.Success))
}

func Update[Ti, To any](ctx *gin.Context, caller func(ctx context.Context, id int, req *Ti) (*To, error)) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	req := new(Ti)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResWithValidationError(
			nil, false, helpers.ValidationError, err,
		))
		return
	}

	res, err := caller(ctx, id, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResWithErr(
			nil, false, helpers.InternalError, err,
		))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.GenerateBaseRes(res, true, helpers.Success))
}

func Delete(ctx *gin.Context, caller func(ctx context.Context, id int) error) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if id == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseRes(nil, false, helpers.ValidationError))
		return
	}

	err := caller(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResWithErr(nil, false, helpers.InternalError, err))
		return
	}
	ctx.JSON(http.StatusOK, helpers.GenerateBaseRes(nil, true, helpers.Success))
}

func GetById[To any](ctx *gin.Context, caller func(ctx context.Context, id int) (*To, error)) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseRes(nil, false, helpers.ValidationError))
		return
	}

	res, err := caller(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResWithErr(nil, false, helpers.InternalError, err))
		return
	}
	ctx.JSON(http.StatusOK, helpers.GenerateBaseRes(res, true, helpers.Success))
}
