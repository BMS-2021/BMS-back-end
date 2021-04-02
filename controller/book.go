package controller

import (
	"BMS-back-end/model"
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

// @tags Book
// @summary Create multiple book items by uploading a csv file
// @router /books [put]
// @accept text/csv
// @param file body object true "A csv file, with book information inside"
// @success 200
func storeBookCsv(c echo.Context) error {
	fileReq, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "cannot find the file key in request body")
	}

	file, err := fileReq.Open()
	if err != nil {
		logrus.Error("error opening file in storeBookCsv API")
		return c.NoContent(http.StatusInternalServerError)
	}
	defer file.Close()

	fileReadBuffer := make([]byte, 512)
	_, err = file.Read(fileReadBuffer)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	fileMimeType := http.DetectContentType(fileReadBuffer)
	if !strings.HasPrefix(fileMimeType, "text") && fileMimeType != "application/octet-stream"  {
		return c.String(http.StatusUnsupportedMediaType, "the uploaded file formant is not text/csv")
	}

	cr := csv.NewReader(file)
	record, err := cr.Read()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	columnNames := map[string]bool{"category": false, "title": false, "press": false,
		"year": false, "author": false, "price": false, "total": false, "stock": false}
	for _, v := range record {
		if _, ok := columnNames[v]; ok {
			columnNames[v] = true
		}
	}
	for _, v := range columnNames {
		if !v {
			return c.String(http.StatusBadRequest, "some required columns in csv file does not exist")
		}
	}

	fileReadBuffer = make([]byte, fileReq.Size)
	_, err = file.Read(fileReadBuffer)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	var books []*model.Book
	if err := gocsv.UnmarshalBytes(fileReadBuffer, &books); err != nil {
		logrus.Error("error when unmarshalling csv")
		logrus.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	err = model.CreateBooks(&books)
	if err != nil {
		logrus.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

// @tags Book
// @summary Create a single book item
// @router /book [put]
// @accept json
// @param book body model.BookReq true "Information of a book"
// @success 200
func storeBook(c echo.Context) error {
	bookReq := model.BookReq{}
	if err := c.Bind(&bookReq); err != nil {
		return c.String(http.StatusBadRequest, "there are some errors with the parameters")
	} else if err = c.Validate(&bookReq); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	book := model.Book{}
	_ = copier.Copy(&book, &bookReq)

	err := model.CreateBooks(&[]*model.Book{&book})
	if err != nil {
		logrus.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

// @tags Book
// @summary Retrieve multiple books under multiple conditions
// @router /books [get]
// @param data query model.BookQueryReq false " "
// @produce json
// @success 200
func retrieveBooks(c echo.Context) error {
	bookQueryReq := model.BookQueryReq{}
	if err := c.Bind(&bookQueryReq); err != nil {
		return c.String(http.StatusBadRequest, "there are some errors with the parameters")
	}

	book := model.Book{}
	_ = copier.Copy(&book, &bookQueryReq)

	dbBooks, err := model.RetrieveBooks(&book,
		[]uint64{bookQueryReq.YearMin, bookQueryReq.YearMax},
		[]uint64{bookQueryReq.PriceMin, bookQueryReq.PriceMax})

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, dbBooks)
}
