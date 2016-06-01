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
	router.Path("/packages").Methods(http.MethodPost).Handler(logger(postPackage, "postPackage"))
	router.Path("/packages").Methods(http.MethodGet).Queries("lookupTerm", "{lookupTerm}").Handler(logger(getPackages, "getPackages"))
	router.Path("/packages/{packageName}").Methods(http.MethodGet).Handler(logger(getPackagesByPackageName, "getPackagesByPackageName"))
	router.Path("/packages/{packageName}/{version}").Methods(http.MethodGet).Handler(logger(getPackage, "getPackage"))

	return router
}
