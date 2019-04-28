package main

import (
	"github.com/astaxie/beego"
	_ "gohello/routers"
)

func main() {

	// 静态文件处理，参数路径，文件夹路径
	beego.SetStaticPath("/static1", "static")

	beego.Run()
}
