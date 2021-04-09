package model

type Admin struct {
	ID       uint64 `gorm:"not null;autoIncrement;primaryKey"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Contact  string `gorm:"not null"`
}

type AdminReq struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AdminResp struct {
	Name     string `json:"name"`
	Contact  string `json:"contact"`
}

func RetrieveAdmin(admin *Admin) error {
	result := db.Where(admin).First(admin)
	return result.Error
}
