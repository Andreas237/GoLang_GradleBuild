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

	myt := makeTransition(1, "a", "b", "", 2)
	printTransition(myt)
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

/*
Transition for Push Down Automaton
*/
type Transition struct {
	currentState int
	input        string
	stackPush    string
	stackPop     string
	finalState   int
} // end type Transition struct

func makeTransition(currentState int,
	input string,
	stackPush string,
	stackPop string,
	finalState int) Transition {
	tr := Transition{currentState: currentState, input: input, stackPush: stackPush, stackPop: stackPop, finalState: finalState}
	return tr
} // end makeTransition

func printTransition(t Transition) {
	fmt.Printf("Current State: %d\nInput: %s\nStack Push: %s\nStack Pop: %s\nFinal State: %d\n", t.currentState, t.input, t.stackPop, t.stackPush, t.finalState)

} // end printTransition
