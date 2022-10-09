package messagecontroller

import (
	"chatApp/database"
	"chatApp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type inputDeleteMessage struct {
	MessageId int `json:"message_id" validate:"required"`
}

func DeleteMessage(c echo.Context) error {

	var input inputDeleteMessage
	var message models.Message
	db := database.DBManager()
	
	c.Bind(&input)
	err := deleteMessageValidate(c, input)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]bool{
			"isDeleted": false,
		})
	}

	result := db.Where("`message_id` = " + strconv.Itoa(input.MessageId)).Delete(&message)
	if result.Error != nil {
		return c.JSON(http.StatusOK, map[string]bool{
			"isDeleted": false,
		})
	}

	return c.JSON(http.StatusOK, map[string]bool{
		"isDeleted": true,
	})

}

func deleteMessageValidate(c echo.Context, input inputDeleteMessage) error {
	v := validator.New()
	err := v.Struct(&input)
	if err != nil {
		return err
	}
	return nil
}
