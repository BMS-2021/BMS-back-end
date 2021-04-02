package model

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

func CreateBooks(books *[]*Book) error {
	result := db.Create(books)
	return result.Error
}
