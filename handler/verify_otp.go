package handler

import (
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// VerifiOTP godoc
// @Summary Users can verify OTP
// @Description Users can verify OTP
// @Tags otp
// @Accept json
// @Produce json
// @Param otp path string true "OTP"
// @Param email body string true "Email"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/users/verify_otp/{otp} [patch]
func VerifiOTP() func(*gin.Context) {
	return func(c *gin.Context) {
		otp_code := c.Param("otp_code")
		if otp_code == "" || otp_code == ":otp_code" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "không tìm thấy otp_code",
			})
			return
		}

		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqOTP{}

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

		data := otpData.ToPayload().ToModel()
		uc := usecases.NewOtpUseCase()

		err = uc.VerifyOtp(c.Request.Context(), data)
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
			Data:       "OTP verified successfully",
		})
		//var Data = admin_payload.AdminFunctionPayload{}
		//
		//if err := c.ShouldBind(&Data); err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"error": err.Error(),
		//	})
		//	return
		//}
		//
		//uc := usecases.NewAdminFunctionUseCase()
		//usersall, err := uc.GetAll(c.Request.Context())
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{
		//		"error": err.Error(),
		//	})
		//	return
		//}

		//c.JSON(http.StatusOK, res.Response{
		//	StatusCode: http.StatusOK,
		//	Message:"Xử lý thành công",
		//	Data:       "usersall",
		//})

	}
}
