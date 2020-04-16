package app

import (
	"bs/services/web/app/admin"
	"bs/services/web/app/node"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type HttpApp struct {
	Address string
}

func (app *HttpApp)Run()  {
	router := gin.Default()

	var store = cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session_id", store))
	router.GET("/test" , testfun)
	router.GET("/projects" , testfun)
	router.GET("/ips",testfun)
	router.GET("/login", admin.Login)
	router.GET("/runserver",testfun)
	router.GET("/nodes" , node.GetNodes)
	router.GET("/node/start" , node.StartOneNode)
	router.GET("/node/info",node.GetNodeServerInfo)

	router.Run(app.Address)
}

func testfun(c *gin.Context)  {
	s := sessions.Default(c)
	id := s.Get("userid")

	fmt.Println("id")
	if id == nil {
		s.Set("userid",1)
		s.Set("userid",2)
		c.JSON(200, gin.H{
			"message": "pong",
		})
		return
	}
	fmt.Println(id)
	switch id.(type) {
	case int:
		idtmp := id.(int)
		idtmp += 1
		fmt.Println(idtmp)
		s.Set("userid",idtmp)
		s.Save()
	}

	c.JSON(200, gin.H{
		"message": "pong",
	})
}