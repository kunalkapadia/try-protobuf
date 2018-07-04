package main

import (
	"encoding/json"
	"log"

	"fmt"

	"bytes"

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
	buf := bytes.Buffer{}

	// Write person to buffer.
	out, err := proto.Marshal(&person)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	if _, err := buf.Write(out); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Read person from buffer
	samePerson := pb.Person{}
	content := buf.Bytes()

	fmt.Println("File contents: ", content)
	if err := proto.Unmarshal(content, &samePerson); err != nil {
		log.Fatalln("Failed to parse person:", err)
	}

	marshaledPerson, _ := json.Marshal(samePerson)
	fmt.Println("person data: ", string(marshaledPerson))
}
