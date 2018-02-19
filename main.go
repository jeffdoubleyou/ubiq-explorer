package main

import (
	"github.com/astaxie/beego"
	//"ubiq-explorer/controllers"
	_ "ubiq-explorer/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
