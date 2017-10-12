package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/loivis/godis/structs"

	"github.com/loivis/godis/utils"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

// ViewChapters ...
func ViewChapters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["hash"])
	session := utils.MongoSession()
	c := session.DB("godis").C("books")
	query := bson.M{"hash": utils.TrimAtoi(vars["hash"])}
	result := structs.Book{}
	c.Find(query).One(&result)

	t := template.Must(template.ParseFiles("templates/view_chapters.html", "templates/header.html"))
	t.Execute(w, result.Chapters)
}
