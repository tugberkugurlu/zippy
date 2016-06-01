package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/packages", postPackage).Methods("POST")
	router.HandleFunc("/packages", getPackages).Methods("GET").Queries("lookupTerm", "{lookupTerm}")
	router.HandleFunc("/packages/{packageName}", getPackagesByPackageName).Methods("GET")
	router.HandleFunc("/packages/{packageName}/{version}", getPackage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getPackages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lookupTerm := vars["lookupTerm"]
	fmt.Fprintf(w, "Hello from getPackages. path: %q, lookupTerm: %q", html.EscapeString(r.URL.Path), lookupTerm)
}

func getPackagesByPackageName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	packageName := vars["packageName"]
	fmt.Fprintf(w, "Hello from getPackagesByPackageName. path: %q, packageName: %q", html.EscapeString(r.URL.Path), packageName)
}

func getPackage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	packageName := vars["packageName"]
	version := vars["version"]
	fmt.Fprintf(w, "Hello from getPackage. path: %q, packageName: %q, version: %q", html.EscapeString(r.URL.Path), packageName, version)
}

func postPackage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from postPackage, %q", html.EscapeString(r.URL.Path))
}
