package handler

import (
	"bs/clients/info"
	"bs/dao/entity"
	"encoding/json"
	"golang.org/x/net/websocket"
	"log"
	"time"
)

func InfoHandler(ws *websocket.Conn) {
	log.Println("into ws infohandler")
	msg := make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Receive: %s\n", msg[:n])

	//send_msg := "[" + string(msg[:n]) + "]"
	//m, err := ws.Write([]byte(send_msg))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("Send: %s\n", msg[:m])
	ninfo := entity.Node{}
	err = json.Unmarshal(msg[:n] , &ninfo)
	log.Println(ninfo.NodeId)
	log.Println(err)
	ninfo.GetById()
	infoclient := info.InfoClient{Address:ninfo.Address}
	//err , allinfo := infoclient.GetInfo()
	//allinfoJson , err := json.Marshal(allinfo)
	//_, err = ws.Write(allinfoJson)

	for ; ; {
		err , allinfo := infoclient.GetInfo()
		allinfoJson , err := json.Marshal(allinfo)
		_, err = ws.Write(allinfoJson)
		if err != nil {
			ws.Write(nil)
		}
		time.Sleep(time.Second * 5)
	}
	if err != nil {
		log.Println("ws send info err")
	}else {
		log.Println("ws send info success")
	}

}
