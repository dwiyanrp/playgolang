package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type P struct {
	X, Y int
	Name string
}

func main() {
	startTime := time.Now()

	for i := 0; i < 1000000; i++ {
		p := P{3, 4, "Pythagoras"}
		str, _ := json.Marshal(p)

		json.Unmarshal(str, P{})
	}

	fmt.Println("Struct	: ", time.Now().Sub(startTime))

	startTime = time.Now()

	for i := 0; i < 1000000; i++ {
		p := &P{3, 4, "Pythagoras"}
		str, _ := json.Marshal(p)

		json.Unmarshal(str, &P{})
	}

	fmt.Println("Pointer	: ", time.Now().Sub(startTime))
}
