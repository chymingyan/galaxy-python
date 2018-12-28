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
	beego.Router("/tree", &controllers.MainController{}, "get:Tree")

	//configController
	beego.Router("/config", &controllers.ConfigController{}, "*:Config")
	beego.Router("/config/groups", &controllers.ConfigController{}, "*:Groups")
	beego.Router("/config/group", &controllers.ConfigController{}, "*:Group")
	beego.Router("/config/addgroup", &controllers.ConfigController{}, "*:AddGroup")
	beego.Router("/config/modifygroup", &controllers.ConfigController{}, "*:ModifyGroup")
	beego.Router("/config/delgroup", &controllers.ConfigController{}, "*:DelGroup")
	beego.Router("/config/targethosts", &controllers.ConfigController{}, "*:TargetHosts")
	beego.Router("/config/singlehost", &controllers.ConfigController{}, "*:SingleHost")
	beego.Router("/config/addhost", &controllers.ConfigController{}, "*:AddHost")
	beego.Router("/config/modifyhost", &controllers.ConfigController{}, "post:ModifyHost")
	beego.Router("/config/delhost", &controllers.ConfigController{}, "*:DelHost")
	beego.Router("/config/testconnhost", &controllers.ConfigController{}, "*:TestConnHost")
	beego.Router("/config/targetdbs", &controllers.ConfigController{}, "get:TargetDbs")
	beego.Router("/config/singledb", &controllers.ConfigController{}, "post:SingleDb")
	beego.Router("/config/adddb", &controllers.ConfigController{}, "post:AddDb")
	beego.Router("/config/modifydb", &controllers.ConfigController{}, "post:ModifyDb")
	beego.Router("/config/deldb", &controllers.ConfigController{}, "post:DelDb")
	beego.Router("/config/testconndb", &controllers.ConfigController{}, "*:TestConnDb")

	//oraclecmdController
	beego.Router("/oradb", &controllers.OracleCmdController{}, "*:OraDb")
	beego.Router("/oraclecmd/cmdindex", &controllers.OracleCmdController{}, "get:OracleCmdIndex")
	beego.Router("/oraclecmd/commands", &controllers.OracleCmdController{}, "*:Commands")
	beego.Router("/oraclecmd/command", &controllers.OracleCmdController{}, "*:Command")
	beego.Router("/oraclecmd/add", &controllers.OracleCmdController{}, "*:Add")
	beego.Router("/oraclecmd/modify", &controllers.OracleCmdController{}, "*:Modify")
	beego.Router("/oraclecmd/del", &controllers.OracleCmdController{}, "*:Del")

	//linuxcmdController
	beego.Router("/linuxcmd/cmdindex", &controllers.LinuxCmdController{}, "get:LinuxCmdIndex")
	beego.Router("/linuxcmd/commands", &controllers.LinuxCmdController{}, "*:LinuxCommands")
	beego.Router("/linuxcmd/command", &controllers.LinuxCmdController{}, "*:LinuxCommand")
	beego.Router("/linuxcmd/add", &controllers.LinuxCmdController{}, "*:AddCommand")
	beego.Router("/linuxcmd/modify", &controllers.LinuxCmdController{}, "*:ModifyCommand")
	beego.Router("linuxcmd/del", &controllers.LinuxCmdController{}, "*:DelCommand")
	//aixcmdController
	beego.Router("/aixcmd/cmdindex", &controllers.AixCmdController{}, "get:AixCmdIndex")
	beego.Router("/aixcmd/commands", &controllers.AixCmdController{}, "*:AixCommands")
	beego.Router("/aixcmd/command", &controllers.AixCmdController{}, "*:AixCommand")
	beego.Router("/aixcmd/add", &controllers.AixCmdController{}, "*:AddCommand")
	beego.Router("/aixcmd/modify", &controllers.AixCmdController{}, "*:ModifyCommand")
	beego.Router("/aixcmd/del", &controllers.AixCmdController{}, "*:DelCommand")
	//inspectionController
	beego.Router("/insp", &controllers.InspectionController{}, "*:Insp")
	beego.Router("/insp/showhost", &controllers.InspectionController{}, "get:ShowHostByDbId")
	//exportContorller
	beego.Router("/export", &controllers.ExportController{}, "get:Export")
	//commonController
	beego.Router("/common/init", &controllers.CommonController{}, "post:InitLocalDb")
	beego.Router("/common/oracleversionarray", &controllers.CommonController{}, "*:OracleVersionArray")
	beego.Router("/common/aixversionarray", &controllers.CommonController{}, "*:AixVersionArray")
	beego.Router("/common/linuxversionarray", &controllers.CommonController{}, "*:LinuxVersionArray")
	beego.Router("/common/import", &controllers.CommonController{}, "*:Import")
	beego.Router("/common/export", &controllers.CommonController{}, "*:Export")
}
