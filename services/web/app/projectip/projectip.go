package projectip

import (
	"bs/dao/entity"
	"github.com/gin-gonic/gin"
	"strconv"
)

func All(c *gin.Context)  {
	p := &entity.ProjectIp{}
	ps , err := p.All()
	if err != nil {
		c.JSON(200 , gin.H{
			"status" : -1 ,
			"message" : "project ip get all err",
		})
		return
	}
	c.JSON(200 , gin.H{
		"status" : 1 ,
		"message" : "success" ,
		"all" : ps,
	})
}

func GetByProject(c *gin.Context)  {
	projectIpString := c.Query("projectIp")
	if projectIpString == "" {
		c.JSON(200 , gin.H{
			"status" : -1 ,
			"message" : "err",
		})
		return
	}
	projectId ,err := strconv.ParseInt(projectIpString , 10 , 64)
	p := &entity.ProjectIp{ProjectId:projectId}
	ps , err := p.FindByProjectId()
	if err != nil {
		c.JSON(200 , gin.H{
			"status" : -1 ,
			"message" : "project ip get all err",
		})
		return
	}
	c.JSON(200 , gin.H{
		"status" : 1 ,
		"message" : "success" ,
		"all" : ps,
	})
}

func GetByNode(c *gin.Context)  {
	nodeIdString := c.Query("projectIp")
	if nodeIdString == "" {
		c.JSON(200 , gin.H{
			"status" : -1 ,
			"message" : "err",
		})
		return
	}
	nodeId ,err := strconv.ParseInt(nodeIdString , 10 , 64)
	p := &entity.ProjectIp{NodeId:nodeId}
	ps , err := p.FindByProjectId()
	if err != nil {
		c.JSON(200 , gin.H{
			"status" : -1 ,
			"message" : "project ip get all err",
		})
		return
	}
	c.JSON(200 , gin.H{
		"status" : 1 ,
		"message" : "success" ,
		"all" : ps,
	})
}

func Start(c *gin.Context)  {

}
