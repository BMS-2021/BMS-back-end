package controller

import (
	echoSwagger "github.com/swaggo/echo-swagger"
)

func addRoutes() {
	e.GET("/doc/*", echoSwagger.WrapHandler)

	e.POST("/login", login)

	// api := e.Group("/api")
}
