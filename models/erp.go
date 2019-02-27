package models

import (
	"time"
)

type Erp struct {
	Id			int
	Name 		string
	count 		int
	Price 		int64		//单价
}

type EroLog struct {
	Id   		int
	ConTime 	time.Time	//消费时间
	ConUser		string		//消费者名字
	ConCount	int			//消费数量
	ConPrice	int			//消费单价
}
