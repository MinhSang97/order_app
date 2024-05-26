package handler

import (
	"github.com/MinhSang97/order_app/usecases"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminDelete() func(*gin.Context) {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")

		if user_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "ko tim thay user_id",
			})
			return
		}

		//var validate *validator.Validate
		//validate = validator.New(validator.WithRequiredStructEnabled())
		//req := req.ReqUpdateUser{}
		//if err := c.ShouldBind(&req); err != nil {
		//	c.JSON(http.StatusBadRequest, res.Response{
		//		StatusCode: http.StatusBadRequest,
		//		Message:    err.Error(),
		//		Data:       nil,
		//	})
		//	return
		//}
		//
		//if err := validate.Struct(req); err != nil {
		//	c.JSON(http.StatusForbidden, res.Response{
		//		StatusCode: http.StatusForbidden,
		//		Message:    err.Error(),
		//		Data:       nil,
		//	})
		//	return
		//}
		//
		//PassHashNew := sercurity.HashAndSalt([]byte(req.PassWord))
		//userAdmin := admindto.Admin{
		//	Name:     req.Name,
		//	Email:    req.Email,
		//	PassWord: PassHashNew,
		//}
		//
		//err := validate.Struct(userAdmin)
		//
		//if err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"error": err.Error(),
		//	})
		//	return
		//}

		//data := userAdmin.ToPayload().ToModel()
		uc := usecases.NewAdminUseCase()

		err := uc.DeleteAdmin(c.Request.Context(), user_id)

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
			Message:    "Xoá thành công",
			Data:       map[string]interface{}{"user_id": user_id},
		})
	}
}
