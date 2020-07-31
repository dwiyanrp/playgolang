package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type P struct {
	X, Y int
	Name string
}

type Q struct {
	X, Y int32
	Name string
}

func main() {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	startTime := time.Now()

	for i := 0; i < 1000000; i++ {
		err := enc.Encode(P{3, 4, "Pythagoras"})
		if err != nil {
			log.Fatal("encode error:", err)
		}

		var q Q
		err = dec.Decode(&q)
		if err != nil {
			log.Fatal("decode error:", err)
		}
	}

	fmt.Println(time.Now().Sub(startTime))
}
