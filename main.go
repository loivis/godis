package main

import (
	"log"
	"net/http"

	"github.com/anaskhan96/soup"
	"github.com/loivis/godis/routers"
	"github.com/loivis/godis/try"
	"github.com/loivis/godis/utils"
)

func init() {
	soup.Header("User-Agent", utils.UserAgent())
}

func main() {
	try.Run()
	// books.UpdateOrigin()
	// os.Exit(0)
	router := routers.Router()
	log.Fatal(http.ListenAndServe(":3001", router))
}
