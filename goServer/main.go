package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	// 2nd argument defines permissions for opening the file
	// os.O_RDWR means we want to read & write
	// os.O_CREATE means create the file if it does not exist
	// 3rd argument sets permissions for the file; 0666 means all users can RDWR the file
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	server := NewPlayerServer(store)

	// server := NewPlayerServer(NewInMemoryPlayerStore())

	// type casting PlayerServer function to HandlerFunc so that it implements ServeHTTP()
	// handler := http.HandlerFunc(PlayerServer)

	// if there is a problem with the web server, an error is returned and logged
	if err := http.ListenAndServe(":8000", server); err != nil {
		log.Fatalf("could not listen on port 8000 %v", err)
	}
}
