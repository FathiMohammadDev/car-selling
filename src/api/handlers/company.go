package handlers

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	service services.CompanyService
}

func NewCompanyHandler(cfg *config.Config) *CompanyHandler {
	return &CompanyHandler{
		*services.NewCompanyService(cfg),
	}
}

// CreateCompany godoc
// @Summary Create new Company
// @Description Create new Company
// @Tags Company
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateCompanyReq true "CreateCompanyReq"
// @Router /v1/company/ [post]
// @Security AuthBearer
func (h *CompanyHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// CreateCompany godoc
// @Summary Update Company
// @Description Update Company
// @Tags Company
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCompanyReq true "UpdateCompanyReq"
// @Router /v1/company/{id} [put]
// @Security AuthBearer
func (h *CompanyHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCompany godoc
// @Summary delete Company
// @Description delete Company
// @Tags Company
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/company/{id} [delete]
// @Security AuthBearer
func (h *CompanyHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)

}

// GetCty godoc
// @Summary Get a Company
// @Description Get a Company
// @Tags Company
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/company/{id} [get]
// @Security AuthBearer
func (h *CompanyHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}
