package main

import (
	"fmt"
	"time"
)

type Foo struct {
	Data string
}

var (
	batch     = make([]Foo, 0)
	batchChan = make(chan Foo)
)

func main() {
	// case1()
	// case2()
	case3()
	time.Sleep(10 * time.Second)
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
				batch = append(batch, f)

				if len(batch) >= 10 {
					tempBatch = batch
					batch = make([]Foo, 0)
					batchInsert(tempBatch)
				}
			}
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(100 * time.Millisecond)
			batchChan <- Foo{}
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
				batch = append(batch, f)
				if len(batch) >= 10 {
					tempBatch = batch
					batch = make([]Foo, 0)
					batchInsert(tempBatch)
				}

			case <-tickerBatch.C:
				tempBatch = batch
				batch = make([]Foo, 0)

				if len(batch) >= 0 {
					batchInsert(tempBatch)
				}
			}
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(430 * time.Millisecond)
			batchChan <- Foo{}
		}
	}()
}
