package api

import (
	"fmt"
	"log"
	"net/http"
	"orca/env"
	"orca/group"

	"github.com/gorilla/mux"
)

func SetUpRoutes(env *env.Env) {
	r := mux.NewRouter()
	r.HandleFunc("/", healthHandler)
	group.Configure(r.PathPrefix("/group").Subrouter(), env)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}
