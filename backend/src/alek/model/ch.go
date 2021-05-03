package model

type Ch struct {
	Id          string   `json:"id"`
	AvgRating   float64  `json:"avgRating"`
	ChType      ChType   `json:"chtype"`
	Description string   `json:"description"`
	Location    Location `json:"location"`
	Name        string   `json:"name"`
	Likes       int64    `json:"likes"`
	Dislikes    int64    `json:"dislikes"`
}
