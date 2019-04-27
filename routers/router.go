package routers

import (
	"github.com/astaxie/beego"
	"itemdetail/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
