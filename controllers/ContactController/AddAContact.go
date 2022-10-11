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

type inputAddAContact struct {
	To int `json:"to_id"`
}

func AddAContact(c echo.Context) error {
	var contact models.Contact
	var input inputAddAContact
	var block models.Block
	db := database.DBManager()

	c.Bind(&input)
	err := addAContactValidate(c, input)
	if err != nil {
		return echo.ErrBadRequest
	}
	result := db.Where("to_id =" + strconv.Itoa(int(input.To))).Where("from_id =" + strconv.Itoa(int(helpers.User(c).ID))).Find(&block)

	if result.RowsAffected > 0 {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	contact.UserId = int(helpers.User(c).ID)
	contact.ToId = input.To
	result = db.Create(&contact)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	return c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
}

func addAContactValidate(c echo.Context, input inputAddAContact) error {
	v := validator.New()
	if err := v.Struct(input); err != nil {
		return err
	}
	return nil
}
