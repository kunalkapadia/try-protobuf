package main

import (
	"io/ioutil"
	"log"

	"fmt"

	"github.com/gin-gonic/gin/json"
	"github.com/golang/protobuf/proto"
	"github.com/kunalkapadia/try-protobuf/pb"
)

func main() {
	person := pb.Person{
		Id:    1,
		Name:  "Gopher",
		Email: "gopher@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "111-2222", Type: pb.Person_MOBILE},
		},
	}
	filename := "temp.txt"

	// Write the person to disk.
	out, err := proto.Marshal(&person)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Read person from disk
	samePerson := pb.Person{}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("File contents: ", content)
	if err := proto.Unmarshal(content, &samePerson); err != nil {
		log.Fatalln("Failed to parse person:", err)
	}

	marshaledPerson, _ := json.Marshal(samePerson)
	log.Println("person data: ", string(marshaledPerson))
}
