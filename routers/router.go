package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/loivis/godis/handlers"
)

// Router ...
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Root).Methods("GET")
	r.HandleFunc("/books", handlers.ViewBooks).Methods("GET")
	r.HandleFunc("/books/{hash:[0-9]+}", handlers.ViewChapters).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return r
}
