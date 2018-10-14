package routers

import (
	"ipp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "get:Login")
	beego.Router("/inspection", &controllers.InspectionController{}, "post:insp")
	beego.Router("/export", &controllers.ExportController{})
	beego.Router("/config", &controllers.ConfigController{})
}
