package model

type Review struct {
	Id           string  `json:"id"`
	ChId         string  `json:"chId"`
	Content      string  `json:"content"`
	UserName     string  `json:"userName"`
	Rating       float64 `json:"rating"`
	TotalRatings int64   `json:"totalRatings"`
}
