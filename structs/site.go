package structs

import (
	"time"
)

// Site ...
type Site struct {
	Name       string
	Home       string
	Update     int
	Copyright  bool
	LastUpdate time.Time `bson:"last_update"`
}
