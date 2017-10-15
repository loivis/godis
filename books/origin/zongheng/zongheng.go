package zongheng

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/loivis/godis/structs"
	"github.com/loivis/godis/utils"
	"gopkg.in/mgo.v2/bson"
)

var siteName = "纵横中文网"

// Update ...
func Update() {
	linkAll := "http://book.zongheng.com/store.html"
	reset(linkAll)
}

func reset(link string) {
	log.Println("reset books in zongheng")
	resp, err := soup.Get(link)
	if err != nil {
		log.Println("error fetching page: ", link)
		log.Println(err)
	}

	doc := soup.HTMLParse(resp)
	count, _ := strconv.Atoi(doc.Find("div", "class", "pagenumber pagebar").Attrs()["count"])
	total, _ := strconv.Atoi(doc.Find("div", "class", "pagenumber pagebar").Attrs()["total"])
	log.Println("number of pages: ", count)
	log.Println("number of books: ", total)
	// count = 1
	for index := 0; index < count; index++ {
		pageNumber := strconv.Itoa(index + 1)
		pageLink := "http://book.zongheng.com/store/c0/c0/b0/u0/p" + pageNumber + "/v9/s9/t0/ALL.html"
		if bookList(pageLink) {
			break
		}
		// time.Sleep(time.Duration(utils.RandomInt(5, 13)) * time.Second)
	}
}

func bookList(link string) bool {
	var lastUpdateNew time.Time
	var shouldStop bool
	lastSiteUpdate := lastSiteUpdate()
	log.Println("last site update:", lastSiteUpdate)
	session := utils.MongoSession()
	defer session.Close()

	log.Println(link)
	resp, _ := soup.Get(link)
	doc := soup.HTMLParse(resp)
	books := doc.Find("ul", "class", "main_con").FindAll("li")[:]
	for _, book := range books {
		if _, ok := book.Attrs()["class"]; !ok {
			bookName := book.Find("a", "class", "fs14").Text()
			bookHash := utils.BookHash(bookName, siteName)
			bookLink := book.Find("a", "class", "fs14").Attrs()["href"]
			wordCount := utils.TrimAtoi(book.Find("span", "class", "number").Text())
			author := book.Find("span", "class", "author").Find("a").Text()
			authorLink := book.Find("span", "class", "author").Find("a").Attrs()["href"]
			category := book.Find("span", "class", "kind").Find("a").Text()
			resp, _ := soup.Get(bookLink)
			doc := soup.HTMLParse(resp)
			cover := doc.Find("div", "class", "book_cover fl").Find("img").Attrs()["src"]
			log.Println(bookName, author, category, wordCount)

			c := session.DB("godis").C("books")
			query := bson.M{"name": bookName, "site": siteName}
			change := bson.M{
				"$set": bson.M{
					"name":        bookName,
					"site":        siteName,
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
			log.Println("last book update:", updateTime)
			if updateTime.After(lastUpdateNew) {
				lastUpdateNew = updateTime
			}
			if !updateTime.After(lastSiteUpdate) {
				log.Println("reached last site update:", lastSiteUpdate)
				shouldStop = true
				break
			}
			// if i != len(books)-1 {
			// 	duration := time.Duration(utils.RandomInt(7, 13)) * time.Second
			// 	log.Printf("sleep %s to continue\n", duration)
			// 	time.Sleep(duration)
			// }
		}
	}
	c := session.DB("godis").C("sites")
	query := bson.M{"name": siteName}
	change := bson.M{"$set": bson.M{
		"name":        siteName,
		"last_update": lastUpdateNew}}
	_, err := c.Upsert(query, change)
	utils.CheckError(err)
	return shouldStop
}

func chapterList(name, link string) time.Time {
	var lastUpdate time.Time
	log.Println(link)
	resp, _ := soup.Get(link)
	doc := soup.HTMLParse(resp)
	chapters := doc.FindAll("td", "class", "chapterBean")
	chaptersM := []bson.M{}
	for _, chapter := range chapters {
		chapterName := chapter.Attrs()["chaptername"]
		updateTime := chapter.Attrs()["updatetime"][:len(chapter.Attrs()["updatetime"])-3]
		lastUpdate = time.Unix(int64(utils.TrimAtoi(updateTime)), 0)
		wordCount := utils.TrimAtoi(chapter.Attrs()["wordnum"])
		chapterLink := chapter.Find("a").Attrs()["href"]
		vip := chapter.Find("em", "class", "vip")
		// log.Println(name, chapterName, lastUpdate, chapterLink)
		chaptersM = append(chaptersM, bson.M{
			"name":        chapterName,
			"update_time": lastUpdate,
			"link":        chapterLink,
			"vip":         isVip(vip),
			"word_count":  wordCount})
	}

	session := utils.MongoSession()
	defer session.Close()
	c := session.DB("godis").C("books")
	query := bson.M{"name": name, "site": siteName}
	change := bson.M{"$set": bson.M{
		"name":        name,
		"site":        siteName,
		"last_update": lastUpdate,
		"chapters":    chaptersM}}
	_, err := c.Upsert(query, change)
	utils.CheckError(err)
	log.Println(len(chaptersM), "chapters in total")
	return lastUpdate
}

func lastSiteUpdate() time.Time {
	session := utils.MongoSession()
	defer session.Close()
	c := session.DB("godis").C("sites")
	query := bson.M{"name": siteName}
	result := structs.Site{}
	err := c.Find(query).One(&result)
	utils.CheckError(err)
	return result.LastUpdate
}

func isVip(root soup.Root) bool {
	if root.Error != nil {
		return false
	}
	return true
}
