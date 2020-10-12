package route

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Mux struct {
	handlers map[string][]*Handler
}

type Handler struct {
	path string
	http.HandlerFunc
}

func (m *Mux) Listen(port string) {
	err := http.ListenAndServe(port, m)
	if err != nil {
		log.Fatal("open HTTP server error")
	}
}

func NewMuxer() *Mux {
	return &Mux{make(map[string][]*Handler)}
}

func (m *Mux) AddRoute(mode string, path string, fun http.HandlerFunc) {
	m.add(mode, path, fun)
}

func (m *Mux) add(mode string, path string, fun http.HandlerFunc) {
	h := &Handler{strings.ToLower(path), fun}
	//debug output
	fmt.Println("h ::", &h)
	fmt.Println("strings.ToLower(path) ::", strings.ToLower(path))
	fmt.Println("strings.ToLower(mode) ::", strings.ToLower(mode))
	//core
	m.handlers[strings.ToLower(mode)] = append(
		m.handlers[strings.ToLower(mode)],
		h,
	)
	fmt.Println("m.handlers", m.handlers[strings.ToLower(mode)])
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(strings.TrimLeft(r.URL.Path, "/"), "/")
	//debug output
	fmt.Println("r", fmt.Sprintf("%+v", r))
	fmt.Println("w", fmt.Sprintf("%+v", w))

	for _, handler := range m.handlers[strings.ToLower(r.Method)] {
		if handler.path == "/"+strings.ToLower(url[0]) {
			handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
	return
}
