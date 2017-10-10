package routers

import (
	"github.com/gorilla/mux"
	"github.com/loivis/godis/handlers"
)

// Router ...
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Root)
	r.HandleFunc("/content", handlers.Chapter)
	return r
}
