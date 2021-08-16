package main

import (
	"fmt"

	example_simple "github.com/jhonzp/proto_buffer_example/src/simple"
)

func main() {
	fmt.Println("Hello Go!!")
	doSimple()
}

func doSimple() {
	sm := example_simple.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "My Jhon Simple Message",
		SampleList: []int32{1, 2, 3, 4},
	}
	fmt.Println(sm)
}
