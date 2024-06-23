package handler

import (
	sercurity2 "github.com/MinhSang97/order_app/pkg/sercurity"
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/dto/admin_dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type DataRes struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	PassWord    string `json:"password"`
	Token       string `json:"token"`
	UserID      string `json:"user_id"`
}

// AdminSignIn godoc
// @Summary Admin can sign in
// @Description Admin can sign in
// @Tags admin
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param pass_word body string true "PassWord"
// @Param phone_number body string true "PhoneNumber"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/admin/sign_in [post]
func AdminSignIn() func(*gin.Context) {
	return func(c *gin.Context) {
		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqSignIn{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, res.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		//vl
		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusForbidden, res.Response{
				StatusCode: http.StatusForbidden,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}
		//gen token
		token, err := sercurity2.GenTokenAdmin(admin_dto.Admin{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, res.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}
		PassHash := sercurity2.HashAndSalt([]byte(req.PassWord))
		userAdmin := admin_dto.ReqSignIn{
			PassWord:    PassHash,
			Email:       req.Email,
			Token:       token,
			PhoneNumber: req.PhoneNumber,
		}

		err = validate.Struct(userAdmin)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data := userAdmin.ToPayload().ToModel()
		uc := usecases.NewAdminUseCase()

		adminPass, err := uc.GetAdmin(c.Request.Context(), data)
		if err != nil {
			c.JSON(http.StatusUnauthorized, res.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		//check pass
		isTheSame := sercurity2.ComparePasswords(adminPass.PassWord, []byte(req.PassWord))
		if !isTheSame {
			c.JSON(http.StatusUnauthorized, res.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Đăng nhập thất bại. Kiểm tra lại email password",
				Data:       nil,
			})
			return
		}

		dataRes := DataRes{
			Email:       adminPass.Email,
			PhoneNumber: adminPass.PhoneNumber,
			PassWord:    req.PassWord,
			Token:       adminPass.Token,
			UserID:      adminPass.UserID,
		}
		c.JSON(http.StatusOK, res.Response{
			StatusCode: http.StatusOK,
			Message:    "Xử lý thành công",
			Data:       dataRes,
		})

	}
}
