package books

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

// StartCron ...
func StartCron() {
	c := cron.New()
	c.AddFunc("@every 5s", func() {
		time.Sleep(time.Duration(10) * time.Second)
		fmt.Println(time.Now())
	})
	c.Start()
}
