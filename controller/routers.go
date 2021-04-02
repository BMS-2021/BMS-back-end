package controller

import (
	"BMS-back-end/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func addRoutes() {
	e.GET("/doc/*", echoSwagger.WrapHandler)

	e.POST("/api/login", login)

	api := e.Group("/api", middleware.Auth)
	api.PUT("/book", storeBook)
	api.PUT("/books", storeBookCsv)
	api.GET("/books", retrieveBooks)
}
