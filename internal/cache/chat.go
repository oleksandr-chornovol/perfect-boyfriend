package cache

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var chats = make(map[int64]*tgbotapi.Chat)

func AddChat(chat *tgbotapi.Chat) {
	chats[chat.ID] = chat
}

func GetAllChats() map[int64]*tgbotapi.Chat {
	return chats
}
