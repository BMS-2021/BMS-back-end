package model

type Book struct {
	ID       uint64  `gorm:"not null;autoIncrement;primaryKey"`
	Category string  `gorm:"not null;size:10"`
	Title    string  `gorm:"not null;size:40"`
	Press    string  `gorm:"not null;size:30"`
	Year     uint64  `gorm:"not null"`
	Author   string  `gorm:"not null;size:20"`
	Price    float64 `gorm:"not null"`
	Total    uint64  `gorm:"not null"`
	Stock    uint64  `gorm:"not null"`
}
