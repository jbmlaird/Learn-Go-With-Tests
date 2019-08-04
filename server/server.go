package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type Player struct {
	Name string
	Wins int
}

type PlayerServer struct {
	Store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.Store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.getLeagueTable())
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) getLeagueTable() League {
	return p.Store.GetLeague()
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, r, player)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, name string) {
	p.Store.RecordWin(name)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request, name string) {
	score := p.Store.GetPlayerScore(name)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
