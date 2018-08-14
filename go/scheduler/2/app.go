package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	taskID   int64
	isActive bool
	timer    *time.Timer
	runAt    time.Time
	action   func(string)
	message  string
}

var channel = make(chan Task)
var mapTask = make(map[int64]Task, 0)

func main() {
	r := gin.New()
	r.GET("/start", func(c *gin.Context) {
		runAt := time.Now().Add(20 * time.Second)
		taskID := registerTask(runAt, c.Query("msg"))

		log.Printf("Task %v scheduled at %v", taskID, runAt)
		c.String(200, fmt.Sprintf("%v", taskID))
	})

	r.GET("/stop/:id", func(c *gin.Context) {
		taskID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		task := mapTask[taskID]
		task.isActive = false

		channel <- task
		c.String(200, fmt.Sprintf("Task %v stopped", taskID))
	})

	go handleScheduler()

	r.Run()
}

func handleScheduler() {
	for {
		select {
		case task := <-channel:
			timer := time.NewTimer(task.runAt.Sub(time.Now()))

			go func() {
				<-timer.C
				task.action(task.message)
				delete(mapTask, task.taskID)
			}()

			if task.isActive == false {
				timer.Stop()
				log.Printf("Task %v stopped", task.taskID)
				delete(mapTask, task.taskID)
			}
		}
	}
}

func registerTask(runAt time.Time, param string) int64 {
	taskID := time.Now().UnixNano()

	task := Task{
		taskID:   taskID,
		isActive: true,
		timer:    time.NewTimer(runAt.Sub(time.Now())),
		runAt:    runAt,
		action:   func(msg string) { log.Println(msg) },
		message:  param,
	}

	mapTask[taskID] = task
	channel <- task

	return taskID
}
