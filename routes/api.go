package routes

import (
	"chatApp/adapter"
	controllers "chatApp/controllers"
	authcontroller "chatApp/controllers/AuthController"
	contactcontroller "chatApp/controllers/ContactController"
	groupcontroller "chatApp/controllers/GroupController"
	messagecontroller "chatApp/controllers/MessageController"
	profilecontroller "chatApp/controllers/ProfileController"
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
	//check auth
	auth.GET("/getAuthUser", authcontroller.GetAuthUser)

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
	auth.GET("/getGroupUsers", groupcontroller.GetGroupUsers)

	//message controllers
	auth.GET("/getMessages", messagecontroller.GetMessages)
	auth.GET("/getContactsMessages", messagecontroller.GetContactsMessages)
	auth.POST("/sendMessage", messagecontroller.SendMessage)
	auth.DELETE("/deleteMessage", messagecontroller.DeleteMessage)
	auth.DELETE("/deleteAllMessage", messagecontroller.DeleteAllMessage)

	//profile controllers
	auth.PUT("/updateProfile", profilecontroller.UpdateProfile)
	auth.PUT("/changePassword", profilecontroller.ChangePassword)

	//index
	auth.GET("/index", controllers.Index)
	/*                            For auth users                            */

}
