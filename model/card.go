package model

import "gorm.io/gorm"

type Card struct {
	ID         uint64
	Name       string `gorm:"not null;size:10"`
	Department string `gorm:"not null;size:40"`
	Type       string `gorm:"not null;size:1"`
}

type CardReq struct {
	Name       string `json:"name" validate:"required"`
	Department string `json:"department" validate:"required"`
	Type       string `json:"type" validate:"required"`
}

func CreateCard(card *Card) error {
	result := db.Create(card)
	return result.Error
}

func GetCard(id uint64) (*Card, error) {
	dbCard := Card{ID: id}
	result := db.First(&dbCard)
	return &dbCard, result.Error
}

func DeleteCard(id uint64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("card_id = ?", id).Delete(&Borrow{}).Error; err != nil {
			return err
		}
		err := tx.Delete(&Card{}, id).Error
		return err
	})
}
