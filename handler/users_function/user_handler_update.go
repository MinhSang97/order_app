package handler

import (
	"github.com/MinhSang97/order_app/usecases"
	usersdto "github.com/MinhSang97/order_app/usecases/dto/users_dto"
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UsersUpdateResponse struct {
	Name string `json:"name"`
	//Email       string `json:"email"`
	//PhoneNumber string `json:"phone_number"`
	Sex       string `json:"sex"`
	BirthDate string `json:"birth_date"`
	Telegram  string `json:"telegram"`
}

// UsersUpdate godoc
// @Summary Users can update
// @Description Users can update
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param name body string true "Name"
// @Param sex body string true "Sex"
// @Param birth_date body string true "BirthDate"
// @Param telegram body string true "Telegram"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /v1/api/users/update/{user_id} [patch]
func UsersUpdate() func(*gin.Context) {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		if user_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "không tìm thấy user_id",
			})
			return
		}

		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqUpdateUser{}
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

		//PassHashNew := sercurity.HashAndSalt([]byte(req.PassWord))
		userUsers := usersdto.Users{
			Name: req.Name,
			//Email:       req.Email,
			//PhoneNumber: req.PhoneNumber,
			Sex:       req.Sex,
			BirthDate: req.BirthDate,
			Telegram:  req.Telegram,
		}

		err := validate.Struct(userUsers)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data := userUsers.ToPayload().ToModel()
		uc := usecases.NewUsersUseCase()

		err = uc.UpdateUsers(c.Request.Context(), user_id, data)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, res.Response{
				StatusCode: http.StatusUnprocessableEntity,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		usersUpdate := usersdto.Users{
			Name: req.Name,
			//Email:       req.Email,
			//PhoneNumber: req.PhoneNumber,
			Sex:       req.Sex,
			BirthDate: req.BirthDate,
			Telegram:  req.Telegram,
		}

		c.JSON(http.StatusOK, res.Response{
			StatusCode: http.StatusOK,
			Message:    "Xử lý cập nhật thành công",
			Data:       usersUpdate,
		})
	}
}
