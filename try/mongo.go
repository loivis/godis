package try

import (
	"log"

	"github.com/loivis/godis/structs"

	"github.com/loivis/godis/utils"
	"gopkg.in/mgo.v2/bson"
)

func mongo() {
	query()
	// insert()
	updateOne()
}

func query() {
	log.Println("### mongodb query")
	session := utils.MongoSession()
	defer session.Close()
	c := session.DB("godis").C("sites")
	result := structs.Site{}
	err := c.Find(bson.M{}).One(&result)
	utils.CheckError(err)
	log.Println(result)
}

type doc struct {
	KeyName string `bson:"key_name"`
	Fruits  []fruit
}

type fruit struct {
	Name  string
	Price int
}

func insert() {
	log.Println("### mongodb key")
	session := utils.MongoSession()
	defer session.Close()
	c := session.DB("test").C("col")
	doc := doc{KeyName: "some-value", Fruits: []fruit{fruit{Name: "apple", Price: 1}, fruit{Name: "banana", Price: 1}}}
	c.Insert(doc)
}

func updateOne() {
	log.Println("### update document")
	session := utils.MongoSession()
	defer session.Close()
	c := session.DB("test").C("col")
	result := doc{}
	query := bson.M{"_id": bson.ObjectIdHex("59dfd3865a374ac2a32f4c43"), "fruits.name": "apple"}
	c.Find(query).One(&result)
	log.Println(result)
	change := bson.M{"$set": bson.M{"fruits.$.price": 111}}
	c.Upsert(query, change)
}
