package structs

// Chapter ...
type Chapter struct {
	Name       string
	Link       string
	UpdateTime int `bson:"update_time"`
	Vip        bool
}
