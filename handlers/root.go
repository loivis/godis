package handlers

import (
	"html/template"
	"net/http"

	"github.com/loivis/godis/structs"
	"github.com/loivis/godis/utils"
	"gopkg.in/mgo.v2/bson"
)

// Root ...
func Root(w http.ResponseWriter, r *http.Request) {
	session := utils.MongoSession()
	defer session.Close()
	c := session.DB("godis").C("books")
	query := bson.M{}
	result := []structs.Book{}
	c.Find(query).All(&result)

	t := template.Must(template.ParseFiles("templates/books.html"))
	t.Execute(w, result)
}
