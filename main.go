///mnt/c/Users/zuni230669/go/src/github.com/jhonzp/proto_buffer_example
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	simplepb "github.com/jhonzp/proto_buffer_example/src/simple"
)

func main() {
	fmt.Println("Hello Go!!")
	sm := doSimple()
	readAndWriteDemo(sm)
}

func readAndWriteDemo(pb proto.Message) {
	writeToFile("simple.bin", pb)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(*sm2)
}

func writeToFile(filename string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}
	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}
	fmt.Println("Data has been read!")
	return nil
}

func doSimple() *simplepb.SimpleMessage {
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
	return &sm
}
