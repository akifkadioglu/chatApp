package contactcontroller

import (
	"chatApp/database"
	"chatApp/helpers"
	"chatApp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type inputDeleteAContact struct {
	To int `json:"to_id" validate:"required"`
}

// the user's all contacts
func DeleteAContact(c echo.Context) error {
	var contact []models.Contact
	var input inputDeleteAContact

	c.Bind(&input)
	deleteAContactValidate(c, input)

	db := database.DBManager()
	result := db.Where("to_id =" + strconv.Itoa(int(input.To))).Where("user_id =" + strconv.Itoa(int(helpers.User(c).ID))).Delete(&contact)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	return c.JSON(http.StatusOK, map[string]bool{
		"isDeleted": true,
	})
}
func deleteAContactValidate(c echo.Context, input inputDeleteAContact) error {
	v := validator.New()
	if err := v.Struct(input); err != nil {
		return err
	}
	return nil
}
