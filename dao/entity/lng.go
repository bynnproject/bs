package entity

type Lng struct {
	LngId          int64     `xorm:"int(10) pk not null autoincr "`

	LngName       string    `xorm:"char(50) not null"`
}

func (l *Lng)Insert() ( int64, error) {
	x := NewEngine()
	return x.Insert(l)
}

func (l *Lng)GetById() (bool , error) {
	x := NewEngine()
	return x.ID(l.LngId).Get(l)

}

func (l *Lng)All() []Lng {
	x := NewEngine()
	all := make([]Lng , 0)
	x.Find(&all)
	return all
}