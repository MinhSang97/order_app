package handler

import (
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/dto/users_dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// UsersChangeAddressDefault godoc
// @Summary Users can change address default
// @Description Users can change address default
// @Tags usersFunction
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param name body string true "Name"
// @Param address body string true "Address"
// @Param phone_number body string true "PhoneNumber"
// @Param type body string true "Type"
// @Param address_default body string true "AddressDefault"
// @Param lat body float64 true "Lat"
// @Param long body float64 true "Long"
// @Param ward_id body string true "WardId"
// @Param ward_text body string true "WardText"
// @Param district_id body string true "DistrictId"
// @Param district_text body string true "DistrictText"
// @Param province_id body string true "ProvinceId"
// @Param province_text body string true "ProvinceText"
// @Param national_id body string true "NationalId"
// @Param national_text body string true "NationalText"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/users/change-address-default/{user_id} [patch]
func UsersChangeAddressDefault() func(*gin.Context) {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		if user_id == "" || user_id == ":user_id" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "không tìm thấy user_id",
			})
			return
		}

		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqAddress{}
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, res.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		userAddress := users_dto.UsersAddressDto{
			Address:        req.Address,
			Name:           req.Name,
			PhoneNumber:    req.PhoneNumber,
			TypeAddress:    req.TypeAddress,
			AddressDefault: req.AddressDefault,
			Lat:            req.Lat,
			Long:           req.Long,
			WardId:         req.WardId,
			WardText:       req.WardText,
			DistrictId:     req.DistrictId,
			DistrictText:   req.DistrictText,
			ProvinceId:     req.ProvinceId,
			ProvinceText:   req.ProvinceText,
			NationalId:     req.NationalId,
			NationalText:   req.NationalText,
		}
		err := validate.Struct(userAddress)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		data := userAddress.ToPayload().ToModel()
		uc := usecases.NewUsersUseCase()
		err = uc.DefaultAddressUsersFunction(c.Request.Context(), user_id, data)
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
			Data:       "UserID: " + user_id + " đã thêm địa chỉ mặc định thành công",
		})
	}
}
