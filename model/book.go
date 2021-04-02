package model

import "strconv"

type Book struct {
	ID       uint64  `gorm:"not null;autoIncrement;primaryKey" csv:"-"`
	Category string  `gorm:"not null;size:10" csv:"category"`
	Title    string  `gorm:"not null;size:40" csv:"title"`
	Press    string  `gorm:"not null;size:30" csv:"press"`
	Year     uint64  `gorm:"not null" csv:"year"`
	Author   string  `gorm:"not null;size:20" csv:"author"`
	Price    float64 `gorm:"not null" csv:"price"`
	Total    uint64  `gorm:"not null" csv:"total"`
	Stock    uint64  `gorm:"not null" csv:"stock"`
}

type BookReq struct {
	Category string  `json:"category" validate:"required"`
	Title    string  `json:"title" validate:"required"`
	Press    string  `json:"press" validate:"required"`
	Year     uint64  `json:"year" validate:"required"`
	Author   string  `json:"author" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	Total    uint64  `json:"total" validate:"required"`
	Stock    uint64  `json:"stock" validate:"required"`
}

type BookQueryReq struct {
	Category string `query:"category"`
	Title    string `query:"title"`
	Press    string `query:"press"`
	Author   string `query:"author"`
	YearMin  uint64 `query:"year_min"`
	YearMax  uint64 `query:"year_max"`
	PriceMin uint64 `query:"price_min"`
	PriceMax uint64 `query:"price_max"`
}

func CreateBooks(books *[]*Book) error {
	result := db.Create(books)
	return result.Error
}

func RetrieveBooks(book *Book, year []uint64, price []uint64) (*[]Book, error) {
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

	var dbBooks []Book
	result := db.Where(book).Find(&dbBooks, sqlConditionString)
	return &dbBooks, result.Error
}
