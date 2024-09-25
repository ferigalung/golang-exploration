package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Feri")
	fmt.Println(message)
}
