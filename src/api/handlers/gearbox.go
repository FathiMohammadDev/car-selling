package handlers

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type GearboxHandler struct {
	service services.GearboxService
}

func NewGearboxHandler(cfg *config.Config) *GearboxHandler {
	return &GearboxHandler{
		*services.NewGearboxService(cfg),
	}
}

// CreateGearbox godoc
// @Summary Create new Gearbox
// @Description Create new Gearbox
// @Tags Gearbox
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateGearboxReq true "CreateGearboxReq"
// @Router /v1/gearbox/ [post]
// @Security AuthBearer
func (h *GearboxHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// CreateGearbox godoc
// @Summary Update Gearbox
// @Description Update Gearbox
// @Tags Gearbox
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.UpdateGearboxReq true "UpdateGearboxReq"
// @Router /v1/gearbox/{id} [put]
// @Security AuthBearer
func (h *GearboxHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteGearbox godoc
// @Summary delete Gearbox
// @Description delete Gearbox
// @Tags Gearbox
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/gearbox/{id} [delete]
// @Security AuthBearer
func (h *GearboxHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)

}

// GetCty godoc
// @Summary Get a Gearbox
// @Description Get a Gearbox
// @Tags Gearbox
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/gearbox/{id} [get]
// @Security AuthBearer
func (h *GearboxHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}
