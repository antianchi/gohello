package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"gohello/controllers"
)

func init() {
	fmt.Println("初始化router")
	// 执行路由注册
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{}, "get:Get2")
}
