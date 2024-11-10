package services

import (
	"fetch-take-home-exercise/models"
	"testing"
)

func TestCalculatePointsRetailerNameRule(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Whole Foods",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
	}

	points := CalculatePoints(receipt)

	expectedPoints := 32

	if points != expectedPoints {
		t.Errorf("Expected %d points but got %d", expectedPoints, points)
	}
}

func TestCalculatePointsTotalAmount(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.00", // multiple of 0.25
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
	}

	points := CalculatePoints(receipt)

	expectedPoints := 103

	if points != expectedPoints {
		t.Errorf("Expected %d points but got %d", expectedPoints, points)
	}
}

func TestCalculatePointsTwoItemsOnReceipt(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
		},
	}

	points := CalculatePoints(receipt)

	expectedPoints := 20

	if points != expectedPoints {
		t.Errorf("Expected %d points but got %d", expectedPoints, points)
	}
}

func TestCalculatePointsTrimmedLength(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []models.Item{
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		},
	}

	points := CalculatePoints(receipt)

	expectedPoints := 15

	if points != expectedPoints {
		t.Errorf("Expected %d points but got %d", expectedPoints, points)
	}
}

func TestCalculatePointsOddDayRule(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01", // Odd day
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
	}

	points := CalculatePoints(receipt)

	expectedPoints := 28

	if points != expectedPoints {
		t.Errorf("Expected %d points but got %d", expectedPoints, points)
	}
}

func TestCalculatePointsEvenDayRule(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-02", // Even day
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
	}

	points := CalculatePoints(receipt)

	expectedPoints := 22

	if points != expectedPoints {
		t.Errorf("Expected %d points but got %d", expectedPoints, points)
	}
}

func TestCalculatePointsPurchaseTime(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "14:01",
		Total:        "35.35",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
	}

	points := CalculatePoints(receipt)

	expectedPoints := 38

	if points != expectedPoints {
		t.Errorf("Expected %d points but got %d", expectedPoints, points)
	}
}
