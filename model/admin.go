package model

type Admin struct {
	ID       uint64 `gorm:"not null;autoIncrement;primaryKey"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Contact  string `gorm:"not null"`
}
