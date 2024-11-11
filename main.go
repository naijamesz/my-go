package main

import (
	"fmt"
)

func main() {
	// normal way to variables and assignment
	// var whatToSay string
	// whatToSay = `Hello World again and again!`
	// sayHelloWorld(whatToSay) // OUTPUT : Hello World again and again!

	// shorthand easy way auto type by value
	whatToSay := `Hello World again by shorthand`
	sayHelloWorld(whatToSay) // OUTPUT : Hello World again by shorthand

	sayHelloWorld(`Hello World again!👋🏻`)      // OUTPUT : Hello World again!👋🏻
	sayHelloWorld(`Bonjour, tout le monde!👋🏻`) // OUTPUT : Bonjour, tout le monde!👋🏻
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
