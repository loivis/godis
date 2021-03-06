package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/loivis/qieshu/structs"

	"github.com/loivis/qieshu/utils"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

// ViewContent ...
func ViewContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("book id:", vars["hash"])
	session := utils.MongoSession()
	c := session.DB("godis").C("books")
	query := bson.M{"hash": utils.TrimAtoi(vars["hash"])}
	result := structs.Book{}
	c.Find(query).One(&result)

	t := template.Must(template.ParseFiles("templates/view_content.html", "templates/header.html", "templates/footer.html"))
	t.Execute(w, result)
}
