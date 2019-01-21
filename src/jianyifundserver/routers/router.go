package routers

import (
	"jianyifundserver/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/userContract", &controllers.UserContractController{})
}
