package try

import (
	"fmt"
	"os"

	"github.com/loivis/godis/utils"
)

// Run ...
func Run() {
	testHeader()
	mongo()
	renderTemplate()
	fmt.Println("# print book hash as id")
	fmt.Println(utils.BookHash("永夜君王", "纵横中文网"))
	os.Exit(0)
}
