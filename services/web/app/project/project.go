package project

import (
	"bs/dao/entity"
	"github.com/gin-gonic/gin"
)

func All(c *gin.Context)  {
	p := & entity.Project{}
	all := p.All()
	c.JSON(200 , gin.H{
		"status" : 1 ,
		"message" : "success",
		"all" : all,
	})
}
