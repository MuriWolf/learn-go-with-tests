package main

import (
	"fmt"
)

const defaultName = "world"
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = defaultName
	}
	if lang == spanish {
		return spanishHelloPrefix + name
	}
	if lang == french {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	output := fmt.Sprintln(Hello("Murillo", ""))
	fmt.Println(output)
}
