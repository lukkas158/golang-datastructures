package main

import (
	"fmt"
	"log"
	"lukkas158/greetings"
)

func main() {

	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")

	if err != nil {

		log.Fatal(err)

	}
	fmt.Println(message)
}
