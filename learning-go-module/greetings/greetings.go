package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
// this func must return string, and error. If succeed, return nil on second return val. If error occurred, return instance of "errors" package on second return val
func Hello(name string) (string, error) {
	// return error when name is empty
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
