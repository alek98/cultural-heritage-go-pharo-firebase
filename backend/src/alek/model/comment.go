package model

type Comment struct {
	Id       string `json:"id"`
	ReviewId string `json:"reviewId"`
	Content  string `json:"content"`
	UserName string `json:"userName"`
}
