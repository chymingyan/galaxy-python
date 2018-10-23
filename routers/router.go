package routers

import (
	"ipp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//loginController
	beego.Router("/signin", &controllers.LoginController{}, "get:Index")
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/signout", &controllers.LoginController{}, "post:SignOut")
	//mainController
	beego.Router("/", &controllers.MainController{})
	beego.Router("/chat", &controllers.MainController{}, "post:Chat")
	beego.Router("/tree", &controllers.MainController{}, "post:Tree")
	//configController
	beego.Router("/config", &controllers.ConfigController{}, "get:Config")
	//oraclecmdController
	//linuxcmdController
	//aixcmdController
	//inspectionController
	beego.Router("/insp", &controllers.InspectionController{}, "get:Insp")
	//exportContorller
	beego.Router("/export", &controllers.ExportController{}, "get:Export")

}
