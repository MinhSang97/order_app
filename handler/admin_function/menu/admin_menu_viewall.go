package menu

import (
	"github.com/MinhSang97/order_app/payload"
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AdminMenuView godoc
// @Summary AdminMenuView
// @Description AdminMenuView
// @Tags AdminMenu
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/admin/menu_view [get]

func AdminMenuView() func(*gin.Context) {
	return func(c *gin.Context) {

		var Data = payload.MenuItemsPayload{}

		if err := c.ShouldBind(&Data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		uc := usecases.NewAdminFunctionUseCase()
		menuAll, err := uc.GetMenuAll(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, res.Response{
			StatusCode: http.StatusOK,
			Message:    "Xử lý thành công",
			Data:       menuAll,
		})
	}
}
