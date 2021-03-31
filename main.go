package main

import (
	"BMS-back-end/controller"
	"BMS-back-end/model"
)

func main()  {
	model.Init()
	controller.InitWebFramework()
	controller.StartServer()
}
