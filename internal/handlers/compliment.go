package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"perfect-boyfriend/internal/database"
	"perfect-boyfriend/internal/models"
)

type Compliment struct {
	db *database.Connection
}

func NewCompliment(db *database.Connection) *Compliment {
	return &Compliment{db: db}
}

func (c Compliment) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello!",
	})
}

func (c Compliment) GetCompliments(ctx *gin.Context) {
	var compliments []models.Compliment
	c.db.Find(&compliments)

	ctx.JSON(http.StatusOK, compliments)
}

func (c Compliment) CreateCompliment(ctx *gin.Context) {
	var compliment models.Compliment

	if err := ctx.BindJSON(&compliment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please check request body",
		})
	}

	c.db.Create(&compliment)

	ctx.JSON(http.StatusCreated, compliment)
}

func (c Compliment) DeleteCompliment(ctx *gin.Context) {
	id := ctx.Param("id")

	var compliment models.Compliment
	c.db.Delete(&compliment, id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Compliment successfully deleted.",
	})
}
