package controller

import (
	"BMS-back-end/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func addRoutes() {
	e.GET("/doc/*", echoSwagger.WrapHandler)
	e.POST("/api/login", login)

	api := e.Group("/api", middleware.Auth)
	api.POST("/book", nil)
	api.POST("/books", storeBookCsv)
}
