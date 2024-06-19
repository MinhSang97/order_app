package order

import (
	"github.com/MinhSang97/order_app/payload"
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AdminOrderHistory godoc
// @Summary AdminOrderHistory
// @Description AdminOrderHistory
// @Tags AdminOrder
// @Accept  json
// @Produce  json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/admin/order_history [get]
func AdminOrderHistory() func(*gin.Context) {
	return func(c *gin.Context) {
		var Data = payload.MenuItemsPayload{}

		if err := c.ShouldBind(&Data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		uc := usecases.NewAdminFunctionUseCase()
		orderAll, err := uc.GetOrderAll(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, res.Response{
			StatusCode: http.StatusOK,
			Message:    "Xử lý thành công",
			Data:       orderAll,
		})

	}
}
