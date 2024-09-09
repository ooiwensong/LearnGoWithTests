package main

import (
	// "bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

// using the more general Writer interface allows us to pass in both os.Stdout and bytes.Buffer
func Greet( /*  writer *bytes.Buffer  */ writer io.Writer, name string) {
	// Fprintf's first argument allows users to choose where to print the string to
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
}
