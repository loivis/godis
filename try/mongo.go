package try

import (
	"fmt"

	"github.com/loivis/godis/structs"

	"github.com/loivis/godis/utils"
	"gopkg.in/mgo.v2/bson"
)

func mongo() {
	query()
	key()
}

func query() {
	fmt.Println("### mongodb query")
	session := utils.MongoSession()
	defer session.Close()
	c := session.DB("godis").C("sites")
	result := structs.Site{}
	err := c.Find(bson.M{}).One(&result)
	utils.CheckError(err)
	fmt.Println(result)
}

type doc struct {
	KeyName string `bson:"key_name"`
}

func key() {
	fmt.Println("### mongodb key")
	session := utils.MongoSession()
	defer session.Close()
	c := session.DB("test").C("col")
	doc := doc{KeyName: "key-value"}
	c.Insert(doc)
}
