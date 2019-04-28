package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

//Get2
func (c *UserController) Get2() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	c.Ctx.WriteString("hello,child ,我收到你的消息")
}
