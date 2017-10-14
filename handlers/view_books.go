package handlers

import (
	"html/template"
	"net/http"

	"github.com/loivis/godis/structs"
	"github.com/loivis/godis/utils"
	"gopkg.in/mgo.v2/bson"
)

// ViewBooks ...
func ViewBooks(w http.ResponseWriter, r *http.Request) {
	session := utils.MongoSession()
	defer session.Close()
	c := session.DB("godis").C("books")
	query := bson.M{}
	result := []structs.Book{}
	err := c.Find(query).All(&result)
	utils.CheckError(err)

	t := template.Must(template.ParseFiles("templates/view_books.html", "templates/header.html", "templates/footer.html"))
	t.Execute(w, result)
}
