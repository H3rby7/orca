package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HealthHandler)
	GroupRoutes(r.PathPrefix("/group").Subrouter())
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}
