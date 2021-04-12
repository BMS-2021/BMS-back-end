package controller

import (
	"BMS-back-end/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @tags Borrow
// @summary Get the books borrowed by a specific Bard
// @router /borrow [get]
// @param cardId query uint64 true "Card ID"
// @produce json
// @success 200 {object} []model.Book
// @failure 400 {string} string "Bad request"
func getBorrowed(c echo.Context) error {
	var cardId uint64
	if err := echo.QueryParamsBinder(c).MustUint64("cardId", &cardId).BindError(); err != nil {
		return c.String(http.StatusBadRequest, "there are some errors with the parameters")
	}

	borrows, err := model.GetBorrowWithBook(cardId)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	booksMap := make(map[uint64]*model.Book)
	books := make([]*model.Book, 0)
	for _, v := range *borrows {
		if _, ok := booksMap[v.BookID]; !ok {
			booksMap[v.BookID] = &v.Book
			books = append(books, &v.Book)
		}
	}

	return c.JSON(http.StatusOK, &books)
}
