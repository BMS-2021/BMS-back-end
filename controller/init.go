package controller

import (
	"BMS-back-end/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var e *echo.Echo

func InitWebFramework() {
	e = echo.New()
	e.HideBanner = true
	addRoutes()
	e.Validator = &utils.CustomValidator{}

	logrus.Info("echo framework initialized")
}

func StartServer() {
	e.Logger.Fatal(e.Start(":1323"))
}
