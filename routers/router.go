package routers

import (
	"ap/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/find/", &controllers.AccountController{}, "get:Find")
}
