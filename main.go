package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/anaskhan96/soup"
	"github.com/gorilla/mux"
	"github.com/loivis/godis/utils"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	soup.Header("User-Agent", utils.UserAgent())
}

func main() {
	// try.Run()
	// books.UpdateOrigin()

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":3001", r))
}

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := mgo.Dial(utils.HostIP())
	defer session.Close()
	c := session.DB("godis").C("books")
	query := bson.M{}
	result := []book{}
	c.Find(query).All(&result)

	t := template.Must(template.ParseFiles("templates/books.html"))
	t.Execute(w, result)
}

type book struct {
	Name   string
	Site   string
	Link   string
	Author string
}
