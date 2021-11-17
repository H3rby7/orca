package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ListOfGroups struct {
	Groups []Group `json:"groups"`
}

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GroupRoutes(s *mux.Router) {
	s.Path("").Methods("GET").HandlerFunc(getAllHandler)
	s.Path("").Methods("POST").HandlerFunc(createHandler)
	s.HandleFunc("/module", moduleInfoHandler)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{}")
}

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, err := json.Marshal(stubGroups())
	if err != nil {
		return
	}

	w.Write(jsonResponse)

}

func moduleInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Group module loaded")
}

func stubGroups() []Group {
	var groups []Group
	var g1 Group
	g1.Id = 1
	g1.Name = "all"
	groups = append(groups, g1)
	return groups
}
