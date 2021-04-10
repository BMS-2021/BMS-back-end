package model

import "time"

type Borrow struct {
	BookID     uint64 `gorm:"not null"`
	Book       Book
	CardID     uint64 `gorm:"not null"`
	Card       Card
	AdminID    uint64 `gorm:"not null"`
	Admin      Admin
	BorrowTime time.Time `gorm:"not null"`
	ReturnTime time.Time
}

func GetBorrowWithBook(cardId uint64) (*[]Borrow, error) {
	var borrows []Borrow
	result := db.Preload("Book").Where(&Borrow{CardID: cardId}).Find(&borrows)
	return &borrows, result.Error
}
