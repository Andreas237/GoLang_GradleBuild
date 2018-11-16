package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	hs := House{rooms: 2, address: "676 west monument\n"}

	fmt.Println("Rooms: ", hs.rooms)

	lst := list.New()

	fmt.Println("List size: ", lst.Len())

	// Read file: https://gobyexample.com/reading-files

} // end func main()

/*
House is a generic container, designed to do nothing
*/
type House struct {
	rooms   int
	address string
} // end type House struct

func (h House) GetRooms() int      { return h.rooms }
func (h House) GetAddress() string { return h.address }
