package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

type Foo struct {
	Value int64
}

func main() {
	c := cache.New(5*time.Minute, 10*time.Minute)

	case1(c)
	case2(c)
	case3(c)
	case4(c)
	case5(c)
}

// Update Data from gocache pointer
// unsafe
func case1(c *cache.Cache) {
	c.SetDefault("case1", &Foo{Value: 1})

	if first, found := c.Get("case1"); found {
		data := first.(*Foo)
		data.Value = 999
	}

	if second, found := c.Get("case1"); found {
		fmt.Println(second)
	}
}

// Update Data from gocache struct
// safe
func case2(c *cache.Cache) {
	c.SetDefault("case2", Foo{Value: 1})

	if first, found := c.Get("case2"); found {
		data := first.(Foo)
		data.Value = 999
	}

	if second, found := c.Get("case2"); found {
		fmt.Println(second)
	}
}

// Update Data from gocache slice
// unsafe
func case3(c *cache.Cache) {
	sl := []Foo{Foo{Value: 1}}
	c.SetDefault("case3", sl)

	if first, found := c.Get("case3"); found {
		data := first.([]Foo)
		data[0].Value = 999
	}

	if second, found := c.Get("case3"); found {
		fmt.Println(second)
	}
}

// Update Data from gocache slice copy
// safe
func case4(c *cache.Cache) {
	sl := []Foo{Foo{Value: 1}}
	c.SetDefault("case4", sl)

	if first, found := c.Get("case4"); found {
		raw := first.([]Foo)
		data := make([]Foo, len(raw))
		copy(data, raw)
		data[0].Value = 999
	}

	if second, found := c.Get("case4"); found {
		fmt.Println(second)
	}
}

// Update Data from gocache slice of pointer copy
// unsafe
func case5(c *cache.Cache) {
	sl := []*Foo{&Foo{Value: 1}}
	c.SetDefault("case5", sl)

	if first, found := c.Get("case5"); found {
		raw := first.([]*Foo)
		data := make([]*Foo, len(raw))
		copy(data, raw)
		data[0].Value = 999
	}

	if second, found := c.Get("case5"); found {
		data := second.([]*Foo)
		fmt.Println(data[0].Value)
	}
}
