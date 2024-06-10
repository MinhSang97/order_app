package member

import (
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AdminMemberDelete godoc
// @Summary Admin can delete member
// @Description Admin can delete member
// @Tags adminfunction
// @Accept json
// @Produce json
// @Param email path string true "Email"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/admin/member_delete/{email} [delete]
func AdminMemberDelete() func(*gin.Context) {
	return func(c *gin.Context) {
		email := c.Param("email")
		if email == "" || email == ":email" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "không tìm thấy email address",
			})
			return
		}

		uc := usecases.NewAdminFunctionUseCase()

		err := uc.DeleteUsers(c.Request.Context(), email)

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
			Data:       map[string]interface{}{"email": email},
		})
	}
}
