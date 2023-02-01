package server

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"log"
	"os"
	"perfect-boyfriend/internal/database"
	"perfect-boyfriend/internal/jobs"
	"perfect-boyfriend/internal/services"
	"time"
)

type Scheduler struct {
	db  *database.Connection
	bot *services.TelegramBot
}

func NewScheduler(db *database.Connection, bot *services.TelegramBot) *Scheduler {
	return &Scheduler{
		db:  db,
		bot: bot,
	}
}

func (s Scheduler) Start() {
	location, err := time.LoadLocation(os.Getenv("SCHEDULER_LOCATION"))
	if err != nil {
		fmt.Println(err)
	}

	scheduler := gocron.NewScheduler(location)

	_, err = scheduler.Every(1).Day().At(os.Getenv("SEND_GREETING_TIME")).Do(jobs.NewSendGreeting(s.db, s.bot).Handle)
	if err != nil {
		log.Println(err)
	}

	scheduler.StartAsync()
}
