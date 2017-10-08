package main

import (
	"github.com/loivis/godis/books"
	"github.com/loivis/godis/try"
)

func main() {
	try.Run()
	books.UpdateOrigin()
}
