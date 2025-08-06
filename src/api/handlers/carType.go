package handlers

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type CarTypeHandler struct {
	service services.CarTypeService
}

func NewCarTypeHandler(cfg *config.Config) *CarTypeHandler {
	return &CarTypeHandler{
		*services.NewCarTypeService(cfg),
	}
}

// CreateCarType godoc
// @Summary Create new CarType
// @Description Create new CarType
// @Tags CarType
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateCarTypeReq true "CreateCarTypeReq"
// @Router /v1/car-type/ [post]
// @Security AuthBearer
func (h *CarTypeHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// CreateCarType godoc
// @Summary Update CarType
// @Description Update CarType
// @Tags CarType
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarTypeReq true "UpdateCarTypeReq"
// @Router /v1/car-type/{id} [put]
// @Security AuthBearer
func (h *CarTypeHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCarType godoc
// @Summary delete CarType
// @Description delete CarType
// @Tags CarType
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-type/{id} [delete]
// @Security AuthBearer
func (h *CarTypeHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)

}

// GetCty godoc
// @Summary Get a CarType
// @Description Get a CarType
// @Tags CarType
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-type/{id} [get]
// @Security AuthBearer
func (h *CarTypeHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}
