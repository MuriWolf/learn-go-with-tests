package main

import (
	"fmt"
)

const defaultName = "world"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, lang string) string {
	if name == "" {
		name = defaultName
	}
	if lang == spanish {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	output := fmt.Sprintln(Hello("Murillo", ""))
	fmt.Println(output)
}
