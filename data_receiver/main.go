package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/vihaan404/toll-microservice/typess"
)

func main() {
	recv := NewDataReceiver()
	fmt.Println("data receiver is working fine")
	http.HandleFunc("/ws", recv.handlerWS)

	http.ListenAndServe(":30000", nil)
}

type DataReceiver struct {
	msgch chan typess.OBUData
	conn  *websocket.Conn
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgch: make(chan typess.OBUData, 128),
	}
}

func (dr *DataReceiver) handlerWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}

	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	dr.conn = conn

	go dr.wsReceiverLoop()
}

func (dr *DataReceiver) wsReceiverLoop() {
	fmt.Println("OBU connected client connected")
	for {
		var data typess.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			println("read error", err)
			continue
		}

		fmt.Printf("reaceived OBU data from id [%d] :: <lat %.2f , long %.2f>\n", data.OBUID, data.Lat, data.Long)
		dr.msgch <- data
	}
}
