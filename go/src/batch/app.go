package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Foo struct {
	Data string
}

var (
	done = make(chan os.Signal, 1)

	batch     = make([]Foo, 0, 10)
	batchChan = make(chan *Foo)
)

func main() {
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// case1()
	// case2()
	case3()

	<-done
	graceful()
}

func graceful() {
	batchInsert(batch) // Make sure batch is empty & all data executed
}

func insertOne() {
	fmt.Println("1 inserted")
}

func batchInsert(mass []Foo) {
	fmt.Println(len(mass), "inserted")
}

// Insert one
func case1() {
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(100 * time.Millisecond)
			insertOne()
		}
	}()
}

// Insert batch, executed after being fully filled
// Will wait a long time if not fully filled
func case2() {
	go func() {
		var tempBatch []Foo

		for {
			select {
			case f := <-batchChan:
				batch = append(batch, *f)

				if len(batch) == 10 {
					tempBatch = batch
					batch = make([]Foo, 0, 10)
					batchInsert(tempBatch)
				}
			}
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(222 * time.Millisecond)
			batchChan <- &Foo{}
		}
	}()
}

// Insert batch, executed after being filled or ticker
// If batch not full filled, will executed
func case3() {
	tickerBatch := time.NewTicker(time.Second)
	go func() {
		var tempBatch []Foo

		for {
			select {
			case f := <-batchChan:
				batch = append(batch, *f)
				if len(batch) == 10 {
					tempBatch = batch
					batch = make([]Foo, 0, 10)
					batchInsert(tempBatch)
				}

			case <-tickerBatch.C:
				tempBatch = batch
				batch = make([]Foo, 0, 10)

				if len(tempBatch) > 0 {
					batchInsert(tempBatch)
				}

			case <-done:
				batchInsert(batch)
			}
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(430 * time.Millisecond)
			batchChan <- &Foo{}
		}
	}()
}
