package handler

import (
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

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

		//PassHash := sercurity.HashAndSalt([]byte(req.PassWord))
		//role := payload.ADMIN.String()
		//
		//userAdminId, err := uuid.NewUUID()
		//
		//if err != nil {
		//	log.Error(err.Error())
		//	c.JSON(http.StatusForbidden, res.Response{
		//		StatusCode: http.StatusForbidden,
		//		Message:    err.Error(),
		//		Data:       nil,
		//	})
		//	return
		//}

		//userAdmin := admindto.Admin{
		//	UserId:   userAdminId.String(),
		//	Name:     req.Name,
		//	PassWord: PassHash,
		//	Email:    req.Email,
		//	Role:     role,
		//	//Token:       token,
		//	PhoneNumber: req.PhoneNumber,
		//	Address:     req.Address,
		//}
		//
		//err = validate.Struct(userAdmin)
		//
		//if err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"error": err.Error(),
		//	})
		//	return
		//}
		//
		//data := userAdmin.ToPayload().ToModel()
		//uc := usecases.NewAdminUseCase()
		//
		//err = uc.CreateAdmin(c.Request.Context(), data)
		//
		//if err != nil {
		//	c.JSON(http.StatusConflict, res.Response{
		//		StatusCode: http.StatusConflict,
		//		Message:    err.Error(),
		//		Data:       nil,
		//	})
		//	return
		//}

		c.JSON(http.StatusOK, res.Response{
			StatusCode: http.StatusOK,
			Message:    "Xử lý thành công",
			Data:       "null",
		})
	}
}
