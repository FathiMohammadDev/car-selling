package handlers

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type CarModelPropertyHandler struct {
	service services.CarModelPropertyService
}

func NewCarModelPropertyHandler(cfg *config.Config) *CarModelPropertyHandler {
	return &CarModelPropertyHandler{
		*services.NewCarModelPropertyService(cfg),
	}
}

// CreateCarModelProperty godoc
// @Summary Create new CarModelProperty
// @Description Create new CarModelProperty
// @Tags CarModelProperty
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateCarModelPropertyReq true "CreateCarModelPropertyReq"
// @Router /v1/car-model-property/ [post]
// @Security AuthBearer
func (h *CarModelPropertyHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// CreateCarModelProperty godoc
// @Summary Update CarModelProperty
// @Description Update CarModelProperty
// @Tags CarModelProperty
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelPropertyReq true "UpdateCarModelPropertyReq"
// @Router /v1/car-model-property/{id} [put]
// @Security AuthBearer
func (h *CarModelPropertyHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCarModelProperty godoc
// @Summary delete CarModelProperty
// @Description delete CarModelProperty
// @Tags CarModelProperty
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-model-property/{id} [delete]
// @Security AuthBearer
func (h *CarModelPropertyHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)

}

// GetCty godoc
// @Summary Get a CarModelProperty
// @Description Get a CarModelProperty
// @Tags CarModelProperty
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-model-property/{id} [get]
// @Security AuthBearer
func (h *CarModelPropertyHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}
