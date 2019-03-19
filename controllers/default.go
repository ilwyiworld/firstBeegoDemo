package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego"
	c.Data["Email"] = "ilwyiworld@gmail.com"
	c.TplName = "index.tpl"
}

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	//c.Ctx.WriteString("Hello World")
	c.Data["json"]=map[string]interface{}{
		"key":"value",
		"age":14,
		"isDeleted":true,
	}
	c.ServeJSON()
}
