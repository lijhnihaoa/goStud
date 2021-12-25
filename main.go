package main

import (
	"fmt"

	"study.io/go/socket"
)

func swapIntValue(a int, b int) {
	fmt.Println("a:", a, " b:", b)
	a, b = b, a
	fmt.Println("a:", a, " b:", b)
}
func example01() {
	var slice = []string{"1", "2"}
	fmt.Printf("%p, %v, %v\n", slice, len(slice), cap(slice))
	fmt.Printf("%p, %v, %v\n", &slice, len(slice), cap(slice))
	slice = append(slice, "3")
	fmt.Printf("%p, %v, %v\n", slice, len(slice), cap(slice))
	s := add(slice)
	fmt.Printf("add(slice) : %p, %v, %v\n", slice, len(slice), cap(slice))
	fmt.Printf("s := add(slice) : %p, %v, %v\n", s, len(s), cap(s))

}

func add(slice []string) []string {
	fmt.Printf("add start: %p, %v, %v\n", slice, len(slice), cap(slice))
	slice = append(slice, "4")
	slice[0] = "0"
	// slice = append(slice, "5")
	// slice = append(slice, "6")
	fmt.Printf("add end: %p, %v, %v\n", slice, len(slice), cap(slice))
	return slice
}
func main() {
	// example01()
	// swapIntValue(1, 10)
	//example.Studhttp()

	socket.StartServer()
}
