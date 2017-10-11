package try

import (
	"fmt"

	"github.com/loivis/godis/utils"
	"gopkg.in/mgo.v2/bson"
)

func mongo() {
	mgoQuery()
}

func mgoQuery() {
	fmt.Println("### mongodb query")
	session := utils.MongoSession()
	defer session.Close()
	c := session.DB("godis").C("sites")
	result := site{}
	err := c.Find(bson.M{}).One(&result)
	utils.CheckError(err)
	fmt.Println(result)
}
