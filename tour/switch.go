package main

import (
	"fmt"
	"runtime"
)

func main() {

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
		fmt.Println("Linux2")
	default:
		fmt.Printf("%s", os)
	}

}
