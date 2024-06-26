package member

import (
	"github.com/MinhSang97/order_app/payload/admin_payload"
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AdminMemberView godoc
// @Summary Admin can view all member
// @Description Admin can view all member
// @Tags adminfunction
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/admin/member_viewall [get]
func AdminMemberView() func(*gin.Context) {
	return func(c *gin.Context) {

		var Data = admin_payload.AdminFunctionPayload{}

		if err := c.ShouldBind(&Data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		uc := usecases.NewAdminFunctionUseCase()
		usersall, err := uc.GetAll(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, res.Response{
			StatusCode: http.StatusOK,
			Message:    "Xử lý thành công",
			Data:       usersall,
		})

	}
}
