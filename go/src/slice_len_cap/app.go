package main

import "fmt"

func main() {
	fmt.Println("Using len only")
	case1() // don't know cap, when append and out of cap, will 2x cap
	fmt.Println("Using len cap")
	case2() // know cap, less memory use, but when append and out of cap, will do the same ( 2x cap )
}

func case1() {
	slice := make([]int, 5)

	fmt.Println(len(slice), cap(slice)) // 5 10
	slice = append(slice, 1)
	fmt.Println(len(slice), cap(slice)) // 6 10
}

func case2() {
	slice := make([]int, 5, 6)

	fmt.Println(len(slice), cap(slice)) // 5 6
	slice = append(slice, 1)
	fmt.Println(len(slice), cap(slice)) // 6 6
}
