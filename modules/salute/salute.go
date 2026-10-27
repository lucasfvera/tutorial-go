package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Me")
	fmt.Println(message)
}
