package routers

import (
	"garage/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/index_v1", &controllers.IndexController{})
    beego.Router("/order", &controllers.OrderController{})
    beego.Router("/addorder", &controllers.AddorderController{})
    beego.Router("/add_order_api", &controllers.AddorderController{},"POST:AddOrderApi")

    erp := beego.NewNamespace("/erp",
    	beego.NSRouter("/stock",&controllers.ErpController{},"*:Get"),

		beego.NSInclude(
			&controllers.ErpController{},
    	),
	)
	beego.AddNamespace(erp)
}
