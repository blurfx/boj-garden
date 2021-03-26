package utils

import (
	"github.com/gocolly/colly"
	"sync"
)

var crawler *colly.Collector
var once sync.Once

func GetInstance() *colly.Collector {
	once.Do(func() {
		crawler = colly.NewCollector(
			colly.Async(true),
		)
	})

	return crawler
}
