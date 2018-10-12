package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "https://github.com/chymingyan"
	c.Data["Email"] = "chen_haiyan@hotmail.com"
	c.TplName = "index.tpl"
}
