package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	sayHelloWorld("Hello, sayHelloWorld!")
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
