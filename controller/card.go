package controller

import (
	"BMS-back-end/model"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

// @tags Card
// @summary Create a library card
// @router /card [put]
// @param data body model.CardReq true " "
// @produce json
// @failure 400 {string} string "Bad request"
// @failure 500
func createCard(c echo.Context) error {
	cardReq := model.CardReq{}
	if err := c.Bind(&cardReq); err != nil {
		return c.String(http.StatusBadRequest, "there are some errors with the parameters")
	} else if err = c.Validate(&cardReq); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	card := model.Card{}
	_ = copier.Copy(&card, &cardReq)

	err := model.CreateCard(&card)
	if err != nil {
		logrus.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, &struct{ID uint64 `json:"id"`}{ID: card.ID})

}

// @tags Card
// @summary Retrieve a library card
// @router /card [get]
// @param id query uint true "Card ID"
// @produce json
// @failure 400 {string} string "Bad request"
// @failure 500
func getCard(c echo.Context) error {
	var id uint64
	if err := echo.QueryParamsBinder(c).MustUint64("id", &id).BindError(); err != nil {
		return c.String(http.StatusBadRequest, "there are some errors with the parameters")
	}

	dbCard, err := model.GetCard(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "Card ID not found in database")
		}
		logrus.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, dbCard)
}

// @tags Card
// @summary Delete a library card
// @router /card [delete]
// @param id query uint true "Card ID"
// @produce plain
// @success 200
// @failure 400 {string} string "Bad request"
// @failure 500
func deleteCard(c echo.Context) error {
	var id uint64
	if err := echo.QueryParamsBinder(c).MustUint64("id", &id).BindError(); err != nil {
		return c.String(http.StatusBadRequest, "there are some errors with the parameters")
	}

	err := model.DeleteCard(id)
	if err != nil {
		logrus.Error(err)
		return c.String(http.StatusBadRequest, "cannot delete due to foreign key constraint")  // maybe contains error about foreign key
	}

	return c.NoContent(http.StatusOK)
}
