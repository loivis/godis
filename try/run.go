package try

import (
	"log"

	"github.com/loivis/godis/utils"
)

// Run ...
func Run() {
	testHeader()
	mongo()
	renderTemplate()
	log.Println("# print book hash as id")
	log.Println(utils.BookHash("永夜君王", "纵横中文网"))
}
