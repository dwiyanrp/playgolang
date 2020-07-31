package main

import (
	"log"
	"time"

	"github.com/rakanalh/scheduler"
	"github.com/rakanalh/scheduler/storage"
)

func TaskWithoutArgs() {
	log.Println(time.Now().String())
}

func TaskWithArgs(message string) {
	log.Println("TaskWithArgs is executed. message:", message)
}

func main() {
	storage := storage.NewMemoryStorage()

	s := scheduler.New(storage)

	// Start a task without arguments
	if _, err := s.RunEvery(2*time.Second, TaskWithoutArgs); err != nil {
		log.Fatal(err)
	}

	if _, err := s.RunAt(time.Now().Add(5*time.Second), TaskWithArgs, "Specific"); err != nil {
		log.Fatal(err)
	}

	s.Start()
	s.Wait()
}
