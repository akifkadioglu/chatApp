package messagecontroller

import (
	"chatApp/database"
	"chatApp/helpers"
	"chatApp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type inputSendMessage struct {
	Message string `json:"message" validate:"required"`
	To_id   int    `json:"to_id" validate:"required"`
}

func SendMessage(c echo.Context) error {
	var message models.Message
	var input inputSendMessage
	var block models.Block
	db := database.DBManager()

	c.Bind(&input)
	err := sendMessageValidate(c, input)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]bool{
			"isSended": false,
		})
	}

	result := db.Where("`to_id` = " + strconv.Itoa(input.To_id)).Where("`from_id` = " + strconv.Itoa(int(helpers.User(c).ID))).Find(&block)
	if result.RowsAffected > 0 {
		return c.JSON(http.StatusOK, map[string]bool{
			"isSended": false,
		})
	}

	message.FromId = int(helpers.User(c).ID)
	message.ToId = input.To_id
	message.Message = input.Message
	result = db.Create(&message)
	if result.Error != nil {
		return c.JSON(http.StatusOK, map[string]bool{
			"isSended": false,
		})
	}

	return c.JSON(http.StatusOK, map[string]bool{
		"isSended": true,
	})
}

func sendMessageValidate(c echo.Context, input inputSendMessage) error {
	v := validator.New()
	err := v.Struct(input)
	if err != nil {
		return err
	}
	return nil
}
