package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type RestController struct {
	beego.Controller
}

func (c *RestController) Get() {
	c.GetString("name", "未设置名字")
	fmt.Println("调用了GET请求")
	c.Ctx.WriteString("hello,rest ,我收到你的消息")
}

func (c *RestController) Post() {
	c.GetString("name", "未设置名字")
	fmt.Println("调用了Post请求")
	c.Ctx.WriteString("hello,rest ,我收到你的消息")
}
func (c *RestController) Delete() {
	c.GetString("name", "未设置名字")
	fmt.Println("调用了delete请求")
	c.Ctx.WriteString("hello,rest ,我收到你的消息")
}
func (c *RestController) Patch() {
	c.GetString("name", "未设置名字")
	fmt.Println("调用了Patch请求")
	c.Ctx.WriteString("hello,rest ,我收到你的消息")
}
func (c *RestController) Put() {
	c.GetString("name", "未设置名字")
	fmt.Println("调用了Put请求")
	c.Ctx.WriteString("hello,rest ,我收到你的消息")
}
func (c *RestController) ApiGet() {
	c.GetString("name", "未设置名字")
	fmt.Println("调用了ApiGet请求")
	c.Ctx.WriteString("hello,rest ,我收到你的消息")
}

func (c *RestController) ApiPost() {
	c.GetString("name", "未设置名字")
	fmt.Println("调用了ApiPost请求")
	c.Ctx.WriteString("hello,rest ,我收到你的消息")
}
