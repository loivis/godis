package try

import (
	"fmt"
	"os"

	"github.com/loivis/godis/utils"
)

// Run ...
func Run() {
	fmt.Println(utils.RandomInt(30, 100))
	os.Exit(0)
}
