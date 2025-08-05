package handlers

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type Property struct {
	service services.PropertyService
}

func NewProperty(cfg *config.Config) *Property {
	return &Property{
		*services.NewPropertyService(cfg),
	}
}

// CreatePropertygodoc
// @Summary Create new Property
// @Description Create new Property
// @Tags Property
// @Accept  json
// @Produce  json
// @Param Request body dto.CreatePropertyRequest true "CreatePropertyRequest"
// @Router /v1/property/ [post]
// @Security AuthBearer
func (h *Property) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// UpdataPropertygodoc
// @Summary Update new Property
// @Description Update new Property
// @Tags Property
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePropertyRequest true "UpdatePropertyRequest"
// @Router /v1/property/{id} [put]
// @Security AuthBearer
func (h *Property) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeletePropertygodoc
// @Summary Delete new Property
// @Description Delete new Property
// @Tags Property
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/property/{id} [delete]
// @Security AuthBearer
func (h *Property) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)

}

// GetPropertygodoc
// @Summary Get new Property
// @Description Get new Property
// @Tags Property
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/property/{id} [get]
// @Security AuthBearer
func (h *Property) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}
