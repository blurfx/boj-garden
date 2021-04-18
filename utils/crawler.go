package utils

import (
	"boj-garden/models"
	"fmt"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
	"math"
	"strconv"
	"time"
)

type Crawler struct {
	Crawler *colly.Collector
}

func GetCrawlerInstance() *Crawler {

	crawler := &Crawler{
		Crawler: colly.NewCollector(
			colly.Async(true),
			colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36"),
		),
	}

	duration, _ := time.ParseDuration("1m")

	crawler.Crawler.SetRequestTimeout(duration)

	return crawler
}

func (c *Crawler) Crawl(db *gorm.DB, username string) {

	var (
		lastSubmission  *models.Submission
		firstSubmission *models.Submission
		user            *models.User
	)

	db.FirstOrCreate(&user, &models.User{Username: username})

	if err := db.Where("user_id = ?", user.ID).First(&firstSubmission).Error; err != nil {
		firstSubmission.ID = math.MaxInt32
	}

	if err := db.Where("user_id = ?", user.ID).Last(&lastSubmission).Error; err != nil {
		lastSubmission.ID = 0
	}

	results := map[string]string {
		"result-ac": "AC",
		"result-pac": "PAC",
	}

	_ = c.Crawler.Visit("https://www.acmicpc.net/status?&result_id=4&user_id=" + user.Username)

	c.Crawler.OnRequest(func(request *colly.Request) {
		fmt.Println(request.URL)
	})

	c.Crawler.OnHTML("table", func(el *colly.HTMLElement) {
		var submissions []models.Submission

		el.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			var result string

			tr.ForEach(".result-text span", func(_ int, span *colly.HTMLElement) {
				for class, value := range results {
					if span.DOM.HasClass(class) {
						result = value
					}
				}
			})

			submissionId64, _ := strconv.ParseUint(tr.ChildText("td:first-of-type"), 10, 32)
			submissionId := uint(submissionId64)
			problemId64, _ := strconv.ParseUint(tr.ChildText("td:nth-of-type(3)"), 10, 32)
			problemId := uint(problemId64)
			language := tr.ChildText("td:nth-of-type(7)")
			timestamp, _ := strconv.ParseInt(tr.ChildAttr("a[data-timestamp]", "data-timestamp"), 10, 64)
			date := time.Unix(timestamp, 0)

			if result != "" && (submissionId > lastSubmission.ID || submissionId < firstSubmission.ID) {
				submissions = append(submissions, models.Submission{
					ID: submissionId,
					UserID: user.ID,
					ProblemID: problemId,
					Language: language,
					Result: result,
					Date: date,
				})
			}
		})

		if len(submissions) > 0 {
			db.Create(&submissions)
		}
	})

	c.Crawler.OnHTML("#next_page", func(el *colly.HTMLElement) {
		err := el.Request.Visit(el.Attr("href"))
		if err != nil {
			fmt.Println(err)
		}
	})

	c.Crawler.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})

	c.Crawler.Wait()
}
