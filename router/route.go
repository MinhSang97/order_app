package router

import (
	"fmt"
	"github.com/MinhSang97/order_app/dbutil"
	"github.com/MinhSang97/order_app/handler"
	admin_function_feedback "github.com/MinhSang97/order_app/handler/admin_function/feedback"
	admin_function_member "github.com/MinhSang97/order_app/handler/admin_function/member"
	admin_function_menu "github.com/MinhSang97/order_app/handler/admin_function/menu"
	admin_function_order "github.com/MinhSang97/order_app/handler/admin_function/order"
	admin_function_promotion "github.com/MinhSang97/order_app/handler/admin_function/promotion"
	admin_login "github.com/MinhSang97/order_app/handler/admin_login"
	users_function "github.com/MinhSang97/order_app/handler/users_function"
	middleware "github.com/MinhSang97/order_app/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/MinhSang97/order_app/docs"
	"github.com/gin-gonic/gin"
)

// Route routing mapping
// @title github.com/MinhSang97/order_app API
// @version 1.0
// @description This is a sample server github.com/MinhSang97/order_app server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1/api/
// @schemes http
func Route() {
	db := dbutil.ConnectDB()
	fmt.Println("Connected: ", db)

	r := gin.Default()
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.SaveLogRequest())

	// Swagger endpoint
	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	{
		api := v1.Group("/api")
		{
			api.GET("/test", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "This is a secure route"})
			})
			//OTP
			api.POST("/send_otp", handler.SendOTP())
			api.POST("/verify_otp/:otp", handler.VerifiOTP())
			api.PATCH("/change_password/:otp", handler.ChangePassWord())

			//admin
			api.POST("/admin/sign_up", admin_login.AdminSignUp())
			api.POST("/admin/sign_in", admin_login.AdminSignIn())
			api.PATCH("/admin/update/:user_id", middleware.JWTMiddlewareAdmin(), admin_login.AdminUpdate())
			api.DELETE("/admin/delete/:user_id", middleware.JWTMiddlewareAdmin(), admin_login.AdminDelete())

			//admin_function_member
			api.GET("/admin/member_view", middleware.JWTMiddlewareAdmin(), admin_function_member.AdminMemberView())
			api.PATCH("/admin/member_edit/:user_id", middleware.JWTMiddlewareAdmin(), admin_function_member.AdminMemberEdit())
			api.POST("/admin/member_add", middleware.JWTMiddlewareAdmin(), admin_function_member.AdminMemberAdd())
			api.DELETE("/admin/member_delete/:email", middleware.JWTMiddlewareAdmin(), admin_function_member.AdminMemberDelete())

			//admin_function_menu
			api.GET("/admin/menu_view", middleware.JWTMiddlewareAdmin(), admin_function_menu.AdminMenuView())
			api.PATCH("/admin/menu_edit/:item_id", middleware.JWTMiddlewareAdmin(), admin_function_menu.AdminMenuEdit())
			api.POST("/admin/menu_add", middleware.JWTMiddlewareAdmin(), admin_function_menu.AdminMenuAdd())
			api.DELETE("/admin/menu_delete/:item_id", middleware.JWTMiddlewareAdmin(), admin_function_menu.AdminMenuDelete())

			//admin_function_promotion
			api.GET("/admin/promotion_view", middleware.JWTMiddlewareAdmin(), admin_function_promotion.AdminPromotionView())
			api.PATCH("/admin/promotion_edit/:discount_code_id", middleware.JWTMiddlewareAdmin(), admin_function_promotion.AdminDiscountCodeEdit())
			api.POST("/admin/promotion_add", middleware.JWTMiddlewareAdmin(), admin_function_promotion.AdminPromotionAdd())
			api.DELETE("/admin/promotion_delete/:discount_code_id", middleware.JWTMiddlewareAdmin(), admin_function_promotion.AdminPromotionDelete())

			//admin_function_feedback
			api.GET("/admin/feedback_view", middleware.JWTMiddlewareAdmin(), admin_function_feedback.AdminFeedbackView())

			//admin_hístory_order
			api.GET("/admin/order_history", middleware.JWTMiddlewareAdmin(), admin_function_order.AdminOrderHistory())

			//user
			api.POST("/users/sign-up", users_function.UsersSignUp())
			api.POST("/users/sign_in", users_function.UsersSignIn())
			api.PATCH("/users/update/:user_id", middleware.JWTMiddlewareUsers(), users_function.UsersUpdate())
			api.DELETE("/users/delete/:user_id", middleware.JWTMiddlewareUsers(), users_function.UsersDelete())

			//user_function
			api.GET("/users/get_address/:user_id", middleware.JWTMiddlewareUsers(), users_function.UsersGetAddress())
			api.POST("/users/add_address/:user_id", middleware.JWTMiddlewareUsers(), users_function.UsersAddAddress())
			api.PATCH("/users/change_address_default/:user_id", middleware.JWTMiddlewareUsers(), users_function.UsersChangeAddressDefault())
			api.GET("/users/get_menu", middleware.JWTMiddlewareUsers(), users_function.UsersGetMenu())
			api.GET("/users/get_promotion", middleware.JWTMiddlewareUsers(), users_function.UsersGetDiscountCodes())

			//user_order
			api.POST("/users/order/:user_id", middleware.JWTMiddlewareUsers(), users_function.UsersOrder())
			api.PATCH("/users/order_status/:user_id", middleware.JWTMiddlewareUsers(), users_function.UsersOrderStatus())
			api.GET("/users/order_history/:user_id", middleware.JWTMiddlewareUsers(), users_function.UsersOrderHistory())

			//user_feedback
			api.POST("/users/feedback/:user_id", middleware.JWTMiddlewareUsers(), users_function.UsersFeedback())
			api.GET("/users/feedback_history/:user_id", middleware.JWTMiddlewareUsers(), users_function.UsersFeedbackHistory())
		}
	}

	r.Run()
}
