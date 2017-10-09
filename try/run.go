package try

import (
	"os"
)

// Run ...
func Run() {
	testHeader()
	mongo()
	renderTemplate()
	os.Exit(0)
}
