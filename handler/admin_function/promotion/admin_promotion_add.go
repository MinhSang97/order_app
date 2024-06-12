package promotion

import (
	"fmt"
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
	"net/http"
)

// AdminPromotionAdd godoc
// @Summary AdminPromotionAdd
// @Description AdminPromotionAdd
// @Tags AdminPromotion
// @Accept  json
// @Produce  json
// @Param title body string true "Title"
// @Param description body string true "Description"
// @Param code body string true "Code"
// @Param discount_percentage body float64 true "DiscountPercentage"
// @Param valid_from body string true "ValidFrom"
// @Param valid_to body string true "ValidTo"
// @Param promotion_id body string true "PromotionID"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/admin/promotion_add [post]
func AdminPromotionAdd() func(*gin.Context) {
	return func(c *gin.Context) {
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

		id_str, err := uuid.NewV4()
		if err != nil {
			fmt.Printf("Lỗi khi tạo DiscountCodesID: %v\n", err)
			return
		}
		discount_code_id := id_str.String()

		menu := dto.DiscountCodesDto{
			DiscountCodeID:     discount_code_id,
			Title:              req.Title,
			Description:        req.Description,
			Code:               req.Code,
			DiscountPercentage: req.DiscountPercentage,
			ValidFrom:          req.ValidFrom,
			ValidTo:            req.ValidTo,
			PromotionID:        req.PromotionID,
		}

		err = validate.Struct(menu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data := menu.ToPayLoad().ToModel()
		uc := usecases.NewAdminFunctionUseCase()
		discountCodesAdd, err := uc.AddDiscount(c.Request.Context(), data)
		if err != nil {
			c.JSON(http.StatusConflict, res.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, res.Response{
			StatusCode: http.StatusOK,
			Message:    "Xử lý thành công",
			Data:       "Thêm max giảm giá thành công với ID: " + discountCodesAdd.DiscountCodeID,
		})
	}
}
