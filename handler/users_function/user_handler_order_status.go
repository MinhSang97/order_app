package handler

import (
	"github.com/MinhSang97/order_app/usecases/dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func UsersOrderStatus() func(*gin.Context) {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		if user_id == "" || user_id == ":user_id" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "không tìm thấy user_id",
			})
			return
		}

		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqOrderStatus{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, res.Response{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid request body",
				Data:       nil,
			})
			return
		}

		orderStatus := dto.OrderDto{
			OrderID: req.OrderID,
			Status:  req.Status,
		}
		err := validate.Struct(orderStatus)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

	}
}
