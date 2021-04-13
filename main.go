// @title Book Management System API
// @version 0.3
// @description This API will be used under staging environment.
// @host ralxyz.dev.zjuqsc.com
// @BasePath /api

package main

import (
	"BMS-back-end/controller"
	_ "BMS-back-end/docs"
	"BMS-back-end/model"
)

func main()  {
	model.Init()
	controller.InitWebFramework()
	controller.StartServer()
}
