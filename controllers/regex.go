package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type RegexController struct {
	beego.Controller
}

func (c *RegexController) F1() {
	//
	fmt.Println("测试正则路由")
	id := c.GetString(":id")
	fmt.Println("收到参数：" + id)
	c.Ctx.WriteString("hello,child ,我收到你的消息:" + id)

}
