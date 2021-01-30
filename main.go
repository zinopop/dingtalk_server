package main

import (
	_ "dingtalk_server/boot"
	_ "dingtalk_server/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	//测试
	g.Server().Run()
}
