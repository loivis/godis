package utils

import (
	"gopkg.in/mgo.v2"
)

// MongoSession ...
func MongoSession() *mgo.Session {
	mongoHost := HostIP()
	session, err := mgo.Dial(mongoHost)
	CheckError(err)
	return session
}
