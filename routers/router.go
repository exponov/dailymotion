package routers

import (
	"log"
	"github.com/astaxie/beego"
	_"github.com/astaxie/beego/context"
	"dailymotion/controller"
	
)

func init() {
	log.Println("Routers Init Called")
	beego.Include(&controller.MainController{})
}
