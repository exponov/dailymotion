package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["dailymotion/controller:MainController"] = append(beego.GlobalControllerRouter["dailymotion/controller:MainController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dailymotion/controller:MainController"] = append(beego.GlobalControllerRouter["dailymotion/controller:MainController"],
		beego.ControllerComments{
			Method: "RequestAd",
			Router: `/request`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
