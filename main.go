package main

import (
	_ "garage/routers"
	"github.com/astaxie/beego"
	"garage/models"
	"github.com/astaxie/beego/orm"
)

func init(){
	models.RegisterDB()
}

func main() {
	beego.SetStaticPath("/static/images", "images")
	beego.SetStaticPath("/static/css", "css")
	beego.SetStaticPath("/static/js", "js")
	orm.Debug = true
	orm.RunSyncdb("default",false,true)
	beego.Run()
}


