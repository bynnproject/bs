package admin

import (
	"bs/dao/entity"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func Login(c *gin.Context)  {
	log.Println("into login")
	account := c.Query("account")
	password := c.Query("password")
	admin := entity.Admin{Account:account}
	admins , err := admin.GetByAccount()

	if err != nil || len(admins) != 1 {
		c.JSON(200, gin.H{
			"status": "-1",
			"message":"get err",
		})
		return
	}
	tmp := admins[0]
	if tmp.Pwd != password {
		c.JSON(200, gin.H{
			"status": "-1",
			"message":"get err",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("isAdmin" , true)
	session.Set("admin" , tmp)
	session.Save()
	c.JSON(200, gin.H{
		"status": "1",
		"message":"get success",
		"admin" : tmp,
	})
}

func AllNotDel(c *gin.Context)  {
	log.Println("into allNotDel")

	admin := &entity.Admin{}
	all := admin.AllNotDel()
	c.JSON(200 , gin.H{
		"status" : 1,
		"message" : "success",
		"allNotDel" : all,
	})
}



func All(c *gin.Context)  {
	log.Println("into allNotDel")

	admin := &entity.Admin{}
	all := admin.All()
	c.JSON(200 , gin.H{
		"status" : 1,
		"message" : "success",
		"allNotDel" : all,
	})
}

func Update(c *gin.Context)  {
	account := c.Query("account")
	adminIdString := c.Query("adminId")
	password := c.Query("password")
	if adminIdString == "" {
		c.JSON(200 , gin.H{
			"status" : -1,
			"message" : "adminId is null" , 
			
		})
		return 
	}
	adminId , err := strconv.ParseInt(adminIdString, 10 , 64)
	if err != nil {
		c.JSON(200 , gin.H{
			"status" : -1 ,
			"message" : "adminId transport err",
		})
		return
	}

	tmp := &entity.Admin{AdminId:adminId}
	tmp.GetById()
	if account != "" {
		tmp.Account = account
	}

	if password != "" {
		tmp.Pwd = password
	}
	_ , err = tmp.Update()
	if err != nil {
		c.JSON(200, gin.H{
			"status" : -1,
			"message" : "update err",

		})
		return
	}

	c.JSON(200,gin.H{
		"status" : 1,
		"message" : "success",
	})
}

func Insert(c *gin.Context)  {
	account := c.Query("account")
	password := c.Query("password")
	if account == "" || password == "" {
		c.JSON(200,gin.H{
			"status" : -1 ,
			"messgage" : "account and password can not be null",
		})
		return
	}
	tmp := &entity.Admin{Account:account , Pwd:password, IsDel:0}
	_ , err := tmp.Insert()
	if err != nil {
		c.JSON(200,gin.H{
			"status" : -1 ,
			"message" : "insert admin err",
		})
		return
	}

	c.JSON(200 , gin.H{
		"status" : 1 ,
		"message" : "success" ,
	})
}

func ChangeDelStatus(c *gin.Context)  {
	isDelString := c.Query("isDel")
	adminIdString := c.Query("adminId")
	if isDelString == "" || adminIdString == "" {
		c.JSON(200 , gin.H{
			"status" : -1 ,
			"message" : "isDel and adminId can not be null",
		})
		return
	}

	isDel , isDelErr := strconv.ParseInt(isDelString , 10 , 64)
	adminId , adminIdErr := strconv.ParseInt(adminIdString , 10 , 64)
	if isDelErr != nil || adminIdErr != nil {
		c.JSON(200, gin.H{
			"status" : -1 ,
			"message" : "isDelErr or adminIdErr'",
		})
		return
	}

	tmp := &entity.Admin{AdminId:adminId , IsDel:isDel}
	tmp.ChangeDelStatus(isDel)
	c.JSON(200, gin.H{
		"status" : 1,
		"message" : "success",
	})

}
