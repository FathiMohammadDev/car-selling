package handlers

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type CarModelColorlHandler struct {
	service services.CarModelColorService
}

func NewCarModelColorlHandler(cfg *config.Config) *CarModelColorlHandler {
	return &CarModelColorlHandler{
		*services.NewCarModelColorService(cfg),
	}
}

// CreateCarModelColorl godoc
// @Summary Create new CarModelColor
// @Description Create new CarModelColor
// @Tags CarModelColor
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateCarModelColorReq true "CreateCarModelColorReq"
// @Router /v1/car-model-color/ [post]
// @Security AuthBearer
func (h *CarModelColorlHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// CreateCarModelColorl godoc
// @Summary Update CarModelColor
// @Description Update CarModelColor
// @Tags CarModelColor
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelColorReq true "UpdateCarModelColorReq"
// @Router /v1/car-model-color/{id} [put]
// @Security AuthBearer
func (h *CarModelColorlHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCarModelColor godoc
// @Summary delete CarModelColor
// @Description delete CarModelColor
// @Tags CarModelColor
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-model-color/{id} [delete]
// @Security AuthBearer
func (h *CarModelColorlHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)

}

// GetCty godoc
// @Summary Get a CarModelColor
// @Description Get a CarModelColor
// @Tags CarModelColor
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-model-color/{id} [get]
// @Security AuthBearer
func (h *CarModelColorlHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}
