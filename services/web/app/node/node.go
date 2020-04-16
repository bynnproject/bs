package node

import (
	"bs/clients/info"
	"bs/dao/entity"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetNodes(c *gin.Context)  {
	n := &entity.Node{}
	all , err := n.All()
	if err != nil {
		c.JSON(200 , gin.H{
			"status" : -1 ,
			"message" : "get nodes err",
		})
		return
	}

	c.JSON(200 , gin.H{
		"status" : 1 ,
		"message" : "success" ,
		"all" : all,
	})
}

func GetNodesByProject()  {
	
}

func DeleteNode()  {
	
}



func GetNodeServerInfo(c *gin.Context)  {
	nodeId := c.Query("nodeId")
	if nodeId == "" {
		c.JSON(200 , gin.H{
			"status" : -1 ,
			"message" : "nodeId can not be null",
		})
		return
	}

	n := &entity.Node{}
	n.NodeId, _ = strconv.ParseInt(nodeId , 10 , 64)

	bol , err := n.GetById()
	if err != nil || bol != true {
		c.JSON(200 , gin.H{
			"status" : -1 ,
			"message" : "nodeId can not be null",
		})
		return
	}
	address := n.Address
	infoclient := info.InfoClient{Address:address}
	err , allinfo := infoclient.GetInfo()
	if err != nil {
		c.JSON(200 , gin.H{
			"status" : -1 ,
			"message" : "get info err",
		})
		return
	}

	c.JSON(200 , gin.H{
		"status" : 1 ,
		"message" : "success",
		"info" : allinfo ,
	})
}

func GetNodeLogInfo()  {
	
}

func StartOneNode(c *gin.Context)  {
	address := c.Query("address")
	n := &entity.Node{Address:address}
	err := n.Start()
	if err != nil {
		c.JSON(200 , gin.H{
			"status" : -1 ,
			"message" : "start node err" ,
		})
		return
	}

	c.JSON(200 , gin.H{
		"status" : 1 ,
		"message" : "success",
	})
}


