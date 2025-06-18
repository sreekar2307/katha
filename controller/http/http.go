package http

import "github.com/gin-gonic/gin"

type Controller struct {
	V1 V1Controller
}

func NewController(v1 V1Controller) Controller {
	return Controller{
		V1: v1,
	}
}

type V1Controller interface {
	NewExpense() gin.HandlerFunc
	Expenses() gin.HandlerFunc
	Balances() gin.HandlerFunc
}
