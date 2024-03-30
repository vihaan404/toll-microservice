package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/vihaan404/toll-microservice/typess"
)

const wsEndpoint = "ws://127.0.0.1:30000/ws"

var sendInterval = time.Second

//	func sendOBUData (conn *websocket.Conn , data OBUData) error {
//	  return conn.WriteJSON(data)
//	}
func getCord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()

	return n + f
}

func genLetLongCord() (float64, float64) {
	return getCord(), getCord()
}

func genrateOBUIDs(n int) []int {
	ids := make([]int, n)

	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}

	return ids
}

func main() {
	obuIDs := genrateOBUIDs(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		for i := 0; i < len(obuIDs); i++ {
			lat, long := genLetLongCord()
			data := typess.OBUData{
				OBUID: obuIDs[i],
				Lat:   lat,
				Long:  long,
			}
			fmt.Printf("%+v\n", data)

			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}

		time.Sleep(sendInterval)
	}
}
