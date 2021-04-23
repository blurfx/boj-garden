package utils

import (
	"boj-garden/models"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"log"
)

func RunCrontab(db *gorm.DB) {
	c := cron.New()

	_, err := c.AddFunc("@every 2h", func() {
		var users []models.User
		db.Find(&users)

		for _, user := range users {
			crawlTask := CrawlTask{
				DB: db,
				Username: user.Username,
			}

			EnqueueCrawlTask(crawlTask)
		}
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Crontab Start")

	c.Start()
}