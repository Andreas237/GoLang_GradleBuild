package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	hs := House{rooms: 2, address: "676 west monument\n"}

	fmt.Println("Rooms: ", hs.rooms)

	vec := fun New() *List"

	fmt.Println("Vector size: ", vec.size())

} // end func main()

type House struct {
	rooms   int
	address string
} // end type House struct

func (h House) GetRooms() int      { return h.rooms }
func (h House) GetAddress() string { return h.address }
