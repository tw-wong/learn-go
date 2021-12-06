package main

import (
	"fmt"
	"log"

	"github.com/tw-wong/learn-go/protobuf/gen"
	"google.golang.org/protobuf/proto"
)


func main() {
	// Compile all .proto files
	// ~ protoc --go_out=. ./protobuf/*.proto
	
	// Execute go
	// ~ go run main.go protobuf/*.proto

	elliot := &gen.Person{
		Name: "Alice",
		Age: 24,
		Gender: gen.Person_FEMALE,
	}

	// Encode
	data, err := proto.Marshal(elliot)
	if err != nil {
			log.Fatal("marshaling error: ", err)
	}

	fmt.Printf("Marshal data: %v\n", data)

	// Decode 
	newElliot := &gen.Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println("Print Person:")
	fmt.Printf("- Age: %d\n", newElliot.GetAge())
	fmt.Printf("- Name: %v\n", newElliot.GetName())
	fmt.Printf("- Gender: %v\n", newElliot.GetGender())
}