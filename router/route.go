package router

import (
	"github.com/MinhSang97/order_app/dbutil"
	"github.com/MinhSang97/order_app/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Route() {
	db := dbutil.ConnectDB()
	fmt.Println("Connected: ", db)

	// CRUD: Create, Read, Update, Delete
	// POST /v1/items (create a new item)
	// GET /v1/items (list items) /v1/items?page=1
	// GET /v1/items/:id (get item detail by id)
	// (PUT | PATCH) /v1/items/:id (update an item by id)
	// DELETE /v1/items/:id (delete item by id)
	//viper.SetConfigFile("config.yaml")
	//if err := viper.ReadInConfig(); err != nil {
	//	panic(err)
	//}

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
			// //admin
			// items.POST("/admin/sign-up", handler.AdminSignUp())
			// items.POST("/admin/sign-in", handler.AdminSignIn())
			// items.PATCH("/admin/update/:user_id", middleware.JWTMiddlewareAdmin(), handler.AdminUpdate())
			// items.DELETE("/admin/delete/:user_id", middleware.JWTMiddlewareAdmin(), handler.AdminDelete())

			// //user
			// items.POST("/users/sign-up", middleware.JWTMiddlewareAdmin(), usersHandler.UsersSignUp())
			// items.POST("/users/sign-in", usersHandler.UsersSignIn())
			// items.PATCH("/users/update/:user_id", middleware.JWTMiddlewareUsers(), usersHandler.UsersUpdate())
			// items.DELETE("/users/delete/:user_id", middleware.JWTMiddlewareAdmin(), usersHandler.UsersDelete())

			// //attendance
			// items.POST("/attendance/check-in", middleware.JWTMiddlewareUsers(), attendance.AttendanceCheckIn())
			// items.POST("/attendance/check-out/:id", middleware.JWTMiddlewareUsers(), attendance.AttendanceCheckOut())
			// items.GET("/attendance/history", middleware.JWTMiddlewareUsers(), attendance.AttendanceHistory())
			//items.GET("", handler.GetAllStudent(db))
			//items.GET("/:id", handler.GetId(db))
			//items.PATCH("/:id", handler.Update_One(db))
			//items.DELETE("/:id", handler.Delete_One(db))
			//items.GET("/search", handler.SearchStudents(db))

		}
	}

	r.Run()

}
