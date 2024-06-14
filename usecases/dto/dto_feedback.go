package dto

import (
	"github.com/MinhSang97/order_app/payload"
	"time"
)

type FeedbackDto struct {
	FeedbackID int       `json:"feedback_id"`
	OrderID    int       `json:"order_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
}

func (c *FeedbackDto) ToPayload() *payload.FeedbackPayload {
	feedbackPayload := &payload.FeedbackPayload{
		FeedbackID: c.FeedbackID,
		OrderID:    c.OrderID,
		Rating:     c.Rating,
		Comment:    c.Comment,
		CreatedAt:  c.CreatedAt,
	}
	return feedbackPayload
}
