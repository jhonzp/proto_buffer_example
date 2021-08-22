///mnt/c/Users/zuni230669/go/src/github.com/jhonzp/proto_buffer_example
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	complexpb "github.com/jhonzp/proto_buffer_example/src/complex"
	enumpb "github.com/jhonzp/proto_buffer_example/src/enum_example"
	"github.com/jhonzp/proto_buffer_example/src/example"
	simplepb "github.com/jhonzp/proto_buffer_example/src/simple"
)

func main() {
	fmt.Println("Hello Go!!")
	sm := doSimple()
	readAndWriteDemo(sm)
	jsonDemo(sm)
	doEnum()
	doComplex()
	readAndWriteAddressBook(doAddressBook(), "AddressBook.bin")
}

func readAndWriteAddressBook(pb proto.Message, filename string) {
	fmt.Println("_____________________readAndWriteAddressBook example________________________")
	writeToFile(filename, pb)
	sm2 := &example.AddressBook{}
	readFromFile(filename, sm2)
	fmt.Println(*sm2)
}

func doAddressBook() *example.AddressBook {
	fmt.Println("_____________________doAddressBook example________________________")

	pncasa23 := example.Person_PhoneNumber{
		Number: "0328232323",
		Type:   example.Person_HOME,
	}

	pm := example.Person{
		Id:    1,
		Name:  "Jhon Message",
		Email: "jhonzp@gmail.com",
		Phones: []*example.Person_PhoneNumber{
			{
				Number: "3014444444",
				Type:   example.Person_WORK,
			},
			{
				Number: "03283666666",
				Type:   example.Person_HOME,
			},
			&pncasa23,
		},
	}

	am := example.AddressBook{
		People: []*example.Person{
			&pm,
			&example.Person{
				Id:    2,
				Name:  "Lucia Message",
				Email: "luropa@gmail.com",
				Phones: []*example.Person_PhoneNumber{
					&pncasa23,
					{
						Number: "310000292",
						Type:   example.Person_WORK,
					},
				},
			},
		},
	}

	fmt.Println(am)
	return &am
}

func doComplex() {
	fmt.Println("_____________________doComplex example________________________")
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			{
				Id:   2,
				Name: "Second message",
			},
			{
				Id:   3,
				Name: "Third message",
			},
		},
	}
	fmt.Println(cm)
}

func doEnum() {
	fmt.Println("_____________________doEnum example________________________")
	em := enumpb.EnumMessage{
		Id:           4321,
		DayOfTheWeek: enumpb.DayOfTheWeek_MONDAY,
	}

	fmt.Println(em)
}

func readAndWriteDemo(pb proto.Message) {
	fmt.Println("_____________________readAndWriteDemo example________________________")
	writeToFile("simple.bin", pb)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(*sm2)
}

func jsonDemo(pb proto.Message) {
	fmt.Println("_____________________jsonDemo example________________________")
	ssm := toJson(pb)
	fmt.Println("marshall to json (string)", ssm)
	sm2 := &simplepb.SimpleMessage{}
	fromJson(ssm, sm2)
	fmt.Println("umarshall from json: ")
	fmt.Println(*sm2)
}

func toJson(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to Json", err)
	}
	return out
}

func fromJson(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Can't convert to object", err)
	}
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
	fmt.Println("_____________________doSimple example________________________")
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
