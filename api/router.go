package api

import (
	"fmt"
	"log"
	"net/http"
	"orca/group"

	"github.com/gorilla/mux"
)

func SetUpRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", healthHandler)
	group.GroupRoutes(r.PathPrefix("/group").Subrouter())
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}
