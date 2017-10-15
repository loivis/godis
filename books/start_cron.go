package books

import (
	"log"

	"github.com/robfig/cron"
)

// StartCron ...
func StartCron() {
	c := cron.New()
	c.AddFunc("@every 5m", UpdateOrigin)
	log.Println("starting cron job")
	c.Start()
}
