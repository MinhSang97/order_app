package menu

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

// AdminMenuAdd godoc
// @Summary AdminMenuAdd
// @Description AdminMenuAdd
// @Tags AdminMenu
// @Accept  json
// @Produce  json
// @Param name body string true "Name"
// @Param description body string true "Description"
// @Param price body float64 true "Price"
// @Param image_url body string true "ImageUrl"
// @Param customization_option1 body string true "CustomizationOption1"
// @Param extra_price1 body float64 true "ExtraPrice1"
// @Param customization_option2 body string true "CustomizationOption2"
// @Param extra_price2 body float64 true "ExtraPrice2"
// @Param customization_option3 body string true "CustomizationOption3"
// @Param extra_price3 body float64 true "ExtraPrice3"
// @Param customization_option4 body string true "CustomizationOption4"
// @Param extra_price4 body float64 true "ExtraPrice4"
// @Param customization_option5 body string true "CustomizationOption5"
// @Param extra_price5 body float64 true "ExtraPrice5"
// @Param customization_option6 body string true "CustomizationOption6"
// @Param extra_price6 body float64 true "ExtraPrice6"
// @Param customization_option7 body string true "CustomizationOption7"
// @Param extra_price7 body float64 true "ExtraPrice7"
// @Param customization_option8 body string true "CustomizationOption8"
// @Param extra_price8 body float64 true "ExtraPrice8"
// @Param customization_option9 body string true "CustomizationOption9"
// @Param extra_price9 body float64 true "ExtraPrice9"
// @Param customization_option10 body string true "CustomizationOption10"
// @Param extra_price10 body float64 true "ExtraPrice10"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/admin/menu_add [post]
func AdminMenuAdd() func(*gin.Context) {
	return func(c *gin.Context) {
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
			fmt.Printf("Lỗi khi tạo UUID: %v\n", err)
			return
		}
		item_id := id_str.String()

		menu := dto.MenuItemsDto{
			ItemID:                item_id,
			Name:                  req.Name,
			Description:           req.Description,
			Price:                 req.Price,
			ImageUrl:              req.ImageUrl,
			CustomizationOption1:  req.CustomizationOption1,
			ExtraPrice1:           req.ExtraPrice1,
			CustomizationOption2:  req.CustomizationOption2,
			ExtraPrice2:           req.ExtraPrice2,
			CustomizationOption3:  req.CustomizationOption3,
			ExtraPrice3:           req.ExtraPrice3,
			CustomizationOption4:  req.CustomizationOption4,
			ExtraPrice4:           req.ExtraPrice4,
			CustomizationOption5:  req.CustomizationOption5,
			ExtraPrice5:           req.ExtraPrice5,
			CustomizationOption6:  req.CustomizationOption6,
			ExtraPrice6:           req.ExtraPrice6,
			CustomizationOption7:  req.CustomizationOption7,
			ExtraPrice7:           req.ExtraPrice7,
			CustomizationOption8:  req.CustomizationOption8,
			ExtraPrice8:           req.ExtraPrice8,
			CustomizationOption9:  req.CustomizationOption9,
			ExtraPrice9:           req.ExtraPrice9,
			CustomizationOption10: req.CustomizationOption10,
			ExtraPrice10:          req.ExtraPrice10,
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
		menuAdd, err := uc.AddMenu(c.Request.Context(), data)
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
			Data:       "Thêm món thành công với ID: " + menuAdd.ItemID,
		})
	}
}
