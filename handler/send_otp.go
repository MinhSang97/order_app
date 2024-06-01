package handler

import (
	"fmt"
	"github.com/MinhSang97/order_app/log"
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gopkg.in/gomail.v2"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var otpStore = make(map[string]string)

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

		otp := generateOTP()
		otpStore[req.Email] = otp

		otpData := dto.OtpDto{
			Email: req.Email,
			Otp:   otp,
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
		if err := sendEmail(req.Email, otp); err != nil {
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

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%04d", rand.Intn(1000000))
}
func sendEmail(to string, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USERNAME"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", fmt.Sprintf("Your OTP code is: %s", otp))

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Printf("Invalid SMTP port: %v", err)
		return err
	}

	d := gomail.NewDialer(smtpHost, port, smtpUsername, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		// Log the error
		log.Printf("Failed to send email to %s: %v", to, err)
		return err
	}
	return nil
}
