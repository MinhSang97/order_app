package handler

import (
	"github.com/MinhSang97/order_app/sercurity"
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UsersUpdateResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ChangePassWord godoc
// @Summary Users can change password
// @Description Users can change password
// @Tags otp
// @Accept json
// @Produce json
// @Param otp path string true "OTP"
// @Param email body string true "Email"
// @Param pass_word_new body string true "PassWordNew"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/users/change_password/{otp} [patch]
func ChangePassWord() func(*gin.Context) {
	return func(c *gin.Context) {
		otp_code := c.Param("otp")
		if otp_code == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "không tìm thấy OTP code",
			})
			return
		}

		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqChangePassword{}

		if err := c.ShouldBindBodyWithJSON(&req); err != nil {
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
		otpData := dto.OtpDto{
			Email: req.Email,
			Otp:   otp_code,
		}
		err := validate.Struct(otpData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		PassHashNew := sercurity.HashAndSalt([]byte(req.PassWordNew))
		changePassword := dto.OtpDto{
			Email:       req.Email,
			PassWordNew: PassHashNew,
		}

		err = validate.Struct(changePassword)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data := changePassword.ToPayload().ToModel()
		uc := usecases.NewOtpUseCase()

		err = uc.ChangePassword(c.Request.Context(), otp_code, data)

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
			Data:       "Thay đổi password thành công cho email " + data.Email,
		})
	}
}
