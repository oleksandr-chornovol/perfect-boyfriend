package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"perfect-boyfriend/internal/models"
)

type Connection struct {
	*gorm.DB
}

func NewConnection() *Connection {
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &Connection{db}
}

func (c Connection) RunMigrations() {
	allModels := [...]any{
		models.Compliment{},
	}

	for _, model := range allModels {
		if err := c.AutoMigrate(&model); err != nil {
			fmt.Println(err)
		}
	}
}
