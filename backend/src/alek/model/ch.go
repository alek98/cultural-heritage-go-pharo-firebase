package model

type Ch struct {
	Id          string   `json:"id"`
	AvgRating   float32  `json:"avgRating"`
	ChType      ChType   `json:"chtype"`
	Description string   `json:"description"`
	Location    Location `json:"location"`
	Name        string   `json:"name"`
}
