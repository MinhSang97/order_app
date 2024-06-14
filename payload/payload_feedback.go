package payload

import (
	"encoding/json"
	"github.com/MinhSang97/order_app/model"
	"log"
	"time"
)

type FeedbackPayload struct {
	FeedbackID int       `json:"feedback_id"`
	OrderID    int       `json:"order_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
}

func (c *FeedbackPayload) ToModel() *model.FeedbackModel {
	feedbackPayload := &model.FeedbackModel{
		FeedbackID: c.FeedbackID,
		OrderID:    c.OrderID,
		Rating:     c.Rating,
		Comment:    c.Comment,
		CreatedAt:  c.CreatedAt,
	}
	return feedbackPayload
}

func (c *FeedbackPayload) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
