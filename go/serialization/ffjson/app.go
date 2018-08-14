package main

import (
	"fmt"
	"time"

	"github.com/pquerna/ffjson/ffjson"
)

type P struct {
	X    int    `json:"X"`
	Y    int    `json:"Y"`
	Name string `json:"Name"`
}

func main() {
	startTime := time.Now()
	var q P

	for i := 0; i < 500000; i++ {
		p := &P{3, 4, "Pythagoras"}
		str, _ := ffjson.Marshal(p)
		ffjson.Unmarshal(str, &q)
	}

	fmt.Println(q)
	fmt.Println("Struct	: ", time.Now().Sub(startTime))
}
