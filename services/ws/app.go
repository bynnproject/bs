package ws

import (
	"bs/services/ws/handler"
	"golang.org/x/net/websocket"
	"net/http"
)

type WSApp struct {
	Port string
}

func (app *WSApp)Run() {
	http.Handle("/info", websocket.Handler(handler.InfoHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe(":"+app.Port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}
