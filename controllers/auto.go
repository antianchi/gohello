package controllers

import "github.com/astaxie/beego"

type AutoController struct {
	beego.Controller
}

func (c *AutoController) Simple() {

	c.Ctx.WriteString("你好，自动路由哦！！！！！！！！！！")
}
