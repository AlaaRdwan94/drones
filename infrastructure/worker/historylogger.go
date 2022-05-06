package worker

import (
	"log"
	"os"
)

func Historylog(msg string) {
	f, err := os.OpenFile("history.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(msg)
}