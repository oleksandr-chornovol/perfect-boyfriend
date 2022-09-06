package services

import (
	"math/rand"
	"time"
)

var emojis = [...]string{"♥️", "🐱", "🐭", "🐯", "🐽", "🦄", "🐝", "💐", "🌼"}

func getRandomEmoji() string {
	rand.Seed(time.Now().UnixNano())

	return emojis[rand.Intn(len(emojis))]
}
