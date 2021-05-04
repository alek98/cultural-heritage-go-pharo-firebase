package model

type User struct {
	Id           string  `json:"id"`
	UserName     string  `json:"userName"`
	Rating       float64 `json:"rating"`
	TotalRatings int64   `json:"totalRatings"`
}
