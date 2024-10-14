package model

type Feedback string

const (
	FeedbackRead     Feedback = "read"
	FeedbackLike     Feedback = "like"
	FeedbackFavorite Feedback = "favorite"
)
