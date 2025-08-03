package handlers

import (
	"net/http"
	"strconv"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/api/helpers"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	service services.CityService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	return &CityHandler{
		*services.NewCityService(cfg),
	}
}

// CreateCity godoc
// @Summary Create new city
// @Description Create new city
// @Tags City
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateCityReq true "CreateCityReq"
// @Router /v1/cities/ [post]
// @Security AuthBearer
func (h *CityHandler) Create(ctx *gin.Context) {
	req := dto.CreateCityReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResWithValidationError(
			nil, false, 409, err,
		))
		return
	}
	res, err := h.service.Create(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResWithErr(
			nil, false, 409, err,
		))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.GenerateBaseRes(res, true, 0))
}

// CreateCity godoc
// @Summary Update city
// @Description Update city
// @Tags City
// @Accept  json
// @Produce  json
// @Param Request body dto.UpdateCityReq true "UpdateCityReq"
// @Router /v1/cities/ [put]
// @Security AuthBearer
func (h *CityHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	req := dto.UpdateCityReq{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResWithValidationError(
			nil, false, 409, err,
		))
		return
	}

	res, err := h.service.Update(ctx, id, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResWithErr(
			nil, false, 409, err,
		))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.GenerateBaseRes(res, true, 0))
}

// DeleteCity godoc
// @Summary delete city
// @Description delete city
// @Tags City
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/cities/ [delete]
// @Security AuthBearer
func (h *CityHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if id == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseRes(nil, false, 121))
		return
	}

	err := h.service.Delete(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResWithErr(nil, false, 121, err))
		return
	}
	ctx.JSON(http.StatusOK, helpers.GenerateBaseRes(nil, true, 0))

}

// GetCty godoc
// @Summary Get a city
// @Description Get a city
// @Tags City
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/cities/{id} [get]
// @Security AuthBearer
func (h *CityHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseRes(nil, false, 121))
		return
	}

	res, err := h.service.GetById(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResWithErr(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusOK, helpers.GenerateBaseRes(res, true, 0))
}
