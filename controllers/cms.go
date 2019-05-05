package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type CmsController struct {
	beego.Controller
}

//在 controller 的 method 方法上面写上 router 注@释（//   router）就可以了
// @router /allbook/:key [get]
func (c *CmsController) AllBook() {

	fmt.Println("收到请求")
	c.Ctx.WriteString("我是注解路由：" + c.GetString(":key", "未获取到输入参数"))
}
