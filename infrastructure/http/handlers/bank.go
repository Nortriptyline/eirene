package handlers

import (
	"github.com/Nortriptyline/Eirene/application"
	"github.com/gin-gonic/gin"
)

type BankRoutes struct {
	handlers *application.Handlers
}

func NewBankRoutes(handlers *application.Handlers) *BankRoutes {
	return &BankRoutes{}
}

func RegisterBankRoutes(group *gin.RouterGroup, handlers *application.Handlers) {
	bankRoutes := NewBankRoutes(handlers)

}

func (r *BankRoutes) CreateBank() {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
