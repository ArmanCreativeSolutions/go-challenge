package cronjobs

import (
	"github.com/go-co-op/gocron"
	"log"
	"net/http"
	"time"
)

const (
	removeSegmentUrl = "/api/v1/cron-jobs/user-segments/remove-segment"
)

func RunCronJobs() {
	s := gocron.NewScheduler(time.UTC)
	do, err := s.Every(14).Day().Do(func() {
		runRemoveSegmentRestClient()
	})
	if err != nil {
		return
	}
	s.StartBlocking()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(do)
}

func runRemoveSegmentRestClient() {
	url := "http://localhost:8080" + removeSegmentUrl
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp)
}
