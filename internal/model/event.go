package model

type ReviewEventDTO struct {
	EventType    string  `json:"event"`
	Id           int     `json:"id"`
	ReviewerId   int     `json:"reviewer_id"`
	RestaurantId int     `json:"restaurant"`
	Rating       float64 `json:"rating"`
	Content      string  `json:"content"`
}

type FavoriteEventDTO struct {
	EventType string `json:"event"`
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
}
