package model

import "time"

type Borrow struct {
	BookID     uint64
	Book       Book
	CardID     uint64
	Card       Card
	AdminID    uint64
	Admin      Admin
	BorrowTime time.Time `gorm:"not null"`
	ReturnTime time.Time `gorm:"not null"`
}
