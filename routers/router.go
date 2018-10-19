package routers

import (
	"ipp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/ipp", &controllers.LoginController{}, "get:Index")
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/", &controllers.MainController{})
	beego.Router("/insp", &controllers.InspectionController{}, "get:Insp")
	beego.Router("/export", &controllers.ExportController{}, "get:Export")
	beego.Router("/config", &controllers.ConfigController{}, "get:Config")
}
