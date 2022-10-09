package messagecontroller

import (
	"chatApp/database"
	"chatApp/helpers"
	"chatApp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

type inputDeleteAllMessage struct {
	ToId int `json:"to_id" validate:"required"`
}

func DeleteAllMessage(c echo.Context) error {
	var input inputDeleteAllMessage
	db := database.DBManager()
	c.Bind(&input)
	err := deleteAllMessageValidate(c, input)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]bool{
			"isDeleted": false,
		})
	}
	result := deleteMessageTransaction(c, db, input)
	if result != nil {
		return result
	}

	return c.JSON(http.StatusOK, map[string]bool{
		"isDeleted": true,
	})
}

func deleteMessageTransaction(c echo.Context, db gorm.DB, input inputDeleteAllMessage) error {
	var message []models.Message

	return db.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("`from_id` = " + strconv.Itoa(int(helpers.User(c).ID))).Where("`to_id` = " + strconv.Itoa(input.ToId)).Delete(&message)
		if result.Error != nil {
			return c.JSON(http.StatusOK, map[string]bool{
				"isDeleted": false,
			})
		}
		result = tx.Where("`to_id` = " + strconv.Itoa(int(helpers.User(c).ID))).Where("`from_id` = " + strconv.Itoa(input.ToId)).Delete(&message)
		if result.Error != nil {
			return c.JSON(http.StatusOK, map[string]bool{
				"isDeleted": false,
			})
		}
		return nil

	})
}
func deleteAllMessageValidate(c echo.Context, input inputDeleteAllMessage) error {
	v := validator.New()
	err := v.Struct(&input)
	if err != nil {
		return err
	}
	return nil
}
