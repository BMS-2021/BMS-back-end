package model

type Card struct {
	ID         uint64
	Name       string `gorm:"not null;size:10"`
	Department string `gorm:"not null;size:40"`
	Type       string `gorm:"not null;size:1"`
}

type CardReq struct {
	Name       string `json:"name"`
	Department string `json:"department"`
	Type       string `json:"type"`
}

func CreateCard(card *Card) error {
	result := db.Create(card)
	return result.Error
}

func DeleteCard(id uint64) error {
	result := db.Delete(id)
	return result.Error
}
