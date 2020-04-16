package main

import (
	"bs/pkg/log_tail"
	"fmt"
	"bs/pkg/info"
	"log"
	"os"
	"bs/cmd"
	"time"
	"bs/dao/entity"
	infoserver "bs/services/info/server"
	infocli "bs/clients/info"
)

func main() {
	InfoClient()
}

func InfoServer()  {
	app := &infoserver.InfoServer{
		Port:                    "8088",
	}
	app.Run()
}

func InfoClient()  {
	cli := infocli.InfoClient{Address:"127.0.0.1:8088"}
	err ,info := cli.GetInfo()
	if err ==nil {
		fmt.Println(info)
	}else {
		fmt.Println("err %v " , err)
	}
}

func Store()  {
	app := cmd.NewCmdApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Dao()  {
	l := &entity.Project{Name:"LVXIAOLUO"}
	l.Insert()
	fmt.Println(l.FindByName())
	l.ProjectId = 1
	fmt.Println(l.Delete())
	fmt.Println(l.FindByName())

	//log.Println(aff)
}

func Log()  {
	log_tail.NewTail("/Users/bynn/bstest1.txt")
}

func Info()  {
	info.GetHosts()
	fmt.Println(info.GetCpu())
	fmt.Println(info.GetDisk())
	fmt.Println(info.GetHosts())
	fmt.Println(info.GetMem())
	fmt.Println(info.GetProcessInfo())


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