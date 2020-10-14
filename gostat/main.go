package main

import (
	"gostat/config"
	"gostat/router"
	"gostat/stat"
	"net/http"
)

func main() {
	r := router.NewMuxer()
	r.AddRoute("GET", "/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Whoa there!"))
	})
	r.AddRoute("GET", "/v1/status/mem", func(w http.ResponseWriter, r *http.Request) {
		w.Write(stat.RetJSON(stat.CatMemory()))
	})
	r.AddRoute("GET", "/v1/status/cpu", func(w http.ResponseWriter, r *http.Request) {
		w.Write(stat.RetJSON(stat.CatCPU()))
	})
	r.AddRoute("GET", "/v1/status/disk", func(w http.ResponseWriter, r *http.Request) {
		w.Write(stat.RetJSON(stat.CatDisk()))
	})
	r.AddRoute("GET", "/v1/status/uptime", func(w http.ResponseWriter, r *http.Request) {
		w.Write(stat.RetJSON(stat.CatUptime()))
	})
	r.AddRoute("GET", "/v1/status/load", func(w http.ResponseWriter, r *http.Request) {
		w.Write(stat.RetJSON(stat.CatLoad()))
	})
	r.AddRoute("GET", "/v1/status/network", func(w http.ResponseWriter, r *http.Request) {
		w.Write(stat.RetJSON(stat.CatNetwork()))
	})
	r.AddRoute("GET", "/v1/status/all", func(w http.ResponseWriter, r *http.Request) {
		w.Write(stat.RetJSON(stat.CatAll()))
	})

	http.ListenAndServe(":"+config.GetPort(), r)
}
