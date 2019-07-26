package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(recipient string) string {
	if recipient == "" {
		recipient = "World"
	}
	return englishHelloPrefix + recipient
}

func main() {
	fmt.Println(Hello("Chris"))
}
