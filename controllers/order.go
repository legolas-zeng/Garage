package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"garage/models"
	"time"
)

type OrderController struct {
	beego.Controller
}
type AddorderController struct {
	beego.Controller
}


func (c *OrderController) Get() {
	c.TplName = "order/order_tables.html"
}

func (c *AddorderController) Get() {
	c.TplName = "order/add_tables.html"
}

func (c *AddorderController) AddOrderApi(){
	fmt.Println("---------------")
	//ob := &models.Customer{}
	//json.Unmarshal(c.Ctx.Input.RequestBody, ob)
	Name := c.GetString("name")
	Arctic := c.GetString("arctic")
	Carid := c.GetString("carid")
	fmt.Println(Name,Arctic,Carid)
	o := orm.NewOrm()
	CustomerInfo := models.Customer{}
	CustomerInfo.Name = Name
	CustomerInfo.Arctic = Arctic
	CustomerInfo.Carid = Carid
	CustomerInfo.Intime = time.Now()
	CustomerInfo.Outtime = time.Now()
	_, err := o.Insert(&CustomerInfo)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"msg": err, "status": 1}
		c.ServeJSON()
		return
	}else {
		c.Data["json"] = map[string]interface{}{"msg": "工单提交成功", "status": 0}
		c.ServeJSON()
		return
	}
}
