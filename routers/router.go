package routers

import (
	"firstBeegoDemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/helloworld", &controllers.HelloController{})
}
