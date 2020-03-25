package entity

type Path struct {

	PathId          int64     `xorm:"int(10) pk not null autoincr "`

	PathName       string    `xorm:"char(50) not null"`

}



func (p *Path)Insert() ( int64, error) {
	x := NewEngine()
	return x.Insert(p)
}

func (p *Path)GetById() (bool , error) {
	x := NewEngine()
	return x.ID(p.PathId).Get(p)

}

func (p *Path)All() []Path {
	x := NewEngine()
	all := make([]Path , 0)
	x.Find(&all)
	return all
}
