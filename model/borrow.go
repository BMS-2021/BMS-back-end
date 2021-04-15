package model

import (
	"gorm.io/gorm"
	"time"
)

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

type BorrowReq struct {
	BookId uint64 `json:"bookId" validate:"required"`
	CardId uint64 `json:"cardId" validate:"required"`
}

func GetBorrowWithBook(cardId uint64) (*[]Borrow, error) {
	var borrows []Borrow
	result := db.Preload("Book").Where(&Borrow{CardID: cardId}).Find(&borrows)
	return &borrows, result.Error
}

func BorrowBook(bookId uint64, cardId uint64, adminId uint64, dbBook *Book) error {
	return db.Transaction(func(tx *gorm.DB) error {
		dbBook.Stock -= 1
		if err := tx.Save(dbBook).Error; err != nil {
			return err
		}

		if err := tx.Create(&Borrow{BookID: bookId, CardID: cardId, AdminID: adminId, BorrowTime: time.Now()}).
			Error; err != nil {
			return err
		}

		return nil
	})
}

func ReturnBook(bookId uint64, cardId uint64, dbBook *Book) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Borrow{}).Where(&Borrow{BookID: bookId, CardID: cardId}).
			Update("return_time", time.Now()).Error; err != nil {
			return err
		}

		dbBook.Stock += 1
		if err := tx.Save(dbBook).Error; err != nil {
			return err
		}

		return nil
	})
}

func GetNearestBorrowTime(bookId uint64) (*time.Time, error) {
	dbBorrow := Borrow{BookID: bookId}
	result := db.Order("borrow_time ASC").First(&dbBorrow)
	return &dbBorrow.BorrowTime, result.Error
}

func CheckBorrowNumByBookIdAndCardId(bookId uint64, cardId uint64) uint64 {
	result := db.Find(&Borrow{BookID: bookId, CardID: cardId})
	return uint64(result.RowsAffected)
}
