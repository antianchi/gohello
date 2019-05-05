package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "gohello/routers"
	"reflect"
)

func main() {

	var (
		a int
		b int16
		c     = 'A'
		d int = 65
	)

	e := string(d)
	fmt.Println(a, b)
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(c))

	fmt.Println("e:" + e)

	// 静态文件处理，参数路径，文件夹路径
	beego.SetStaticPath("/static1", "static")

	beego.Run()
}
