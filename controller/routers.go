package controller

import (
	"BMS-back-end/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func addRoutes() {
	e.GET("/doc/*", echoSwagger.WrapHandler)

	e.POST("/api/login", login)
	e.GET("/api/login", getLoginStatus)
	e.GET("/api/books", retrieveBooks)

	api := e.Group("/api", middleware.Auth)
	api.PUT("/book", storeBook)
	api.PUT("/books", storeBookCsv)

	api.PUT("/card", createCard)
	api.GET("/card", getCard)
	api.DELETE("/card", deleteCard)

	api.POST("/borrow", createBorrow)
	api.GET("/borrow", getBorrowed)
	api.POST("/return", updateReturn)
}
