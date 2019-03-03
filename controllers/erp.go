package controllers

import (
	"github.com/astaxie/beego"
	"garage/models"
	"fmt"
)

type ErpController struct {
	beego.Controller
}

func (c *ErpController) Get() {
	beego.ReadFromRequest(&c.Controller)

	OrderInfo := &models.Erp{}
	erps:= OrderInfo.FindAllErpInfo()
	fmt.Println(erps)
	c.Data["erps"] = erps
	c.TplName = "erp/stock.html"
	c.Render()
}
