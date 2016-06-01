package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type packageModel struct {
	Name      string
	Version   string
	CreatedOn time.Time
}

type packages []packageModel

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/packages", postPackage).Methods(http.MethodPost)
	router.HandleFunc("/packages", getPackages).Methods(http.MethodGet).Queries("lookupTerm", "{lookupTerm}")
	router.HandleFunc("/packages/{packageName}", getPackagesByPackageName).Methods(http.MethodGet)
	router.HandleFunc("/packages/{packageName}/{version}", getPackage).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getPackages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lookupTerm := vars["lookupTerm"]

	pkgs := packages{
		packageModel{Name: "Foo", Version: "1.0.0"},
		packageModel{Name: "Bar", Version: "1.0.0-beta.1"},
		packageModel{Name: "FooBar", Version: "2.25.2"},
	}

	fmt.Printf("Hello from getPackages. path: %q, lookupTerm: %q", html.EscapeString(r.URL.Path), lookupTerm)

	json.NewEncoder(w).Encode(pkgs)
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
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Hello from postPackage, %q", html.EscapeString(r.URL.Path))
}
