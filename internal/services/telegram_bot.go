package services

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	"math/rand"
	"os"
	"perfect-boyfriend/internal/clients"
	"perfect-boyfriend/internal/database"
	"perfect-boyfriend/internal/models"
	"time"
)

type TelegramBot struct {
	db  *database.Connection
	bot *tgbotapi.BotAPI
}

func NewTelegramBot(db *database.Connection) *TelegramBot {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	return &TelegramBot{
		db:  db,
		bot: bot,
	}
}

func (tgb TelegramBot) Start() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := tgb.bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message != nil && tgb.checkAccess(update.Message) {
			go tgb.handleIncomingMessage(update.Message)
		}
	}
}

func (tgb TelegramBot) checkAccess(message *tgbotapi.Message) bool {
	var chat models.Chat
	tgb.db.Find(&chat, message.Chat.ID)

	if chat.ID == 0 {
		if message.Text == os.Getenv("PASSWORD") {
			tgb.db.Create(models.Chat{
				ID: message.Chat.ID,
			})

			return true
		} else {
			tgb.SendNewMessage("Please enter a password.", message.Chat.ID)

			return false
		}
	}

	return true
}

func (tgb TelegramBot) handleIncomingMessage(message *tgbotapi.Message) {
	var weather string
	var err error
	var compliments []models.Compliment

	rand.Seed(time.Now().UnixNano())
	if rand.Intn(3) == 1 {
		weather, err = clients.GetCurrentWeather()
		if err != nil {
			log.Println(err)
		} else {
			tgb.db.Where("is_greeting = false AND proper_weather = ?", weather).Find(&compliments)
		}
	}

	if len(compliments) == 0 {
		tgb.db.Where("is_greeting = false AND proper_weather = ''").Find(&compliments)
	}

	randomCompliment := compliments[rand.Intn(len(compliments))]

	tgb.SendNewMessage(randomCompliment.Text, message.Chat.ID)
}

func (tgb TelegramBot) SendNewMessage(message string, chatId int64) {
	msg := tgbotapi.NewMessage(chatId, message)

	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(getRandomEmoji()),
		),
	)

	if _, err := tgb.bot.Send(msg); err != nil {
		log.Println(err)
	}
}
