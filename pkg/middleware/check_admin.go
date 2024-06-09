package middleware

import (
	"github.com/MinhSang97/order_app/usecases/req"
	"github.com/MinhSang97/order_app/usecases/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ISAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Perform logic to check if the user is an admin
		req := req.ReqSignIn{}
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, res.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			})
			c.Abort() // Abort the request chain
			return
		}
		if req.Email != "" {
			c.JSON(http.StatusBadRequest, res.Response{
				StatusCode: http.StatusBadRequest,
				Message:    "Bạn không quyền gọi API này",
				Data:       nil,
			})
			c.Abort() // Abort the request chain
			return
		}
		// If the user is an admin, continue with the request chain
		c.Next()
	}
}
