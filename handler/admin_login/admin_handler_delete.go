package handler

import (
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AdminDelete godoc
// @Summary Admin can delete user
// @Description Admin can delete user
// @Tags admin
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/admin/delete/{user_id} [delete]
func AdminDelete() func(*gin.Context) {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		if user_id == "" || user_id == ":user_id" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "không tìm thấy user_id",
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
