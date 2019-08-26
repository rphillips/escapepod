package crawler

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type Job struct {
}

// Crawler
type Crawler struct {
	db         *gorm.DB
	workerChan chan *Job
	client     *http.Client
}

func NewCrawler(db *gorm.DB) *Crawler {
	crawler := &Crawler{
		db:         db,
		workerChan: make(chan *Job, 50),
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
	return crawler
}

// Run intended to be run from a go routine
func (c *Crawler) Run(concurrency int, stopCh <-chan struct{}) {
	go c.Scheduler(stopCh)
	for i := 0; i < concurrency; i++ {
		go c.Worker(stopCh)
	}
	<-stopCh
}

// Scheduler intended to be run from a goroutine
func (c *Crawler) Scheduler(stopCh <-chan struct{}) {
	log.Debug("Starting scheduler")
	defer log.Debug("Stopping scheduler")
	ticker := time.NewTicker(30 * time.Minute)
	for {
		select {
		case <-stopCh:
			return
		case <-ticker.C:
			c.schedule()
		}
	}
}

func (c *Crawler) schedule() {
}

func (c *Crawler) Worker(stopCh <-chan struct{}) {
	log.Debug("Starting worker")
	select {
	case <-stopCh:
		return
	case job := <-c.workerChan:
		c.crawl(job)
	}
}

func (*Crawler) crawl(job *Job) {
	log.Debugf("Crawling %v", job)
}
