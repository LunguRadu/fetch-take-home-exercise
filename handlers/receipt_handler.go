package handlers

import (
	"net/http"
	"fetch-take-home-exercise/models"
	"fetch-take-home-exercise/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var receipts = make(map[string]models.Receipt)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid"})
		return
	}

	receipt.ID = uuid.New().String()
	receipt.Points = services.CalculatePoints(receipt)

	receipts[receipt.ID] = receipt

	c.JSON(http.StatusOK, gin.H{"id": receipt.ID})
}

func GetReceiptPoints(c *gin.Context) {
	id := c.Param("id")
	receipt, exists := receipts[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": receipt.Points})
}
