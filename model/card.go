package model

type Card struct {
	ID         uint64
	Name       string `gorm:"not null;size:10"`
	Department string `gorm:"not null;size:40"`
	Type       string `gorm:"not null;size:1"`
}
