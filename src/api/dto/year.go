package dto

import "time"

type CreateYearReq struct {
	Title string    `json:"Title" binding:"min=4,max=4"`
	Year         int       `json:"year"`
	StartAt      time.Time `json:"startAt"`
	EndAt        time.Time `json:"endAt"`
}

type UpdateYearReq struct {
	Title string    `json:"Title,omitempty" binding:"min=4,max=4"`
	Year         int       `json:"year,omitempty"`
	StartAt      time.Time `json:"startAt,omitempty"`
	EndAt        time.Time `json:"endAt,omitempty"`
}

type YearRes struct {
	Id           int       `json:"id"`
	Title string    `json:"Title,omitempty"`
	Year         int       `json:"year,omitempty"`
	StartAt      time.Time `json:"startAt,omitempty"`
	EndAt        time.Time `json:"endAt,omitempty"`
}

type YearWithoutDateRes struct {
	Id           int    `json:"id"`
	Title string `json:"Title,omitempty"`
	Year         int    `json:"year,omitempty"`
}
