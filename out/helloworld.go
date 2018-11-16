package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
)

// Panic if input isn't empty
// Example usage: check error when file opened
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Hello World")

	hs := House{rooms: 2, address: "676 west monument\n"}

	fmt.Println("Rooms: ", hs.rooms)

	my_list := list.New()
	fmt.Println("Vector size: ", my_list.Len())

	/*
	  Read a file, save the contents to a vector, sort them
	*/
	f, err := ioutil.ReadFile("../input_files/test_ints.txt")
	check(err)

	fmt.Println("File Contents: %s", f)

} // end func main()

type House struct {
	rooms   int
	address string
} // end type House struct

func (h House) GetRooms() int      { return h.rooms }
func (h House) GetAddress() string { return h.address }
