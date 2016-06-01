package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getPackages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lookupTerm := vars["lookupTerm"]

	pkgs := packages{
		packageModel{Name: "Foo", Version: "1.0.0", CreatedOn: time.Now().Add(-10)},
		packageModel{Name: "Bar", Version: "1.0.0-beta.1", CreatedOn: time.Now().Add(-200000)},
		packageModel{Name: "FooBar", Version: "2.25.2", CreatedOn: time.Now().Add(-3000000000)},
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
