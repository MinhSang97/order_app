// package handler
//
// import (
//
//	"github.com/MinhSang97/order_app/sercurity"
//	"github.com/MinhSang97/order_app/usecases"
//	"github.com/MinhSang97/order_app/usecases/dto/users_dto"
//	"github.com/MinhSang97/order_app/usecases/req"
//	"github.com/MinhSang97/order_app/usecases/res"
//	"github.com/gin-gonic/gin"
//	"github.com/go-playground/validator/v10"
//	"net/http"
//
// )
//
//	func UsersSignIn() func(*gin.Context) {
//		return func(c *gin.Context) {
//			var validate *validator.Validate
//			validate = validator.New(validator.WithRequiredStructEnabled())
//			req := req.ReqSignIn{}
//			if err := c.ShouldBind(&req); err != nil {
//				c.JSON(http.StatusBadRequest, res.Response{
//					StatusCode: http.StatusBadRequest,
//					Message:    err.Error(),
//					Data:       nil,
//				})
//				return
//			}
//
//			if err := validate.Struct(req); err != nil {
//				c.JSON(http.StatusForbidden, res.Response{
//					StatusCode: http.StatusForbidden,
//					Message:    err.Error(),
//					Data:       nil,
//				})
//				return
//			}
//
//			userUsers := users_dto.ReqSignIn{
//				PassWord: req.PassWord,
//				Email:    req.Email,
//			}
//
//			err := validate.Struct(userUsers)
//
//			if err != nil {
//				c.JSON(http.StatusBadRequest, gin.H{
//					"error": err.Error(),
//				})
//				return
//			}
//
//			err = validate.Struct(userUsers)
//
//			if err != nil {
//				c.JSON(http.StatusBadRequest, gin.H{
//					"error": err.Error(),
//				})
//				return
//			}
//
//			data := userUsers.ToPayload().ToModel()
//			uc := usecases.NewUsersUseCase()
//
//			//admin := Data.ToPayload().ToModel()
//			//uc := usecases.NewAdminUseCase()
//
//			usersPass, err := uc.GetUsers(c.Request.Context(), data)
//			if err != nil {
//				c.JSON(http.StatusUnauthorized, res.Response{
//					StatusCode: http.StatusUnauthorized,
//					Message:    err.Error(),
//					Data:       nil,
//				})
//				return
//			}
//
//			// check pass
//			isTheSame := sercurity.ComparePasswords(usersPass.PassWord, []byte(req.PassWord))
//			if !isTheSame {
//				c.JSON(http.StatusUnauthorized, res.Response{
//					StatusCode: http.StatusUnauthorized,
//					Message:    "Đăng nhập thất bại",
//					Data:       nil,
//				})
//				return
//			}
//
//			//gen token
//			token, err := sercurity.GenTokenUsers(users_dto.Users{})
//			if err != nil {
//				c.JSON(http.StatusInternalServerError, res.Response{
//					StatusCode: http.StatusInternalServerError,
//					Message:    err.Error(),
//					Data:       nil,
//				})
//				return
//			}
//
//			c.JSON(http.StatusOK, res.Response{
//				StatusCode: http.StatusOK,
//				Message:    "Xử lý thành công",
//				Data:       token,
//			})
//
//		}
//	}
package handler

import (
	"github.com/MinhSang97/order_app/sercurity"
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/dto/users_dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func UsersSignIn() func(*gin.Context) {
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

		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusForbidden, res.Response{
				StatusCode: http.StatusForbidden,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}
		//gen token
		token, err := sercurity.GenTokenUsers(users_dto.Users{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, res.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}
		PassHash := sercurity.HashAndSalt([]byte(req.PassWord))
		users := users_dto.ReqSignIn{
			PassWord: PassHash,
			Email:    req.Email,
			Token:    token,
		}

		err = validate.Struct(users)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data := users.ToPayload().ToModel()
		uc := usecases.NewUsersUseCase()

		usersPass, err := uc.GetUsers(c.Request.Context(), data)
		if err != nil {
			c.JSON(http.StatusUnauthorized, res.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		//check pass
		isTheSame := sercurity.ComparePasswords(usersPass.PassWord, []byte(req.PassWord))
		if !isTheSame {
			c.JSON(http.StatusUnauthorized, res.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Đăng nhập thất bại. Kiểm tra lại email password",
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, res.Response{
			StatusCode: http.StatusOK,
			Message:    "Xử lý thành công",
			Data:       usersPass.Token,
		})

	}
}
