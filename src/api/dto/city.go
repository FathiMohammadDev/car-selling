package dto

type CreateCityReq struct {
	Name      string `json:"name" binding:"required,alpha,min=3,max=20"`
	CountryId int    `json:"countryId" binding:"required"`
}

type UpdateCityReq struct {
	Name      string `json:"name,omitempty" binding:"alpha,min=3,max=20"`
	CountryId int    `json:"countryId,omitempty"`
}
type CityRes struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}
