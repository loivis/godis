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
	// static resources
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("static/img"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	return r
}
