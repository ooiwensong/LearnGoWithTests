package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	// type casting PlayerServer function to HandlerFunc so that it implements ServeHTTP()
	// handler := http.HandlerFunc(PlayerServer)

	// if there is a problem with the web server, an error is returned and logged
	log.Fatal(http.ListenAndServe(":8000", server))
}
