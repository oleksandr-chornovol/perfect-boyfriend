package jobs

import (
	"log"
	"math/rand"
	"perfect-boyfriend/internal/cache"
	"perfect-boyfriend/internal/clients"
	"perfect-boyfriend/internal/database"
	"perfect-boyfriend/internal/models"
	"perfect-boyfriend/internal/services"
	"time"
)

type SendGreeting struct {
	bot *services.TelegramBot
	db  *database.Connection
}

func NewSendGreeting(db *database.Connection, bot *services.TelegramBot) *SendGreeting {
	return &SendGreeting{
		bot: bot,
		db:  db,
	}
}

func (sg SendGreeting) Handle() {
	var compliments []models.Compliment

	weather, err := clients.GetCurrentWeather()
	if err != nil {
		log.Println(err)
	} else {
		sg.db.Where("is_greeting = true AND proper_weather = ?", weather).Find(&compliments)
	}

	if len(compliments) == 0 {
		sg.db.Where("is_greeting = true AND proper_weather = ''").Find(&compliments)
	}

	for _, chat := range cache.GetAllChats() {
		rand.Seed(time.Now().UnixNano())
		randomCompliment := compliments[rand.Intn(len(compliments))]

		go sg.bot.SendNewMessage(randomCompliment.Text, chat.ID)
	}
}
