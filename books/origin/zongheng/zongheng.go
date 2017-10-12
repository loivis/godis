package zongheng

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/loivis/godis/structs"
	"github.com/loivis/godis/utils"
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
		time.Sleep(time.Duration(utils.RandomInt(7, 13)) * time.Second)
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
			bookHash := utils.BookHash(bookName, "纵横中文网")
			bookLink := book.Find("a", "class", "fs14").Attrs()["href"]
			wordCount := utils.TrimAtoi(book.Find("span", "class", "number").Text())
			author := book.Find("span", "class", "author").Find("a").Text()
			authorLink := book.Find("span", "class", "author").Find("a").Attrs()["href"]
			category := book.Find("span", "class", "kind").Find("a").Text()
			resp, _ := soup.Get(bookLink)
			doc := soup.HTMLParse(resp)
			cover := doc.Find("div", "class", "book_cover fl").Find("img").Attrs()["src"]
			fmt.Println(bookName, bookLink, author, authorLink, category, wordCount, cover)

			session := utils.MongoSession()
			defer session.Close()
			c := session.DB("godis").C("books")
			query := bson.M{"name": bookName, "site": "纵横中文网"}
			change := bson.M{
				"$set": bson.M{
					"name":        bookName,
					"site":        "纵横中文网",
					"hash":        bookHash,
					"link":        bookLink,
					"cover":       cover,
					"author":      author,
					"author_link": authorLink,
					"category":    category,
					"word_count":  wordCount}}
			_, err := c.Upsert(query, change)
			utils.CheckError(err)
			chapterLink := strings.Replace(bookLink, "/book/", "/showchapter/", -1)
			updateTime := chapterList(bookName, chapterLink)
			fmt.Println(updateTime, lastUpdate())
			time.Sleep(time.Duration(utils.RandomInt(7, 13)) * time.Second)
		}
	}
}

func chapterList(name, link string) int {
	var updateTime int
	fmt.Println(link)
	resp, _ := soup.Get(link)
	doc := soup.HTMLParse(resp)
	chapters := doc.FindAll("td", "class", "chapterBean")
	chaptersD := []bson.M{}
	for _, chapter := range chapters {
		chapterName := normalizeKey(chapter.Attrs()["chaptername"])
		updateTime = utils.TrimAtoi(chapter.Attrs()["updatetime"]) / 1e3
		chapterLink := chapter.Find("a").Attrs()["href"]
		vip := chapter.Find("em", "class", "vip")
		fmt.Println(name, chapterName, updateTime, chapterLink)
		chaptersD = append(chaptersD, bson.M{
			"name":        chapterName,
			"update_time": updateTime,
			"link":        chapterLink,
			"vip":         isVip(vip)})
	}

	session := utils.MongoSession()
	defer session.Close()
	c := session.DB("godis").C("books")
	query := bson.M{"name": name, "site": "纵横中文网"}
	change := bson.M{"$set": bson.M{
		"name":        name,
		"site":        "纵横中文网",
		"last_update": updateTime,
		"chapters":    chaptersD}}
	_, err := c.Upsert(query, change)
	utils.CheckError(err)
	return updateTime
}

func lastUpdate() int {
	session := utils.MongoSession()
	defer session.Close()
	c := session.DB("godis").C("sites")
	query := bson.M{"name": "纵横中文网"}
	result := structs.Site{}
	err := c.Find(query).One(&result)
	utils.CheckError(err)
	return result.Update
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
