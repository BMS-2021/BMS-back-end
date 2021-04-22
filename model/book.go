package model

import (
	"gorm.io/gorm"
	"strconv"
)

type Book struct {
	ID       uint64  `gorm:"not null;autoIncrement;primaryKey" csv:"-" json:"id"`
	Category string  `gorm:"not null;size:10" csv:"category" json:"category"`
	Title    string  `gorm:"not null;size:40" csv:"title" json:"title"`
	Press    string  `gorm:"not null;size:30" csv:"press" json:"press"`
	Year     uint64  `gorm:"not null" csv:"year" json:"year"`
	Author   string  `gorm:"not null;size:20" csv:"author" json:"author"`
	Price    float64 `gorm:"not null" csv:"price" json:"price"`
	Total    uint64  `gorm:"not null" csv:"total" json:"total"`
	Stock    uint64  `gorm:"not null" csv:"stock" json:"stock"`
}

type BookReq struct {
	Category string  `json:"category" validate:"required"`
	Title    string  `json:"title" validate:"required"`
	Press    string  `json:"press" validate:"required"`
	Year     uint64  `json:"year" validate:"required"`
	Author   string  `json:"author" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	Total    uint64  `json:"total" validate:"required"`
}

type BookQueryReq struct {
	Category string `query:"category"`
	Title    string `query:"title"`
	Press    string `query:"press"`
	Author   string `query:"author"`
	YearMin  uint64 `query:"yearMin"`
	YearMax  uint64 `query:"yearMax"`
	PriceMin uint64 `query:"priceMin"`
	PriceMax uint64 `query:"priceMax"`
	Order    string `query:"order"`
	Desc     bool   `query:"desc"`
}

func CreateBooks(books *[]*Book) error {
	result := db.Create(books)
	return result.Error
}

func RetrieveBooks(book *Book, year []uint64, price []uint64, orderBy string, isDesc bool) (*[]Book, error) {
	sqlConditionCount := 0
	sqlConditionString := ""
	if year[0] != 0 {
		sqlConditionCount++
		sqlConditionString += "year >= " + strconv.FormatUint(year[0], 10)
	}
	if year[1] != 0 {
		if sqlConditionCount != 0 {
			sqlConditionString += " AND "
		}
		sqlConditionCount++
		sqlConditionString += "year <= " + strconv.FormatUint(year[1], 10)
	}
	if price[0] != 0 {
		if sqlConditionCount != 0 {
			sqlConditionString += " AND "
		}
		sqlConditionCount++
		sqlConditionString += "price >= " + strconv.FormatUint(price[0], 10)
	}
	if price[1] != 0 {
		if sqlConditionCount != 0 {
			sqlConditionString += " AND "
		}
		sqlConditionCount++
		sqlConditionString += "price <= " + strconv.FormatUint(price[1], 10)
	}

	order := ""
	if orderBy != "" {
		order += orderBy
		if isDesc {
			order += " DESC"
		}
	}

	var dbBooks []Book
	var result *gorm.DB
	if order == "" {
		result = db.Where(book).Find(&dbBooks, sqlConditionString)
	} else {
		result = db.Where(book).Order(order).Find(&dbBooks, sqlConditionString)
	}

	return &dbBooks, result.Error
}

func RetrieveBook(id uint64) (*Book, error) {
	dbBook := Book{ID: id}
	result := db.First(&dbBook)
	return &dbBook, result.Error
}
