package model

import (
	"encoding/json"
	"log"
	"time"
)

type FeedbackModel struct {
	FeedbackID int       `json:"feedback_id" db:"feedback_id"`
	OrderID    int       `json:"order_id" db:"order_id"`
	Rating     int       `json:"rating" db:"rating"`
	Comment    string    `json:"comment" db:"comment"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

func (c *FeedbackModel) TableName() string {
	return "feedbacks"
}

func (c *FeedbackModel) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}
func (c *FeedbackModel) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
