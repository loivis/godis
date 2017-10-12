package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/loivis/godis/books"
	"github.com/loivis/godis/routers"
	"github.com/loivis/godis/try"
	"github.com/loivis/godis/utils"
)

func init() {
	soup.Header("User-Agent", utils.UserAgent())
}

func main() {
	try.Run()
	books.UpdateOrigin()
	os.Exit(0)
	// books.StartCron()

	router := routers.Router()
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:3001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
