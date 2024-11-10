package main

import (
	"github.com/gin-gonic/gin"
	"fetch-take-home-exercise/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/receipts/process", handlers.ProcessReceipt)
	r.GET("/receipts/:id/points", handlers.GetReceiptPoints)

	r.Run(":8080")
}