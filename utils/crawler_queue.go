package utils

import (
	"gorm.io/gorm"
	"sync"
)

type CrawlTask struct {
	DB       *gorm.DB
	Username string
}

var (
	wg      sync.WaitGroup
	channel = make(chan CrawlTask, 500)
	once    sync.Once
)

func EnqueueCrawlTask(crawlTask CrawlTask) bool {
	once.Do(func() {
		wg.Add(1)
		go worker(&wg)
	})

	select {
	case channel <- crawlTask:
		return true
	default:
		return false
	}
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range channel {
		crawler := GetCrawlerInstance()
		crawler.Crawl(task.DB, task.Username)
	}
}
