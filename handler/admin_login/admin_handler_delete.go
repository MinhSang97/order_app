package handler

import (
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminDelete() func(*gin.Context) {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")

		if user_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "ko tim thay user_id",
			})
			return
		}

		uc := usecases.NewAdminUseCase()

		err := uc.DeleteAdmin(c.Request.Context(), user_id)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, res.Response{
				StatusCode: http.StatusUnprocessableEntity,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, res.Response{
			StatusCode: http.StatusOK,
			Message:    "Xoá thành công",
			Data:       map[string]interface{}{"user_id": user_id},
		})
	}
}
