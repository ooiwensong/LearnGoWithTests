package main

import (
	"log"
	"net/http"

	poker "github.com/ooiwensong/LearnGoWithTests/goServer"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)

	// server := NewPlayerServer(NewInMemoryPlayerStore())

	// type casting PlayerServer function to HandlerFunc so that it implements ServeHTTP()
	// handler := http.HandlerFunc(PlayerServer)

	if err := http.ListenAndServe(":8000", server); err != nil {
		// if there is a problem with the web server, an error is returned and logged
		log.Fatalf("could not listen on port 8000 %v", err)
	}
}
