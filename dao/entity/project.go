package entity

type Project struct {
	ProjectId int64 `xorm:"int(10) pk not null autoincr "`

	Name string `xorm:"char(50) not null"`

	SavePath string `xorm:"char(50) not null"`

	LngId int64 `xorm:"int(10)  not null  "`

	Tag string `xorm:"char(50) not null"`



}

func (p *Project)Insert() (int64 , error) {
	x := NewEngine()
	return x.Insert(p)
}

func (p *Project)GetById()  (bool , error){
	return NewEngine().ID(p.ProjectId).Get(p)
}


func (p *Project)Delete() (int64 , error) {
	return NewEngine().ID(p.ProjectId).Delete(p)
}

func (p *Project)UpdateById() (int64 , error) {
	return NewEngine().ID(p.ProjectId).Update(p)
}

func (p *Project)All() []Project {
	all := make([]Project , 0)
	NewEngine().Find(&all)
	return all
}

func (p *Project)FindByName()  []Project{
	all := make([]Project , 0)
	NewEngine().Where("name = ?" , p.Name).Find(&all)
	return all
}
