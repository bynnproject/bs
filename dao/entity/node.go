package entity

import "errors"

type Node struct {

	NodeId int64 `xorm:"int(10) pk not null autoincr "`

	Address string `xorm:"char(50) not null"`

	Status int64 `xorm:"int(10)  not null  "`

	IsDel int64 `xorm:"int(10)  not null  "`
}

func (n *Node)All() ([]Node , error)  {
	x :=NewEngine()
	all := make([]Node , 0)
	err := x.Find(&all)
	if err != nil {
		return all , err
	}
	return all , nil
	
}

func (n *Node)Insert() (int64 ,error) {
	x := NewEngine()
	return x.ID(n.NodeId).Insert(n)
}

func (n *Node)Start() error {
	nodes , err := n.GetByAddress()
	if err != nil || len(nodes) >= 1 {
		return errors.New("start err")
	}
	if len(nodes) == 0 {
		_ , err = n.Insert()
		if err != nil {
			return err
		}
		return nil
	}

	node := nodes[0]
	if node.Status == 1 {
		return errors.New("start err")
	}// start status is 1

	n.NodeId = node.NodeId
	x := NewEngine()
	_  , err = x.ID(n.NodeId).Update(n)
	if err != nil {
		return err
	}
	return nil
}

func (n *Node)ChangeStatus() error  {
	_ , err := n.GetById()
	if err != nil {
		return err
	}
	x := NewEngine()
	tmp := &Node{
		NodeId: n.NodeId,
		Status:n.Status,
	}
	_, err = x.ID(tmp.NodeId).Update(tmp)
	if err != nil {
		return err
	}
	return nil
}

func (n *Node)ChangeDel() error {
	_ , err := n.GetById()
	if err != nil {
		return err
	}
	x := NewEngine()
	tmp := &Node{
		NodeId: n.NodeId,
		IsDel:n.IsDel,
	}
	_, err = x.ID(tmp.NodeId).Update(tmp)
	if err != nil {
		return err
	}
	return nil
}

func (* Node)AllNotDel() ([]Node, error)  {
	x :=NewEngine().Where("is_del = 0")
	all := make([]Node , 0)
	err := x.Find(&all)
	if err != nil {
		return all , err
	}
	return all , nil
}

func (n *Node)GetByAddress() ([]Node,error) {
	x := NewEngine()
	all := make([]Node , 0)
	x.ShowSQL(true)
	x.Where("address = ? and is_del = 0" ,  n.Address).Find(&all)
	return all , nil
}

func (n *Node)GetById() (bool , error)  {
	x := NewEngine()
	return x.ID(n.NodeId).Get(n)
}