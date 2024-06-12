package promotion

import (
	"github.com/MinhSang97/order_app/payload"
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AdminPromotionView godoc
// @Summary AdminPromotionView
// @Description AdminPromotionView
// @Tags AdminPromotion
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/admin/promotion_view [get]
func AdminPromotionView() func(*gin.Context) {
	return func(c *gin.Context) {

		var Data = payload.DiscountCodesPayload{}

		if err := c.ShouldBind(&Data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		uc := usecases.NewAdminFunctionUseCase()
		discountCodesAll, err := uc.GetDiscountAll(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, res.Response{
			StatusCode: http.StatusOK,
			Message:    "Xử lý thành công",
			Data:       discountCodesAll,
		})
	}
}
