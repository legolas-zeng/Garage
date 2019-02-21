package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_"github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

const (
	_DB_NAME 		= "data/garage.db"
	_SQLITE3_DRVIER = "sqlite3"
)

type Customer struct {
	Id 			int64		//用户id
	Name 		string		//客户名字
	Intime 		time.Time	//入店时间
	Outtime 	time.Time	`orm:"default('')"`//离店时间
	GoodsId		int64		//商品
	Money 		int64		//金额
	Ticket		bool		`orm:"default(0)"`//是否开票
	Carid		string		//车牌号
	Arctic		string		//车型
	Closer		bool		`orm:"default(0)"`//是否关闭
}

type Goods struct {
	Id 			int64
	Name 		string
	Count 		int64
	Price 		int64
}

func RegisterDB(){
	//if !com.IsExist(_DB_NAME){
	//	os.MkdirAll(path.Dir(_DB_NAME),os.ModePerm)
	//	os.Create(_DB_NAME)
	//}
	orm.Debug = true
	orm.RegisterModel(new(Customer),new(Goods))
	orm.RegisterDataBase("default", "mysql", "root:qq1005521@tcp(127.0.0.1:3306)/garage?charset=utf8", 30)
	//err := orm.RunSyncdb("default", true, true)
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println("数据库创建失败!!")
		fmt.Println(err)
	} else {
		fmt.Printf("数据库初始化已完成！！")
	}
	//orm.RegisterDriver(_SQLITE3_DRVIER,orm.DRSqlite)
	//orm.RegisterDataBase("default",_SQLITE3_DRVIER,_DB_NAME,10)

}



