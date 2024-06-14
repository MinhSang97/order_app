package req

type ReqFeedback struct {
	//FeedbackID string `json:"feedback_id" validate:"required"`
	OrderID int    `json:"order_id" validate:"required"`
	Rating  int    `json:"rating" validate:"required"`
	Comment string `json:"comment" validate:"required"`
}

type ReqFeedbackView struct {
	OrderID int `json:"order_id" validate:"required"`
}
