package messagecontroller

import (
	"chatApp/database"
	"chatApp/helpers"
	"chatApp/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type inputSendMessageToGroup struct {
	Message  string `json:"message" validate:"required"`
	Group_Id int    `json:"group_id" validate:"required"`
}

func SendMessageToGroup(c echo.Context) error {
	var message models.GroupMessages
	var input inputSendMessageToGroup
	db := database.DBManager()

	c.Bind(&input)

	err := sendMessageToGroupValidate(c, input)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]bool{
			"isSended": false,
		})
	}
	message.UserId = int(helpers.User(c).ID)
	message.GroupId = input.Group_Id
	message.Message = input.Message
	result := db.Create(&message)

	if result.Error != nil {
		return c.JSON(http.StatusOK, map[string]bool{
			"isSended": false,
		})
	}

	return c.JSON(http.StatusOK, map[string]bool{
		"isSended": true,
	})
}

func sendMessageToGroupValidate(c echo.Context, input inputSendMessageToGroup) error {
	v := validator.New()
	err := v.Struct(input)
	if err != nil {
		return err
	}
	return nil
}
