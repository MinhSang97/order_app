package handler

import (
	"github.com/MinhSang97/order_app/sercurity"
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/dto/admin_dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func AdminSignIn() func(*gin.Context) {
	return func(c *gin.Context) {
		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqSignIn{}
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
		userAdmin := admin_dto.ReqSignIn{
			PassWord: req.PassWord,
			Email:    req.Email,
		}

		err := validate.Struct(userAdmin)

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
		isTheSame := sercurity.ComparePasswords(adminPass.PassWord, []byte(req.PassWord))
		if !isTheSame {
			c.JSON(http.StatusUnauthorized, res.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Đăng nhập thất bại",
				Data:       nil,
			})
			return
		}

		//gen token
		token, err := sercurity.GenTokenAdmin(admin_dto.Admin{})
		if err != nil {
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
			Data:       token,
		})

	}
}
