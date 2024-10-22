package model

type ReviewEventDTO struct {
	EventType    string  `json:"event"`
	Id           int     `json:"id"`
	ReviewerId   int     `json:"reviewer_id"`
	RestaurantId string  `json:"restaurant_id"`
	FoodId       string  `json:"food_id"`
	Rating       float32 `json:"rating"`
	Content      string  `json:"content"`
}

type FavoriteEventDTO struct {
	EventType string `json:"event"`
	UserId    int    `json:"user_id"`
	FoodId    string `json:"food_id"`
}
