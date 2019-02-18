package main

import (
	_ "jianyifundserver/routers"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods: []string{"PUT", "POST", "GET", "DELETE", "OPTIONS"},        
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},     
		AllowCredentials: true, })) 
	beego.BConfig.Listen.Graceful=false
	beego.BConfig.Listen.HTTPPort = 8080
	beego.SetLogFuncCall(true)
	beego.Run()
}
