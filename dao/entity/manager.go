package entity


import (
	//"bs/dao/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

func NewEngine()  *xorm.Engine{
	x, err := xorm.NewEngine("mysql", "root:123456@tcp(119.23.187.108:3306)/bs?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	if err := x.Sync(new(Path), new(Lng), new(Project), new(ProjectIp)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
	return x;
}
