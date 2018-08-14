package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/dwiyanrp/playgolang/go/scheduler/2/task"
)

type Scheduler struct {
	funcRegistry *task.FuncRegistry
	channel      chan *task.Task
	tasks        map[int64]*task.Task
}

func NewScheduler() *Scheduler {
	funcRegistry := task.NewFuncRegistry()
	scheduler := &Scheduler{
		funcRegistry: funcRegistry,
		channel:      make(chan *task.Task),
		tasks:        make(map[int64]*task.Task),
	}

	go scheduler.HandleScheduler()

	return scheduler
}

func (scheduler *Scheduler) RunAt(time time.Time, function task.Function) (int64, error) {
	funcMeta, err := scheduler.funcRegistry.Add(function)
	if err != nil {
		return 0, err
	}

	task := task.NewTask(funcMeta, time)
	scheduler.tasks[task.TaskID] = task

	log.Printf("Task %v scheduled at %v", task.TaskID, task.RunAt)
	scheduler.channel <- task

	return task.TaskID, nil
}

func (scheduler *Scheduler) Cancel(taskID int64) error {
	task, found := scheduler.tasks[taskID]
	if !found {
		return fmt.Errorf("Task %v not found", taskID)
	}

	task.Stop = true
	scheduler.channel <- task
	delete(scheduler.tasks, taskID)
	log.Printf("Task %v stopped", taskID)
	return nil
}

func (scheduler *Scheduler) HandleScheduler() {
	for {
		select {
		case task := <-scheduler.channel:
			go func() {
				<-task.Timer.C
				task.Run()
			}()

			if task.Stop {
				task.Timer.Stop()
			}
		}
	}
}
