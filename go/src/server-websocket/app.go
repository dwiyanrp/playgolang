package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var userConnected = int64(0)
var messageCount = int64(0)
var messageMinuteCount = int64(0)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	go f()
	go f60()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})
	r.Run(":8080")
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	userConnected++

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			userConnected--
			break
		}

		if msg != nil {
			// fmt.Println(string(msg))
			messageCount++
			messageMinuteCount++
		}

		conn.WriteMessage(t, msg)
	}
}

func f() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("%v user	| %v rps	| %v \n", userConnected, messageCount, time.Now().Format(time.RFC3339))
		messageCount = 0

		if userConnected < 0 {
			userConnected = 0
		}
	}
}

func f60() {
	for {
		time.Sleep(30 * time.Second)
		fmt.Printf("%v user	| %v rpm	| %v rps | %v \n", userConnected, messageMinuteCount, messageMinuteCount/30, time.Now().Format(time.RFC3339))
		messageMinuteCount = 0
	}
}
