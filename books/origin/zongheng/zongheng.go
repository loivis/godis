package zongheng

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/loivis/godis/utils"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Update ...
func Update() {
	linkAll := "http://book.zongheng.com/store.html"
	reset(linkAll)
}

func reset(link string) {
	fmt.Println("reset books in zongheng")
	resp, err := soup.Get(link)
	if err != nil {
		fmt.Println("error fetching page: ", link)
		fmt.Println(err)
	}
	// fmt.Println(resp)

	doc := soup.HTMLParse(resp)
	count, _ := strconv.Atoi(doc.Find("div", "class", "pagenumber pagebar").Attrs()["count"])
	total, _ := strconv.Atoi(doc.Find("div", "class", "pagenumber pagebar").Attrs()["total"])
	fmt.Println("number of pages: ", count)
	fmt.Println("number of books: ", total)
	count = 1
	for index := 0; index < count; index++ {
		pageNumber := strconv.Itoa(index + 1)
		pageLink := "http://book.zongheng.com/store/c0/c0/b0/u0/p" + pageNumber + "/v9/s9/t0/ALL.html"
		bookList(pageLink)
		time.Sleep(time.Duration(utils.RandomInt(7, 13)) * 1000 * time.Microsecond)
	}
}

func bookList(link string) {
	fmt.Println(link)
	resp, _ := soup.Get(link)
	doc := soup.HTMLParse(resp)
	books := doc.Find("ul", "class", "main_con").FindAll("li")[:1]
	for _, book := range books {
		if _, ok := book.Attrs()["class"]; !ok {
			bookName := book.Find("a", "class", "fs14").Text()
			bookLink := book.Find("a", "class", "fs14").Attrs()["href"]
			wordCount := utils.TrimAtoi(book.Find("span", "class", "number").Text())
			author := book.Find("span", "class", "author").Find("a").Text()
			authorLink := book.Find("span", "class", "author").Find("a").Attrs()["href"]
			category := book.Find("span", "class", "kind").Find("a").Text()
			resp, _ := soup.Get(bookLink)
			doc := soup.HTMLParse(resp)
			cover := doc.Find("div", "class", "book_cover fl").Find("img").Attrs()["src"]
			fmt.Println(bookName, bookLink, author, authorLink, category, wordCount, cover)

			mongoHost := utils.HostIP()
			session, err := mgo.Dial(mongoHost)
			utils.CheckError(err)
			defer session.Close()
			c := session.DB("godis").C("books")
			query := bson.M{"name": bookName, "site": "纵横中文网"}
			change := bson.M{
				"$set": bson.M{
					"name":        bookName,
					"link":        bookLink,
					"cover":       cover,
					"author":      author,
					"author_link": authorLink,
					"category":    category,
					"word_count":  wordCount,
					"site":        "纵横中文网"}}
			_, err = c.Upsert(query, change)
			utils.CheckError(err)
			chapterLink := strings.Replace(bookLink, "/book/", "/showchapter/", -1)
			chapterList(chapterLink)
			time.Sleep(time.Duration(utils.RandomInt(7, 13)) * time.Second)
		}
	}
}

func chapterList(link string) {
	fmt.Println(link)
	resp, _ := soup.Get(link)
	doc := soup.HTMLParse(resp)
	bookName := doc.Find("div", "class", "tc txt").Find("h1").Text()
	chapters := doc.FindAll("td", "class", "chapterBean")
	chaptersD := bson.D{}
	for i, chapter := range chapters {
		chapterName := normalizeKey(chapter.Attrs()["chaptername"])
		updateTime := chapter.Attrs()["updatetime"]
		chapterLink := chapter.Find("a").Attrs()["href"]
		vip := chapter.Find("em", "class", "vip")
		fmt.Println(bookName, chapterName, updateTime, chapterLink)
		chaptersD = append(chaptersD, bson.DocElem{
			Name: strconv.Itoa(i + 1),
			Value: bson.M{
				"name":        chapterName,
				"update_time": updateTime,
				"link":        chapterLink,
				"vip":         isVip(vip)}})
	}
	mongoHost := utils.HostIP()
	session, err := mgo.Dial(mongoHost)
	utils.CheckError(err)
	defer session.Close()
	c := session.DB("godis").C("books")
	query := bson.M{"name": bookName, "site": "纵横中文网"}
	change := bson.M{"$set": bson.M{
		"name":     bookName,
		"site":     "纵横中文网",
		"chapters": chaptersD}}
	_, err = c.Upsert(query, change)
	utils.CheckError(err)
}

func normalizeKey(name string) string {
	return strings.Replace(name, ".", "_", -1)
}

func isVip(root soup.Root) bool {
	if root.Error != nil {
		return false
	}
	return true
}
