package handler

import (
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
)

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

		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusForbidden, res.Response{
				StatusCode: http.StatusForbidden,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		PassHash := sercurity.HashAndSalt([]byte(req.PassWord))
		role := payload.EMPLOYEE.String()

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
		//
		//token1, ok := c.Get("token")
		//fmt.Println(token1)
		//if !ok {
		//	// Handle error: the token was not found in the context
		//}
		//token2, ok := token1.(*jwt.Token)
		//if !ok {
		//	// Handle error: the token is not of type *jwt.Token
		//}
		//fmt.Println(token2)

		//authHeader := c.Request.Header.Get("Authorization")
		//parts := strings.Split(authHeader, " ")
		//if len(parts) != 2 || parts[0] != "Bearer" {
		//	// Handle error: the Authorization header format is not correct
		//}
		//
		//token1 := parts[1]
		//token1 := c.Request.Header.Values("Authorization").(*jwt.Token)
		//fmt.Println(token1)
		//claims := token1.Claims.(*usecases.JwtCustomClaims)
		//user := usersdto.Users{
		//	UserId: claims.UserId,
		//	Name:   req.Name,
		//	Email:  req.Email,
		//}
		//fmt.Println(user)
		//authHeader := c.Request.Header.Get("Authorization")
		//parts := strings.Split(authHeader, " ")
		//if len(parts) != 2 || parts[0] != "Bearer" {
		//	// Handle error: the Authorization header format is not correct
		//}
		//tokenString := parts[1]
		//token1, err := jwt.ParseWithClaims(tokenString, &usecases.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		//	// Validate the token here
		//	return []byte(sercurity.SECRET_KEY_ADMIN), nil
		//})
		//if err != nil {
		//	// Handle error: the token could not be parsed
		//}

		//claims := token1.Claims.(*usecases.JwtCustomClaims)
		//user := usersdto.Users{
		//	UserId: claims.UserId,
		//	Name:   req.Name,
		//	Email:  req.Email,
		//}
		//fmt.Println(user)

		//gen token
		token, err := sercurity.GenTokenUsers(usersdto.Users{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, res.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		userUsers := usersdto.Users{
			UserId:   userUsersId.String(),
			Name:     req.Name,
			PassWord: PassHash,
			Email:    req.Email,
			Role:     role,
			Token:    token,
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
			Data:       data.UserId,
		})
	}
}
