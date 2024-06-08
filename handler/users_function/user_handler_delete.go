package handler

import (
	"fmt"
	"net/http"

	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
)

// UsersDelete godoc
// @Summary Delete user
// @Description Delete user
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/users/delete/{user_id} [delete]
func UsersDelete() func(*gin.Context) {
	return func(c *gin.Context) {
		// Trích xuất Bear token từ tiêu đề yêu cầu
		token1 := c.GetHeader("Authorization")
		if token1 == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Missing authorization token",
			})
			return
		}
		fmt.Println("token1: ", token1)

		user_id := c.Param("user_id")

		if user_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "ko tim thay user_id",
			})
			return
		}

		uc := usecases.NewUsersUseCase()
		err := uc.DeleteUsers(c.Request.Context(), user_id)

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
