///mnt/c/Users/zuni230669/go/src/github.com/jhonzp/proto_buffer_example
package main

import (
	"fmt"

	simplepb "github.com/jhonzp/proto_buffer_example/src/simple"
)

func main() {
	fmt.Println("Hello Go!!")
	doSimple()
}

func doSimple() {
	sm := simplepb.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "My Jhon Simple Message",
		SampleList: []int32{1, 2, 3, 4},
	}
	fmt.Println(sm)
	sm.Name = "Jhon Message was renamed"
	fmt.Println(sm)
	//if value is nill it will return 0
	fmt.Println("The Id is:", sm.GetId())
	fmt.Println("The Id is:", sm.Id)
}
