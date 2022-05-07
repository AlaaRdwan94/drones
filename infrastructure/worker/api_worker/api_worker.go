package api_worker

import (
	"fmt"
	"task/infrastructure/seeds"
	"time"
	"github.com/go-co-op/gocron"
)

func task() {
	fmt.Println("seed the database cron")
	seeds.Seed()
}
func RunCron() {
	s1 := gocron.NewScheduler(time.UTC)
	s1.Every(60).Seconds().Do(task)
	s1.StartAsync()
}


