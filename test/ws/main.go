package main

import "bs/services/ws"

func main()  {
	app := &ws.WSApp{Port:"8081"}
	app.Run()
}
