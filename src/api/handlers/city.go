package handlers

import (
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
	Create(ctx, h.service.Create)
}

// CreateCity godoc
// @Summary Update city
// @Description Update city
// @Tags City
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCityReq true "UpdateCityReq"
// @Router /v1/cities/{id} [put]
// @Security AuthBearer
func (h *CityHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCity godoc
// @Summary delete city
// @Description delete city
// @Tags City
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/cities/{id} [delete]
// @Security AuthBearer
func (h *CityHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)

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
func (h *CityHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}
