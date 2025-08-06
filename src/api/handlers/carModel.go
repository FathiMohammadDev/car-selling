package handlers

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type CarModelHandler struct {
	service services.CarModelService
}

func NewCarModelHandler(cfg *config.Config) *CarModelHandler {
	return &CarModelHandler{
		*services.NewCarModelService(cfg),
	}
}

// CreateCarModel godoc
// @Summary Create new CarModel
// @Description Create new CarModel
// @Tags CarModel
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateCarModelReq true "CreateCarModelReq"
// @Router /v1/car-model/ [post]
// @Security AuthBearer
func (h *CarModelHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// CreateCarModel godoc
// @Summary Update CarModel
// @Description Update CarModel
// @Tags CarModel
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelReq true "UpdateCarModelReq"
// @Router /v1/car-model/{id} [put]
// @Security AuthBearer
func (h *CarModelHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCarModel godoc
// @Summary delete CarModel
// @Description delete CarModel
// @Tags CarModel
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-model/{id} [delete]
// @Security AuthBearer
func (h *CarModelHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)

}

// GetCty godoc
// @Summary Get a CarModel
// @Description Get a CarModel
// @Tags CarModel
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-model/{id} [get]
// @Security AuthBearer
func (h *CarModelHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}
