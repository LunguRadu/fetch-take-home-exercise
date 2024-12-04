package handlers

import (
	"fetch-take-home-exercise/constants"
	"fetch-take-home-exercise/models"
	"fetch-take-home-exercise/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The JSON is invalid"})
		return
	}

	user, service_error := services.AddUser(user.ID, user)

	if service_error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": service_error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var intId, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.InvalidIdFormat})
		return
	}
	user, service_err := services.GetUser(intId)
	if service_err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "there was an error fetching the user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUserPoints(c *gin.Context) {
	id := c.Param(("id"))
	var intId, err = strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.InvalidIdFormat})
		return
	}

	points, service_err := services.GetUserPoints(intId)

	if service_err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "there was an error fetching the user points"})
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}

func AddReceiptUser(c *gin.Context) {
	id := c.Param(("id"))
	var intId, idErr = strconv.Atoi(id)
	var receipt models.Receipt

	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.InvalidIdFormat})
		return
	}

	jsonErr := c.ShouldBindJSON(&receipt)

	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.InvalidIdFormat})
		return
	}

	user, serviceErr := services.AddUserReceipt(receipt, intId)

	if serviceErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": serviceErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUserReceipts(c *gin.Context) {
	id := c.Param("id")
	var intId, idErr = strconv.Atoi(id)

	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.InvalidIdFormat})
	}

	receipts, serviceErr := services.GetUserReceipts(intId)

	if serviceErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": serviceErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"receipts": receipts, "count": len(receipts)})
}

func ReddemPoints(c *gin.Context) {
	id := c.Param("id")
	var intId, idErr = strconv.Atoi(id)

	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.InvalidIdFormat})
		return
	}

	var req models.ReedemRequest
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	points, serviceError := services.ReddemPoints(intId, req.Points)

	if serviceError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": serviceError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
