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