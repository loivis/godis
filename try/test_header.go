package try

import "github.com/anaskhan96/soup"
import "fmt"

func testHeader() {
	fmt.Println("### print request header")
	resp, _ := soup.Get("http://httpbin.org/headers")
	fmt.Println(resp)
}
