package member

import (
	"github.com/MinhSang97/order_app/log"
	"github.com/MinhSang97/order_app/payload"
	"github.com/MinhSang97/order_app/sercurity"
	"github.com/MinhSang97/order_app/usecases"
	admindto "github.com/MinhSang97/order_app/usecases/dto/admin_dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
)

func AdminMemberAdd() func(*gin.Context) {
	return func(c *gin.Context) {
		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqAdminFunctionAdd{}
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
		var reqRole string
		if req.Role == "admin" {
			reqRole = payload.ADMIN.String()
		} else if req.Role == "users" {
			reqRole = payload.USERS.String()
		} else {
			c.JSON(http.StatusForbidden, res.Response{
				StatusCode: http.StatusForbidden,
				Message:    "Sai role",
				Data:       req.Role,
			})
			return
		}
		role := reqRole

		userId, err := uuid.NewUUID()
		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusForbidden, res.Response{
				StatusCode: http.StatusForbidden,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		user := admindto.AdminFunctionDto{
			UserId:      userId.String(),
			Name:        req.Name,
			PassWord:    PassHash,
			Email:       req.Email,
			Role:        role,
			PhoneNumber: req.PhoneNumber,
			Address:     req.Address,
		}

		err = validate.Struct(user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data := user.ToPayload().ToModel()
		uc := usecases.NewAdminFunctionUseCase()
		//uc := usecases.NewAdminUseCase()
		//
		//err = uc.CreateAdmin(c.Request.Context(), data)
		//
		err = uc.AddUser(c.Request.Context(), data)
		if err != nil {
			c.JSON(http.StatusConflict, res.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}
		usersAdd := admindto.AdminFunctionDto{
			UserId:   data.UserId,
			Name:     req.Name,
			Email:    req.Email,
			PassWord: req.PassWord,
			Role:     req.Role,
		}

		c.JSON(http.StatusOK, res.Response{
			StatusCode: http.StatusOK,
			Message:    "Xử lý thành công",
			Data:       usersAdd,
		})
	}
}
