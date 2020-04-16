package entity

type Admin struct {

	AdminId int64 `xorm:"int(10) pk not null autoincr "`

	Account string `xorm:"char(50) not null"`

	Pwd string `xorm:"char(50) not null"`

	IsDel int64 `xorm:"int(10)  not null  "`

}


func (a *Admin)Insert() ( int64, error) {
	x := NewEngine()
	return x.Insert(a)
}

func (a *Admin)GetById() (bool , error) {
	x := NewEngine()
	return x.ID(a.AdminId).Get(a)

}

func (a *Admin)All() []Admin {
	x := NewEngine()
	all := make([]Admin , 0)
	x.Find(&all)
	return all
}

func (a *Admin)GetByAccount()([]Admin , error)  {
	x := NewEngine()
	all := make([]Admin , 0)
	x.ShowSQL(true)
	x.Where("account = ? and is_del = 0" ,  a.Account).Find(&all)
	return all , nil

}

func (a *Admin)AllNotDel() []Admin {
	x := NewEngine()
	all := make([]Admin , 0)
	x.Where("is_del != 1")
	x.Find(&all)
	return all
}

func (a *Admin)ChangeDelStatus(status int64) (int64, error)  {
	x := NewEngine()
	tmp := &Admin{}
	x.ID(a.AdminId).Get(tmp)
	tmp.IsDel = status
	return x.ID(a.AdminId).Update(tmp)
}

func (a *Admin)Update() (int64, error) {
	x := NewEngine()
	tmp := &Admin{}
	x.ID(a.AdminId).Get(tmp)
	if a.Account != "" {
		tmp.Account = a.Account
	}
	if a.Pwd != "" {
		tmp.Pwd = a.Pwd
	}

	return x.ID(a.AdminId).Update(tmp)
}