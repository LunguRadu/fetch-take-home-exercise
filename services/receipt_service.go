package services

import (
	"math"
	"strconv"
	"strings"
	"fetch-take-home-exercise/models"
	"fetch-take-home-exercise/utils"
)

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// One point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
		if utils.IsAlphanumeric(char) {
			points++
		}
	}

	// 50 points if the total is a round dollar amount with no cents
	if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil {
		if total == float64(int(total)) {
			points += 50
		}
	}

	// 25 points if the total is a multiple of 0.25
	if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil {
		if math.Mod(total, 0.25) == 0 {
			points += 25
		}
	}

	// 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	// Points based on item description length
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			if price, err := strconv.ParseFloat(item.Price, 64); err == nil {
				points += int(math.Ceil(price * 0.2))
			}
		}
	}

	// 6 points if the day in the purchase date is odd
	if day, err := utils.GetDay(receipt.PurchaseDate); err == nil {
		if day%2 != 0 {
			points += 6
		}
	}

	// 10 points if the purchase time is between 2:00pm and 4:00pm
	if hour, minute, err := utils.GetTime(receipt.PurchaseTime); err == nil {
		if (hour >= 14 && hour < 16) || (hour == 16 && minute == 0) {
			points += 10
		}
	}

	return points
}
