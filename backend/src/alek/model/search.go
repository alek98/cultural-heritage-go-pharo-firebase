package model

type Search struct {
	AvgRatingFrom    float64 `json:"avgRatingFrom"`
	AvgRatingTo      float64 `json:"avgRatingTo"`
	ChTypeName       string  `json:"chtypeName"`
	Street           string  `json:"street"`
	City             string  `json:"city"`
	Country          string  `json:"country"`
	Name             string  `json:"name"`
	SortByName       string  `json:"sortByName"`
	SortByRating     string  `json:"sortByRating"`
	SortByChTypeName string  `json:"sortByChTypeName"`

	/*
		sort values, (order by)
		-----------
		ascending
		descending
		ratingAscending
		ratingDescending
		mostCommented
		mostLiked
		mostDisliked
	*/
	// Sort Sort `json:"sort"`
}

// type Sort struct {
// 	SortByName       string `json:"sortByName"`
// 	SortByRating     string `json:"sortByRating"`
// 	SortByChTypeName string `json:"sortByChTypeName"`
// }
