package main

import (
	"fmt"
	"log"
	"os"
	"bs/cmd"
	"time"
	"bs/dao/entity"
)

func main() {
	Store()
}

func Store()  {
	app := cmd.NewCmdApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Dao()  {
	//x := dao.NewEngine()
	//lng := entity.Lng{
	//	LngName: "GO",
	//}
	//aff , err := x.Insert(lng)
	//if err != nil {
	//	log.Println(err)
	//}
	l := &entity.Project{Name:"LVXIAOLUO"}
	l.Insert()
	fmt.Println(l.FindByName())
	l.ProjectId = 1
	fmt.Println(l.Delete())
	fmt.Println(l.FindByName())

	//log.Println(aff)
}

type Member struct {

	Memberid          int64     `xorm:"int(10) pk not null autoincr 'member_id'"`

	Member_name       string    `xorm:"char(50) not null"`

	Member_pwd        string    `xorm:"char(32) not null"`

	Member_count      int64     `xorm:"smallint(6) not null default 0 "`

	Member_ok         string    `xorm:"varchar(50) not null"`

	Member_del        int64     `xorm:"bigint(1) not null default 0 "`

	Member_email      string    `xorm:"varchar(40) not null"`

	Member_logintime  time.Time `xorm:"updated"`

	Member_createtime int64     `xorm:"int(11) not null default 0"`

	Member_updatetime int64     `xorm:"int(11) not null default 0"`

	Member_ip         string    `xorm:"varchar(40) not null"`

	Member_version    int64       `xorm:"version"`

}