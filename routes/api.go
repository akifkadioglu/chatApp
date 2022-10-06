package routes

import (
	"chatApp/adapter"
	controllers "chatApp/controllers"
	"chatApp/models"

	"github.com/labstack/echo/v4/middleware"
)

func Api() {
	api := E.Group("/api")

	api.POST("/login", controllers.Login)

	auth := E.Group("/restricted")
	
	admin := api.Group("/admin", adapter.AdminAdapter)
	//for everybody
	config := middleware.JWTConfig{
		KeyFunc: models.GetKey,
	}
	auth.Use(middleware.JWTWithConfig(config))
	auth.GET("/getContacts", controllers.GetContacts)
	auth.GET("/connectAUser", controllers.ConnectAUser)
	auth.POST("/searchUsers", controllers.SearchUser)

	//for admin
	admin.GET("/getUsers", controllers.GetUsers)
}
