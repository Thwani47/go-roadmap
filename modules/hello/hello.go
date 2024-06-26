package main

import (
	"fmt"
	"log"

	"github.com/Thwani47/go-roadmap/modules/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Alice")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	names := []string{"Alice", "Bob", "Cindy"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
