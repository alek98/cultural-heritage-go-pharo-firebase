package model

import "time"

type Search struct {
	AvgRatingFrom float64   `json:"avgRatingFrom"`
	AvgRatingTo   float64   `json:"avgRatingTo"`
	ChTypeName    string    `json:"chtypeName"`
	Street        string    `json:"street"`
	City          string    `json:"city"`
	Country       string    `json:"country"`
	Name          string    `json:"name"`
	DateFrom      time.Time `json:"dateFrom"`
	DateTo        time.Time `json:"dateTo"`

	/*
		sort values
		-----------
		ascending
		descending
		newestFirst
		oldestFirst
		ratingAscending
		ratingDescending
		mostCommented
		mostLiked
		mostDisliked
	*/
	Sort string `json:"sort"`
}
