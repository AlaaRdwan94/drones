package api_worker

import (
	"fmt"
	"task/infrastructure/seeds"
	"task/infrastructure/worker"
	"time"
	"github.com/go-co-op/gocron"
)

func task() {
	fmt.Println("seed the database cron")
	seeds.Seed()
	worker.Historylog("test")
}
func RunCron() {
	s1 := gocron.NewScheduler(time.UTC)
	s1.Every(60).Seconds().Do(task)
	s1.StartAsync()
}


