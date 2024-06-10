package handler

import (
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UsersGetAddress godoc
// @Summary Users can get address
// @Description Users can get address
// @Tags usersFunction
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/users/get-address/{user_id} [get]
func UsersGetAddress() func(*gin.Context) {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		if user_id == "" || user_id == ":user_id" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "không tìm thấy user_id",
			})
			return
		}

		uc := usecases.NewUsersUseCase()
		data, err := uc.GetAddressUsersFunction(c.Request.Context(), user_id)
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
			Message:    "Xử lý thành công",
			Data:       data,
		})
	}
}
