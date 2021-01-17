package main

import (
	"fmt"
	"log"
	"lukkas158/greetings"
)

func main() {

	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	messages, err := greetings.Hellos([]string{"Lucas", "Roberta", "Letícia"})

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println(messages)
}
