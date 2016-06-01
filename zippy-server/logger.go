package main

import (
	"log"
	"net/http"
	"time"
)

func logger(inner func(http.ResponseWriter, *http.Request), name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner(w, r)
		log.Printf("%s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start))
	})
}
