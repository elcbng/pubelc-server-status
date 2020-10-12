package main

import (
	"gostat/router"
	"net/http"
)

func main() {
	r := router.NewMuxer()
	r.AddRoute("GET", "/v1/mem", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("memstat"))
	})
	r.AddRoute("GET", "/v1/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":8080", r)
}
