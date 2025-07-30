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

type CountryHandler struct {
	service services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	return &CountryHandler{
		*services.NewCountryService(cfg),
	}
}

// CreateCountry godoc
// @Summary Create new country
// @Description Create new country
// @Tags Country
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateUpdateCountryReq true "CreateUpdateCountryReq"
// @Router /v1/countries/ [post]
// @Security AuthBearer
func (h *CountryHandler) Create(ctx *gin.Context) {
	req := dto.CreateUpdateCountryReq{}
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

// UpdateCountry godoc
// @Summary Update country
// @Description Update country
// @Tags Country
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateUpdateCountryReq true "CreateUpdateCountryReq"
// @Router /v1/countries/{id} [put]
// @Security AuthBearer
func (h *CountryHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	req := dto.CreateUpdateCountryReq{}

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

// DeleteCountry godoc
// @Summary Delete a country
// @Description Delete a country
// @Tags Country
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/countries/{id} [delete]
// @Security AuthBearer
func (h *CountryHandler) Delete(ctx *gin.Context) {
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

// GetCountry godoc
// @Summary Get a country
// @Description Get a country
// @Tags Country
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/countries/{id} [get]
// @Security AuthBearer
func (h *CountryHandler) GetById(c *gin.Context) {
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
