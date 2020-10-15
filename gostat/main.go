package main

import (
	"fmt"
	"gostat/config"
	"gostat/router"
	"gostat/stat"
	"net/http"
)

func main() {
	r := router.NewMuxer()
	URL := config.GetURL()
	r.AddRoute("GET", URL+"/raw"+"/mem", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write(stat.Catter("mem")); err != nil {
			fmt.Println("Http write failed:", err)
		}
	})
	r.AddRoute("GET", URL+"/raw"+"/cpu", func(w http.ResponseWriter, r *http.Request) {

		if _, err := w.Write(stat.Catter("cpu")); err != nil {
			fmt.Println("Http write failed:", err)
		}
	})
	r.AddRoute("GET", URL+"/raw"+"/disk", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write(stat.Catter("disk")); err != nil {
			fmt.Println("Http write failed:", err)
		}
	})
	r.AddRoute("GET", URL+"/raw"+"/uptime", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write(stat.Catter("uptime")); err != nil {
			fmt.Println("Http write failed:", err)
		}
	})
	r.AddRoute("GET", URL+"/raw"+"/load", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write(stat.Catter("load")); err != nil {
			fmt.Println("Http write failed:", err)
		}
	})
	r.AddRoute("GET", URL+"/raw"+"/network", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write(stat.Catter("network")); err != nil {
			fmt.Println("Http write failed:", err)
		}
	})
	r.AddRoute("GET", URL+"/raw"+"/all", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write(stat.CatAll()); err != nil {
			fmt.Println("Http write failed:", err)
		}
	})

	if err := http.ListenAndServe(":"+config.GetPort(), r); err != nil {
		fmt.Println("open HTTP server error")
	}
}
