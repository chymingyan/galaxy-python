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
	beego.Router("/config/targetdbs", &controllers.ConfigController{}, "get:TargetDbs")
	beego.Router("/config/singledb", &controllers.ConfigController{}, "post:SingleDb")
	beego.Router("/config/adddb", &controllers.ConfigController{}, "post:AddDb")
	beego.Router("/config/modifydb", &controllers.ConfigController{}, "post:ModifyDb")
	beego.Router("/config/deldb", &controllers.ConfigController{}, "post:DelDb")
	beego.Router("/config/testconndb", &controllers.ConfigController{}, "*:TestConnDb")
	//oraclecmdController
	beego.Router("/oraclecmd", &controllers.OracleCmdController{}, "get:OracleCommand")
	//linuxcmdController
	beego.Router("/linuxcmd", &controllers.LinuxCmdController{}, "get:LinuxCommand")
	//aixcmdController
	beego.Router("/aixcmd", &controllers.AixCmdController{}, "get:AixCommand")
	//inspectionController
	beego.Router("/insp", &controllers.InspectionController{}, "get:Insp")
	//exportContorller
	beego.Router("/export", &controllers.ExportController{}, "get:Export")
	//commonController
	beego.Router("/common/init", &controllers.CommonController{}, "post:InitLocalDb")
	beego.Router("/common/oracleversionarray", &controllers.CommonController{}, "*:OracleVersionArray")
	beego.Router("/common/aixversionarray", &controllers.CommonController{}, "*:AixVersionArray")
	beego.Router("/common/linuxversionarray", &controllers.CommonController{}, "*:LinuxVersionArray")
}
