package try

import (
	"fmt"

	"github.com/loivis/godis/utils"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type site struct {
	Name      string
	Home      string
	Copyright bool
}

func mongo() {
	mgoQuery()
}

func mgoQuery() {
	fmt.Println("### mongodb query")
	session, err := mgo.Dial(utils.HostIP())
	utils.CheckError(err)
	defer session.Close()
	c := session.DB("godis").C("sites")
	result := site{}
	err = c.Find(bson.M{}).One(&result)
	utils.CheckError(err)
	fmt.Println(result)
}
