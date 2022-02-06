package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Emily")
	fmt.Println(message)
}
