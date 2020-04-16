package main

import "bs/services/web/app"

func main()  {
	a := &app.HttpApp{Address:"127.0.0.1:8080"}
	a.Run()
}
