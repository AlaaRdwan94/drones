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
	
}

func task2() {
	worker.Historylog("task2 -- ")
	seeds.UpdateDroneToDelevering()
}
func task3() {
	worker.Historylog("task3 -- ")

	seeds.UpdateDroneToReturing()
}
func task4()  {
	worker.Historylog("task4 -- ")

	seeds.UpdateDroneToDelevered()
}
func RunCron() {
	s1 := gocron.NewScheduler(time.UTC)
	s1.Every(6).Seconds().Do(task)
	s1.Every(6).Second().Do(task2)
	s1.Every(30).Second().Do(task3)
	s1.Every(30).Second().Do(task4)
	s1.StartAsync()
}


