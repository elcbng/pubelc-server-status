package main

import (
	"gostat/config"
	"gostat/router"
	"net/http"
)

func main() {
	r := router.NewMuxer()
	r.AddRoute("GET", "/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Whoa there!"))
	})
	r.AddRoute("GET", "/v1/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	r.AddRoute("GET", "/v1/mem", func(w http.ResponseWriter, r *http.Request) {
		w.Write(utilities.catMemory())
	})
	http.ListenAndServe(":"+config.GetPort(), r)
}
