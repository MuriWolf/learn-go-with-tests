package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, text string) {
	fmt.Fprintf(writer, "Bonjour, %s", text)
}

func myGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "monde")
}

func main() {
	fmt.Println("Server running at http://localhost:5001")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(myGreeterHandler)))
}
