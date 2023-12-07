package main

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

func writeTofile(fname string, pb proto.Message) {
	out, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return
	}

	if err = os.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return
	}

	fmt.Println("Data has been written")
}

func readFromFile(fname string, pb proto.Message) {
	in, err := os.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can't read file", err)
		return
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Can't Unmarshal", err)
		return
	}

	// fmt.Println(pb)
}
