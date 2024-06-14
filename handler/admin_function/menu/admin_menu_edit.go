package menu

import (
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// AdminMenuEdit godoc
// @Summary AdminMenuEdit
// @Description AdminMenuEdit
// @Tags AdminMenu
// @Accept  json
// @Produce  json
// @Param item_id path string true "Item ID"
// @Param name body string true "Name"
// @Param description body string true "Description"
// @Param price body float64 true "Price"
// @Param image_url body string true "ImageUrl"
// @Param customization_option_1 body string true "CustomizationOption1"
// @Param extra_price_1 body float64 true "ExtraPrice1"
// @Param customization_option_2 body string true "CustomizationOption2"
// @Param extra_price_2 body float64 true "ExtraPrice2"
// @Param customization_option_3 body string true "CustomizationOption3"
// @Param extra_price_3 body float64 true "ExtraPrice3"
// @Param customization_option_4 body string true "CustomizationOption4"
// @Param extra_price_4 body float64 true "ExtraPrice4"
// @Param customization_option_5 body string true "CustomizationOption5"
// @Param extra_price_5 body float64 true "ExtraPrice5"
// @Param customization_option_6 body string true "CustomizationOption6"
// @Param extra_price_6 body float64 true "ExtraPrice6"
// @Param customization_option_7 body string true "CustomizationOption7"
// @Param extra_price_7 body float64 true "ExtraPrice7"
// @Param customization_option_8 body string true "CustomizationOption8"
// @Param extra_price_8 body float64 true "ExtraPrice8"
// @Param customization_option_9 body string true "CustomizationOption9"
// @Param extra_price_9 body float64 true "ExtraPrice9"
// @Param customization_option_10 body string true "CustomizationOption10"
// @Param extra_price_10 body float64 true "ExtraPrice10"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/admin/menu_edit/{item_id} [patch]
func AdminMenuEdit() func(*gin.Context) {
	return func(c *gin.Context) {
		item_id := c.Param("item_id")
		if item_id == "" || item_id == ":item_id" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "không tìm thấy item_id",
			})
			return
		}

		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqMenuItems{}
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, res.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		if len(req.CustomizationOption) != len(req.ExtraPrice) {
			c.JSON(http.StatusBadRequest, res.Response{
				StatusCode: http.StatusBadRequest,
				Message:    "CustomizationOption và ExtraPrice  must have the same number of elements",
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

		menu := dto.MenuItemsDto{
			Name:                req.Name,
			Description:         req.Description,
			Price:               req.Price,
			ImageUrl:            req.ImageUrl,
			CustomizationOption: req.CustomizationOption,
			ExtraPrice:          req.ExtraPrice,
		}

		err := validate.Struct(menu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data := menu.ToPayLoad().ToModel()
		uc := usecases.NewAdminFunctionUseCase()

		err = uc.EditMenu(c.Request.Context(), item_id, data)
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
			Data:       "Sửa món thành công với ID: " + item_id,
		})
	}
}
