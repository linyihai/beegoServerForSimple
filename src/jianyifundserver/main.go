package main

import (
	_ "jianyifundserver/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.Listen.Graceful=false
	beego.BConfig.Listen.HTTPPort = 8080
	beego.SetLogFuncCall(true)
	beego.Run()
}
