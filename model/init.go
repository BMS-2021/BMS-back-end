package model

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	ConnectDatabase()
}

func ConnectDatabase() {
	viper.SetConfigName("conf")
	loginInfo := viper.GetStringMapString("sql")

	dbArgs := loginInfo["username"] + ":" + loginInfo["password"] +
		"@(localhost)/" + loginInfo["db_name"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dbArgs), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}
}
