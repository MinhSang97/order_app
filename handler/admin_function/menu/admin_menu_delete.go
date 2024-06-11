package menu

import (
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminMenuDelete() func(*gin.Context) {
	return func(c *gin.Context) {
		item_id := c.Param("item_id")
		if item_id == "" || item_id == ":item_id" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "không tìm thấy item_id",
			})
			return
		}

		uc := usecases.NewAdminFunctionUseCase()

		err := uc.DeleteMenu(c.Request.Context(), item_id)

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
			Data:       map[string]interface{}{"item_id": item_id},
		})
	}
}
