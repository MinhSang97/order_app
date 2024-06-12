package promotion

import (
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// AdminDiscountCodeEdit godoc
// @Summary AdminDiscountCodeEdit
// @Description AdminDiscountCodeEdit
// @Tags AdminPromotion
// @Accept  json
// @Produce  json
// @Param discount_code_id path string true "Discount Code ID"
// @Param title body string true "Title"
// @Param description body string true "Description"
// @Param code body string true "Code"
// @Param discount_percentage body int true "DiscountPercentage"
// @Param valid_from body string true "ValidFrom"
// @Param valid_to body string true "ValidTo"
// @Param promotion_id body string true "PromotionID"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/admin/discount_code_edit/{discount_code_id} [patch]
func AdminDiscountCodeEdit() func(*gin.Context) {
	return func(c *gin.Context) {
		discount_code_id := c.Param("discount_code_id")
		if discount_code_id == "" || discount_code_id == ":discount_code_id" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "không tìm thấy discount_code_id",
			})
			return
		}

		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqDiscountCodes{}
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, res.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusForbidden, res.Response{
				StatusCode: http.StatusForbidden,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}
		discountCodes := dto.DiscountCodesDto{
			Title:              req.Title,
			Description:        req.Description,
			Code:               req.Code,
			DiscountPercentage: req.DiscountPercentage,
			ValidFrom:          req.ValidFrom,
			ValidTo:            req.ValidTo,
			PromotionID:        req.PromotionID,
		}

		err := validate.Struct(discountCodes)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data := discountCodes.ToPayLoad().ToModel()
		uc := usecases.NewAdminFunctionUseCase()

		err = uc.EditDiscount(c.Request.Context(), discount_code_id, data)

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
			Message:    "Xử lý cập nhật thành công",
			Data:       "Sửa mã giảm giá thành công với ID: " + discount_code_id,
		})

	}
}
