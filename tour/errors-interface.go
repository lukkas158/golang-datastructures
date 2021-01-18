package main

import "fmt"

type StringError struct {
	message string
}

func (err *StringError) Error() string {
	return fmt.Sprintf("%v", err.message)
}

func run() (string, error) {
	return "", &StringError{"Error Panic"}
}

func main() {

	_, err := run()

	if err != nil {
		fmt.Println(err)
	}
}
