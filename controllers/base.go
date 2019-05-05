package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Get() {
	c.GetString("name", "未设置名字")
}
