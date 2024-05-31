package router

import (
	"fmt"
	"github.com/MinhSang97/order_app/dbutil"
	"github.com/MinhSang97/order_app/handler"
	admin_function_member "github.com/MinhSang97/order_app/handler/admin_function/member"
	admin_login "github.com/MinhSang97/order_app/handler/admin_login"
	users_login "github.com/MinhSang97/order_app/handler/users_login"
	"github.com/MinhSang97/order_app/middleware"

	"github.com/gin-gonic/gin"
)

func Route() {
	db := dbutil.ConnectDB()
	fmt.Println("Connected: ", db)

	r := gin.Default()
	r.Use(middleware.ErrorHandler())
	// Sử dụng middleware SaveLogRequest()
	r.Use(middleware.SaveLogRequest())
	//r.Use(middleware.BasicAuthMiddleware())

	v1 := r.Group("/v1")
	{
		api := v1.Group("/api")
		{
			api.GET("/test", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "This is a secure route"})
			})

			api.POST("/send_otp", handler.SendOTP())
			api.GET("/verify_otp/:otp", handler.VerifiOTP())

			//admin
			api.POST("/admin/sign_up", admin_login.AdminSignUp())
			api.GET("/admin/sign_in", admin_login.AdminSignIn())
			api.PATCH("/admin/update/:user_id", middleware.JWTMiddlewareAdmin(), admin_login.AdminUpdate())
			api.PATCH("/admin/forget_password/:user_id", middleware.JWTMiddlewareAdmin(), admin_login.AdminForgetPassword())
			api.DELETE("/admin/delete/:user_id", middleware.JWTMiddlewareAdmin(), admin_login.AdminDelete())

			//admin_function_member
			api.GET("/admin/member_view", middleware.JWTMiddlewareAdmin(), admin_function_member.AdminMemberView())
			api.PATCH("/admin/member_edit/:user_id", middleware.JWTMiddlewareAdmin(), admin_function_member.AdminMemberEdit())

			//user
			api.POST("/users/sign_up", users_login.UsersSignUp())
			api.GET("/users/sign_in", users_login.UsersSignIn())
			api.PATCH("/users/update/:user_id", middleware.JWTMiddlewareUsers(), users_login.UsersUpdate())
			api.DELETE("/users/delete/:user_id", middleware.JWTMiddlewareAdmin(), users_login.UsersDelete())

		}
	}

	r.Run()

}
