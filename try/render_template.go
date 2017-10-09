package try

import (
	"fmt"
	"html/template"
	"os"

	"github.com/loivis/godis/utils"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func renderTemplate() {
	fmt.Println("### render template")

	session, _ := mgo.Dial(utils.HostIP())
	defer session.Close()
	c := session.DB("godis").C("sites")
	query := bson.M{}
	result := []site{}
	c.Find(query).All(&result)

	t := template.Must(template.ParseFiles("templates/sites.html"))
	t.Execute(os.Stdout, result)
}
