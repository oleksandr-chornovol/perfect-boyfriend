package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"perfect-boyfriend/internal/models"
)

type Connection struct {
	*gorm.DB
}

func NewConnection() *Connection {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("RDS_USERNAME"), os.Getenv("RDS_PASSWORD"), os.Getenv("RDS_HOSTNAME"), os.Getenv("RDS_PORT"), os.Getenv("RDS_DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	return &Connection{db}
}

func (c Connection) RunMigrations() {
	allModels := [...]any{
		models.Compliment{},
		models.Chat{},
	}

	for _, model := range allModels {
		if err := c.AutoMigrate(&model); err != nil {
			fmt.Println(err)
		}
	}
}
