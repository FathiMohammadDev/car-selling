package handlers

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type PropertyCategory struct {
	service services.PropertyCategoryService
}

func NewPropertyCategory(cfg *config.Config) *PropertyCategory {
	return &PropertyCategory{
		*services.NewPropertyCategoryService(cfg),
	}
}

// CreatePropertyCategorygodoc
// @Summary Create new PropertyCategory
// @Description Create new PropertyCategory
// @Tags PropertyCategory
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.CreatePropertyCategoryRequest true "CreatePropertyCategoryRequest"
// @Router /v1/property-category/ [post]
// @Security AuthBearer
func (h *PropertyCategory) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// UpdataPropertyCategorygodoc
// @Summary Update new PropertyCategory
// @Description Update new PropertyCategory
// @Tags PropertyCategory
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePropertyCategoryRequest true "UpdatePropertyCategoryRequest"
// @Router /v1/property-category/{id} [put]
// @Security AuthBearer
func (h *PropertyCategory) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeletePropertyCategorygodoc
// @Summary Delete new PropertyCategory
// @Description Delete new PropertyCategory
// @Tags PropertyCategory
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/property-category/{id} [delete]
// @Security AuthBearer
func (h *PropertyCategory) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)

}

// GetPropertyCategorygodoc
// @Summary Get new PropertyCategory
// @Description Get new PropertyCategory
// @Tags PropertyCategory
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/property-category/{id} [get]
// @Security AuthBearer
func (h *PropertyCategory) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}
