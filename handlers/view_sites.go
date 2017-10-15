package handlers

import (
	"html/template"
	"net/http"

	"github.com/loivis/godis/structs"
	"github.com/loivis/godis/utils"
	"gopkg.in/mgo.v2/bson"
)

// ViewSites ...
func ViewSites(w http.ResponseWriter, r *http.Request) {
	session := utils.MongoSession()
	defer session.Close()
	c := session.DB("godis").C("sites")
	query := bson.M{}
	result := []structs.Site{}
	c.Find(query).All(&result)

	t := template.Must(template.ParseFiles("templates/view_sites.html", "templates/header.html", "templates/footer.html"))
	t.Execute(w, result)
}
