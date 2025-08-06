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
