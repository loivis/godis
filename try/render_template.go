package try

import (
	"html/template"
	"log"
	"os"

	"github.com/loivis/godis/structs"
	"github.com/loivis/godis/utils"
	"gopkg.in/mgo.v2/bson"
)

func renderTemplate() {
	log.Println("### render template")

	session := utils.MongoSession()
	defer session.Close()
	c := session.DB("godis").C("sites")
	query := bson.M{}
	result := []structs.Site{}
	c.Find(query).All(&result)

	t := template.Must(template.ParseFiles("templates/sites.html"))
	t.Execute(os.Stdout, result)
}
