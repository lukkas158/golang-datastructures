package main

import (
	"fmt"
	"lukkas158/greetings"
)

func main() {

	message := greetings.Hello("lucas")

	fmt.Println(message)
}
