package main

import (
	"log"
	"github.com/astaxie/beego"
	_ "dailymotion/routers"
)

func init() {
	log.Print("Init called")
}

func main() {
	beego.SetStaticPath("/images", "static/images")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.Run()
}
