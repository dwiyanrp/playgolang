package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dwiyanrp/playgolang/go/scheduler/2/scheduler"
	"github.com/gin-gonic/gin"
)

func main() {
	s := scheduler.NewScheduler()

	r := gin.New()

	r.GET("/start", func(c *gin.Context) {
		taskID, _ := s.RunAt(time.Now().Add(10*time.Second), printNow)
		c.String(200, fmt.Sprint(taskID))
	})

	r.GET("/stop/:id", func(c *gin.Context) {
		taskID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		err := s.Cancel(taskID)
		if err != nil {
			c.String(200, fmt.Sprint(err))
		} else {
			c.String(200, fmt.Sprintf("Task %v stopped", taskID))
		}
	})

	r.Run()
}

func printNow() {
	log.Println("Lalala")
}
