package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gohello/controllers"
)

func init() {
	fmt.Println("初始化router")
	// 执行路由注册
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{}, "get:Get2")

	// 路由测试配置
	// 1、基础路由设置
	// 基本GET
	beego.Get("/baseget", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello ,我是基本GET请求"))
	})
	beego.Router("/base1", &controllers.BaseController{})

	// 固定路由方式，典型的RESTFul风格，根据请求的方法（http方法）调用不同的函数
	beego.Router("/resttest", &controllers.RestController{})

	// 正则路由
	beego.Router("/regex/?:id", &controllers.RegexController{}, "get:F1")

	// 自定义方法及RestFull风格
	beego.Router("/api", &controllers.RestController{}, "get:ApiGet;post:ApiPost")

	// 自动路由
	beego.AutoRouter(&controllers.AutoController{})

	// 注解路由
	beego.Include(&controllers.CmsController{})

	// Namespace
	// 初始化namepace
	ns := beego.NewNamespace("v1", beego.NSCond(func(ctx *context.Context) bool {
		if ctx.Input.Domain() == "api.beego.me" {
			return true
		}
		return false
	}))

	// 注册Namespace
	beego.AddNamespace(ns)

}
