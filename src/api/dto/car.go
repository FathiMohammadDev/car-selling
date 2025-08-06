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
	Id             int                `json:"id"`
	Name           string             `json:"name"`
	CarType        CarTypeRes         `json:"carType"`
	Company        CompanyRes         `json:"company"`
	Gearbox        GearboxRes         `json:"gearbox"`
	CarModelColors []CarModelColorRes `json:"carModelColors,omitempty"`
	CarModelYears  []CarModelYearRes  `json:"carModelYears,omitempty"`
}

type CreateCarModelColorReq struct {
	CarModelId int `json:"carModelId" binding:"required"`
	ColorId    int `json:"colorId" binding:"required"`
}

type UpdateCarModelColorReq struct {
	CarModelId int `json:"carModelId,omitempty"`
	ColorId    int `json:"colorId,omitempty"`
}

type CarModelColorRes struct {
	Id    int      `json:"id"`
	Color ColorRes `json:"color,omitempty"`
}

type CreateCarModelYearReq struct {
	CarModelId    int `json:"carModelId" binding:"required"`
	PersianYearId int `json:"persianYearId" binding:"required"`
}

type UpdateCarModelYearReq struct {
	CarModelId    int `json:"carModelId,omitempty"`
	PersianYearId int `json:"persianYearId,omitempty"`
}

type CarModelYearRes struct {
	Id          int                `json:"id"`
	PersianYear YearWithoutDateRes `json:"persianYear,omitempty"`
	CarModelId  int                `json:"carModelId,omitempty"`
}
