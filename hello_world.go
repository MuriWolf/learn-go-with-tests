package main

import (
	"fmt"
)

func Hello(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	output := fmt.Sprintln(Hello("Murillo"))
	fmt.Println(output)
}
