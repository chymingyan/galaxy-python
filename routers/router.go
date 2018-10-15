package routers

import (
	"ipp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "get:Login")
	beego.Router("/insp", &controllers.InspectionController{}, "get:Insp")
	beego.Router("/export", &controllers.ExportController{}, "get:Export")
	beego.Router("/config", &controllers.ConfigController{}, "get:Config")
}
