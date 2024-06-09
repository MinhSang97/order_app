package handler

import (
	"github.com/MinhSang97/order_app/pkg/otp"
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var otpStore = make(map[string]string)

// SendOTP godoc
// @Summary Send OTP to email
// @Description Send OTP to email
// @Tags otp
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/users/send_otp [post]
func SendOTP() func(*gin.Context) {
	return func(c *gin.Context) {
		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqOTP{}

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

		GenOtp := otp.GenerateOTP()
		otpStore[req.Email] = GenOtp

		otpData := dto.OtpDto{
			Email: req.Email,
			Otp:   GenOtp,
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

		err = uc.SendOtp(c.Request.Context(), data)
		if err != nil {
			c.JSON(http.StatusConflict, res.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		if err := otp.SendEmail(req.Email, GenOtp); err != nil {
			c.JSON(http.StatusInternalServerError, res.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, res.Response{
			StatusCode: http.StatusOK,
			Message:    "Xử lý thành công",
			Data:       "OTP sent to email " + req.Email,
		})
	}
}
