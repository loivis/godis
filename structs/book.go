package structs

// Book ...
type Book struct {
	Name       string
	Site       string
	Link       string
	Author     string
	AuthorLink string `bson:"author_link"`
	Hash       int
	LastUpdate int `bson:"last_update"`
	WordCount  int `bson:"word_count"`
	Chapters   []Chapter
}
