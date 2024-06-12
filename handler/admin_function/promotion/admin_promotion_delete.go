package promotion

import (
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AdminPromotionDelete godoc
// @Summary AdminPromotionDelete
// @Description AdminPromotionDelete
// @Tags AdminPromotion
// @Accept  json
// @Produce  json
// @Param discount_code_id path string true "Discount Code ID"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/admin/promotion_delete/{discount_code_id} [delete]
func AdminPromotionDelete() func(*gin.Context) {
	return func(c *gin.Context) {
		discount_code_id := c.Param("discount_code_id")
		if discount_code_id == "" || discount_code_id == ":discount_code_id" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "không tìm thấy discount_code_id",
			})
			return
		}

		uc := usecases.NewAdminFunctionUseCase()

		err := uc.DeleteDiscount(c.Request.Context(), discount_code_id)

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
			Data:       map[string]interface{}{"discount_code_id": discount_code_id},
		})
	}
}
