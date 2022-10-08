package routes

import (
	"chatApp/adapter"
	controllers "chatApp/controllers"
	authcontroller "chatApp/controllers/AuthController"
	contactcontroller "chatApp/controllers/ContactController"
	groupcontroller "chatApp/controllers/GroupController"
	usercontroller "chatApp/controllers/UserController"
)

func Api() {
	api := Network.Group("/api")
	/*                            For all users                            */
	//Authenticate
	api.POST("/register", authcontroller.Register)
	api.POST("/login", authcontroller.Login)
	api.POST("/forgot-password", authcontroller.ForgotPassword)
	/*                            For all users                            */

	/*                            For auth users                            */
	auth := api.Group("")
	auth.Use(adapter.AuthAdapter())

	//contact controllers
	auth.GET("/getContacts", contactcontroller.GetContacts)
	auth.POST("/addAContact", contactcontroller.AddAContact)

	//user controllers
	auth.POST("/searchUsers", usercontroller.SearchUser)
	auth.POST("/getSingleUser", usercontroller.GetSingleUser)

	//group controllers
	auth.POST("/createGroup", groupcontroller.CreateGroup)
	auth.POST("/addGroupUser", groupcontroller.AddGroupUser)
	auth.GET("/getUserGroups", groupcontroller.GetUserGroups)

	//index
	auth.GET("/index", controllers.Index)
	/*                            For auth users                            */

}
