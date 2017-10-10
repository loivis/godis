package handlers

import (
	"html/template"
	"net/http"

	"github.com/loivis/godis/utils"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Root ...
func Root(w http.ResponseWriter, r *http.Request) {
	session, _ := mgo.Dial(utils.HostIP())
	defer session.Close()
	c := session.DB("godis").C("books")
	query := bson.M{}
	result := []book{}
	c.Find(query).All(&result)

	t := template.Must(template.ParseFiles("templates/books.html"))
	t.Execute(w, result)
}
