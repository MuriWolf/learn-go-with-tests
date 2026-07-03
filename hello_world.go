package main

import (
	"fmt"
)

func Hello(name, lang string) string {
	if name == "" {
		name = "world"
	}
	if lang == "Spanish" {
		return "Hola, " + name
	}

	return "Hello, " + name
}

func main() {
	output := fmt.Sprintln(Hello("Murillo", ""))
	fmt.Println(output)
}
