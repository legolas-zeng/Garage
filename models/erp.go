package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Erp struct {
	Id			int
	Name 		string		//商品名称
	Count 		int			//总库存
	Sale		int			//销售额
	Price 		int64		//单价
}

type EroLog struct {
	Id   		int
	ConTime 	time.Time	//消费时间
	ConUser		string		//消费者名字
	ConCount	int			//消费数量
	ConPrice	int			//消费单价
}

//TODO 联表查询
func (c *Erp) FindAllErpInfo() ([]*Erp) {

	var erp []*Erp
	o := orm.NewOrm()
	o.QueryTable("erp").RelatedSel().All(&erp)

	//common.PanicIf(err)
	return erp

}
