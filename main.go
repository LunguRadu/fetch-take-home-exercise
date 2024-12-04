package main

import (
	"fetch-take-home-exercise/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/receipts/process", handlers.ProcessReceipt)
	r.GET("/receipts/:id/points", handlers.GetReceiptPoints)

	r.POST("/user/create", handlers.AddUser)
	r.POST("user/:id/receipts", handlers.AddReceiptUser)
	r.GET("/user/:id", handlers.GetUser)
	r.GET("/user/:id/points", handlers.GetUserPoints)
	r.GET("/user/:id/receipts", handlers.GetUserReceipts)
	r.PUT("/user/:id/points", handlers.ReddemPoints)
	r.Run(":8080")
}
