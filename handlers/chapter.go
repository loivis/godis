package handlers

import (
	"html/template"
	"net/http"

	"github.com/anaskhan96/soup"
)

// Chapter ...
func Chapter(w http.ResponseWriter, r *http.Request) {
	chapterLink := "http://book.zongheng.com/chapter/709687/39147304.html"
	resp, _ := soup.Get(chapterLink)
	doc := soup.HTMLParse(resp)

	result := doc.Find("div", "id", "readerFs").FindAll("p")
	var content []string
	for _, p := range result {
		content = append(content, p.Text())
	}

	t := template.Must(template.ParseFiles("templates/content.html"))
	t.Execute(w, content)
}
