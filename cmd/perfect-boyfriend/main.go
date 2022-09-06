package main

import (
	"perfect-boyfriend/internal/database"
	"perfect-boyfriend/internal/server"
	"perfect-boyfriend/internal/services"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(3)

	db := database.NewConnection()
	db.RunMigrations()

	httpServer := server.NewHTTPServer(db)
	httpServer.InitRoutes()
	go httpServer.Start()

	telegramBot := services.NewTelegramBot(db)
	go telegramBot.Start()

	scheduler := server.NewScheduler(db, telegramBot)
	go scheduler.Start()

	waitGroup.Wait()
}
