package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const jsonContentType = "application/json"

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

// Our PlayerServer struct containts a reference to the PlayerStore
type PlayerServer struct {
	store PlayerStore
	// embedding the http.Handler interface in our struct
	// this makes it such that PlayerServer now has all its associated methods
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux() // returns a http.ServeMux type

	// Handler function to write Status Ok response at "/league" endpoint
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	// Handler function to process "/players" endpoint
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	// since http.ServeMux implements ServeHTTP, router can be assigned to
	// "fill in" the http.Handler
	// this allows us to explicitly remove the ServeHTTP method (see below)
	p.Handler = router

	return p
}

/**

// PlayerServer has a method which is to serve a http response using data from PlayerStore it references
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Handle the request using the router's ServeHTTP method
	p.router.ServeHTTP(w, r)
}

**/

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())

	// w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
