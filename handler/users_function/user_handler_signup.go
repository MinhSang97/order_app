package handler

import (
	"fmt"
	"github.com/MinhSang97/order_app/log"
	"github.com/MinhSang97/order_app/payload"
	"github.com/MinhSang97/order_app/sercurity"
	"github.com/MinhSang97/order_app/usecases"
	usersdto "github.com/MinhSang97/order_app/usecases/dto/users_dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"net/http"
	"strings"
)

// UsersSignUp godoc
// @Summary Users sign up with email and password
// @Description Users sign up with email and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param   name     body   string     true       "Name"
// @Param   email    body   string     true       "Email"
// @Param   pass_word    body   string     true       "PassWord"
// @Param   phone_number    body   string     true       "PhoneNumber"
// @Param   address    body   string     true       "Address"
// @Param   telegram    body   string     true       "Telegram"
// @Param   lat    body   float64     true       "Lat"
// @Param   long    body   float64     true       "Long"
// @Param   ward_id    body   string     true       "WardId"
// @Param   ward_text    body   string     true       "WardText"
// @Param   district_id    body   string     true       "DistrictId"
// @Param   district_text    body   string     true       "DistrictText"
// @Param   province_id    body   string     true       "ProvinceId"
// @Param   province_text    body   string     true       "ProvinceText"
// @Param   national_id    body   string     true       "NationalId"
// @Param   national_text    body   string     true       "NationalText"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/users/sign-up [post]
func UsersSignUp() func(*gin.Context) {
	return func(c *gin.Context) {
		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqSignUp{}
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, res.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}
		fmt.Println("req", req)
		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusForbidden, res.Response{
				StatusCode: http.StatusForbidden,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		PassHash := sercurity.HashAndSalt([]byte(req.PassWord))
		role := payload.USERS.String()

		userUsersId, err := uuid.NewUUID()

		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusForbidden, res.Response{
				StatusCode: http.StatusForbidden,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		userUsers := usersdto.Users{
			UserId:       userUsersId.String(),
			Name:         req.Name,
			PassWord:     PassHash,
			Email:        req.Email,
			Role:         strings.ToLower(role),
			PhoneNumber:  req.PhoneNumber,
			Address:      req.Address,
			Telegram:     req.Telegram,
			Lat:          req.Lat,
			Long:         req.Long,
			WardId:       req.WardId,
			WardText:     req.WardText,
			DistrictId:   req.DistrictId,
			DistrictText: req.DistrictText,
			ProvinceId:   req.ProvinceId,
			ProvinceText: req.ProvinceText,
			NationalId:   req.NationalId,
			NationalText: req.NationalText,
		}

		err = validate.Struct(userUsers)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data := userUsers.ToPayload().ToModel()
		uc := usecases.NewUsersUseCase()

		err = uc.CreateUsers(c.Request.Context(), data)

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
			Data:       "UserID: " + data.UserId,
		})
	}
}
