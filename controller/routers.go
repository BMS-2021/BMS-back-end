package controller

import (
	echoSwagger "github.com/swaggo/echo-swagger"
)

func addRoutes() {
	e.GET("/doc/*", echoSwagger.WrapHandler)

	api := e.Group("/api")
	api.POST("/login", login)
}
