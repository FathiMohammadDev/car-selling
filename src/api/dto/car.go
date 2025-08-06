package dto

type CreateCarTypeReq struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type UpdateCarTypeReq struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type CarTypeRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateGearboxReq struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type UpdateGearboxReq struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type GearboxRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCarModelReq struct {
	Name      string `json:"name" binding:"required,min=3,max=15"`
	CompanyId int    `json:"companyId" binding:"required"`
	CarTypeId int    `json:"carTypeId" binding:"required"`
	GearboxId int    `json:"gearboxId" binding:"required"`
}

type UpdateCarModelReq struct {
	Name      string `json:"name,omitempty"`
	CompanyId int    `json:"companyId,omitempty"`
	CarTypeId int    `json:"carTypeId,omitempty"`
	GearboxId int    `json:"gearboxId,omitempty"`
}

type CarModelRes struct {
	Id      int        `json:"id"`
	Name    string     `json:"name"`
	CarType CarTypeRes `json:"carType"`
	Company CompanyRes `json:"company"`
	Gearbox GearboxRes `json:"gearbox"`
}
