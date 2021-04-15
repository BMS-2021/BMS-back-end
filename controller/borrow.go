package controller

import (
	"BMS-back-end/model"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// @tags Borrow
// @summary Borrow a new book
// @router /borrow [post]
// @param data body model.BorrowReq true " "
// @produce plain
// @success 200
// @failure 400 {string} string "Bad request"
// @failure 404 {string} string "Book not found or Card not found"
// @produce json
// @failure 403 {object} object "The requested book has been all borrowed, return the last borrowed time object"
func createBorrow(c echo.Context) error {
	var borrowReq model.BorrowReq
	if err := c.Bind(&borrowReq); err != nil {
		return c.String(http.StatusBadRequest, "there are some errors with the parameters")
	} else if err = c.Validate(&borrowReq); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var dbBook *model.Book
	dbBook, err := model.RetrieveBook(borrowReq.BookId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "Book ID not found in database")
		}
		logrus.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	if _, err := model.GetCard(borrowReq.CardId); err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "Card ID not found in database")
		}
		logrus.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	if dbBook.Stock == 0 {
		t, err := model.GetNearestBorrowTime(borrowReq.BookId)
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusForbidden, &struct {Time time.Time `json:"time"`}{Time: *t})
	}

	if err := model.BorrowBook(borrowReq.BookId, borrowReq.CardId, c.Get("uid").(uint64), dbBook); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

// @tags Borrow
// @summary Return a book
// @router /return [post]
// @param data body model.BorrowReq true " "
// @produce plain
// @success 200
// @failure 404 {string} string "Book not found or Card not found or Borrow not found"
func updateReturn(c echo.Context) error {
	var borrowReq model.BorrowReq
	if err := c.Bind(&borrowReq); err != nil {
		return c.String(http.StatusBadRequest, "there are some errors with the parameters")
	} else if err = c.Validate(&borrowReq); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var dbBook *model.Book
	dbBook, err := model.RetrieveBook(borrowReq.BookId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "Book ID not found in database")
		}
		logrus.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	if _, err := model.GetCard(borrowReq.CardId); err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "Card ID not found in database")
		}
		logrus.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	switch model.CheckBorrowNumByBookIdAndCardId(borrowReq.BookId, borrowReq.CardId) {
	case 0:
		return c.String(http.StatusNotFound, "Borrow not found in database")
	case 1:
		break
	default:
		logrus.WithField("bookId", borrowReq.BookId).
			WithField("cardId", borrowReq.CardId).
			Error("A person has borrowed multiple same books")
		return c.NoContent(http.StatusInternalServerError)
	}

	if err := model.ReturnBook(borrowReq.BookId, borrowReq.CardId, dbBook); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

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
