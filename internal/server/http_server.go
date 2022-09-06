package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"perfect-boyfriend/internal/database"
	"perfect-boyfriend/internal/handlers"
)

type HTTPServer struct {
	db     *database.Connection
	engine *gin.Engine
}

func NewHTTPServer(db *database.Connection) *HTTPServer {
	return &HTTPServer{
		db:     db,
		engine: gin.Default(),
	}
}

func (s HTTPServer) InitRoutes() {
	complimentHandler := handlers.NewCompliment(s.db)

	s.engine.GET("/compliments", complimentHandler.GetCompliments)
	s.engine.POST("/compliments", complimentHandler.CreateCompliment)
	s.engine.DELETE("/compliments", complimentHandler.DeleteCompliment)
}

func (s HTTPServer) Start() {
	if err := s.engine.Run(":" + os.Getenv("PORT")); err != nil {
		log.Println(err)
	}
}
