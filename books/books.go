package books

import (
	"github.com/anaskhan96/soup"
	"github.com/loivis/godis/books/origin"
	"github.com/loivis/godis/utils"
)

// UpdateOrigin ...
func UpdateOrigin() {
	soup.Header("User-Agent", utils.UserAgent())
	origin.Update()
}
