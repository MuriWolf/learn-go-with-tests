package main

import (
	"fmt"
)

const (
	defaultName        = "world"
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = defaultName
	}

	return getGreetingPrefix(lang) + name
}

func getGreetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	output := fmt.Sprintln(Hello("Murillo", ""))
	fmt.Println(output)
}
