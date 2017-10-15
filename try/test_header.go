package try

import (
	"log"

	"github.com/anaskhan96/soup"
)

func testHeader() {
	log.Println("### print request header")
	resp, _ := soup.Get("http://httpbin.org/headers")
	log.Println(resp)
}
