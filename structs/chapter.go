package structs

import (
	"time"
)

// Chapter ...
type Chapter struct {
	Name       string
	Link       string
	UpdateTime time.Time `bson:"update_time"`
	WordCount  int       `bson:"word_count"`
	Vip        bool
	Hash       int
}
