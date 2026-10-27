package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
// Exported named functions start with capital letter. This method can be called
// from another module
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// The := initializes and declares the variable in a single step
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message, nil
}
