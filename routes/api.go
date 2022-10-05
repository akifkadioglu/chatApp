package routes

import (
	"chatApp/adapter"
	controllers "chatApp/controllers"
)

func Api() {
	api := E.Group("/api")
	admin := api.Group("/admin", adapter.AdminAdapter)
	//for everybody
	api.GET("/getContacts", controllers.GetContacts)
	api.GET("/connectAUser", controllers.ConnectAUser)
	api.POST("/searchUsers", controllers.SearchUser)

	//for admin
	admin.GET("/getUsers", controllers.GetUsers)
}
