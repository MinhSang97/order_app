package handler

import (
	"fmt"
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
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
		fmt.Println("user_id: ", user_id)

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

		data := orderStatus.ToPayload().ToModel()
		uc := usecases.NewUsersUseCase()
		err = uc.StatusOrderUserOrder(c.Request.Context(), user_id, data)
		if err != nil {
			c.JSON(http.StatusConflict, res.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, res.Response{
			StatusCode: http.StatusOK,
			Message:    "Cập nhật thành công",
			Data:       "Status của order: " + strconv.FormatInt(req.OrderID, 10) + " " + req.Status,
		})
	}
}
