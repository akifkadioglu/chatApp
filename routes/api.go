package routes

import (
	"chatApp/adapter"
	controllers "chatApp/controllers"
	authcontroller "chatApp/controllers/AuthController"
	contactcontroller "chatApp/controllers/ContactController"
	usercontroller "chatApp/controllers/UserController"
)

func Api() {
	api := Network.Group("/api")

	//Authenticate
	api.POST("/register", authcontroller.Register)
	api.POST("/login", authcontroller.Login)
	api.POST("/forgot-password", authcontroller.ForgotPassword)

	/*                            For auth users                            */
	auth := api.Group("")
	auth.Use(adapter.AuthAdapter())

	//contact controllers
	auth.GET("/getContacts", contactcontroller.GetContacts)
	auth.POST("/addAContact", contactcontroller.AddAContact)

	//user controller
	auth.POST("/searchUsers", usercontroller.SearchUser)

	//index
	auth.GET("/index", controllers.Index)
}
