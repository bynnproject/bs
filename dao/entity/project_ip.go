package entity

type ProjectIp struct {
	Id          int64     `xorm:"int(10) pk not null autoincr "`

	Address       string    `xorm:"char(50) not null"`

	ProjectId int64 `xorm:"int(10)  not null  "`

	PathId int64 `xorm:"int(10)  not null  "`

	Tag string `xorm:"char(50) not null"`

	NodeId int64 `xorm:"int(10) "`
}

func (p *ProjectIp)Insert() (int64 , error) {
	return NewEngine().Insert(p)
}

//func (p *ProjectIp)Update()  {
//
//}

func (p *ProjectIp)GetById() (bool,error) {
	return NewEngine().ID(p.Id).Get(p)
}

func (p *ProjectIp)All() ([]ProjectIp , error) {
	all := make([]ProjectIp , 0)
	err := NewEngine().Find(&all)
	if err != nil {
		return nil , err
	}
	return all , nil
}

func (p *ProjectIp)FindByProjectId() ([]ProjectIp , error) {
	all := make([]ProjectIp , 0)
	err := NewEngine().Where("project_id = ?" ,  p.ProjectId).Find(&all)
	if err != nil {
		return nil , err
	}
	return all , nil
}

func (p *ProjectIp)Delete()  {

}

func (p *ProjectIp)GetByNodeId() ([]ProjectIp , error) {
	all := make([]ProjectIp , 0)
	err := NewEngine().Where("node_id = ?" ,  p.ProjectId).Find(&all)
	if err != nil {
		return nil , err
	}
	return all , nil
}
