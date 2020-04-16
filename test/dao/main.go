package main

import (
	"bs/dao/entity"
	"fmt"
)

func main()  {
	//p := entity.Project{
	//	ProjectId: 2,
	//	Name:      "",
	//	SavePath:  "",
	//	LngId:     0,
	//	Tag:       "",
	//}
	//all := p.All()
	//fmt.Println(all)
	//a := entity.Admin{}
	//al := a.All()
	//fmt.Println(al)


	a := make([]entity.Node , 8)
	n := entity.Node{
		NodeId: 1,
	}
	//a[7] = n
	a = append(a, n)
	fmt.Println(a)


}
