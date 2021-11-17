package group

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orca/env"

	"github.com/gorilla/mux"
)

type ListOfGroups struct {
	Groups []Group `json:"groups"`
}

var model *GroupModel

func Configure(s *mux.Router, env *env.Env) {
	model = &GroupModel{DB: env.DB}
	model.LoadFixtures()
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

	groups, err := model.All(r.Context())
	if err != nil {
		return
	}
	jsonResponse, err := json.Marshal(groups)
	if err != nil {
		return
	}

	w.Write(jsonResponse)

}

func moduleInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Group module loaded")
}
