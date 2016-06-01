package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := createZippyRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createZippyRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/packages", postPackage).Methods(http.MethodPost)
	router.HandleFunc("/packages", getPackages).Methods(http.MethodGet).Queries("lookupTerm", "{lookupTerm}")
	router.HandleFunc("/packages/{packageName}", getPackagesByPackageName).Methods(http.MethodGet)
	router.HandleFunc("/packages/{packageName}/{version}", getPackage).Methods(http.MethodGet)

	return router
}
