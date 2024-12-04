package services

import (
	"fetch-take-home-exercise/models"
	"fmt"
	"github.com/sirupsen/logrus"
)

var (
	users  map[int]models.User
	logger *logrus.Logger
)

func init() {
	users = make(map[int]models.User)

	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
}

func getUserIfExists(id int) (models.User, error) {
	user, exists := users[id]
	if !exists {
		logger.WithFields(logrus.Fields{
			"user_id": id,
			"action":  "get_user",
		}).Warn("User not found")
		return models.User{}, fmt.Errorf("user with id %d not found", id)
	}
	return user, nil
}

func AddUser(id int, user models.User) (models.User, error) {
	_, exists := users[id]
	if exists {
		logger.WithFields(logrus.Fields{
			"user_id": id,
		}).Warn("Attempt to add existing user")
		return models.User{}, fmt.Errorf("user with id %d already exists", id)
	}

	users[id] = user
	logger.WithFields(logrus.Fields{
		"user_id": id,
	}).Info("User added")
	return user, nil
}

func GetUser(id int) (models.User, error) {
	return getUserIfExists((id))
}

func GetUserPoints(id int) (int, error) {
	user, err := getUserIfExists(id)
	if err != nil {
		return 0, err
	} else {
		logger.WithFields(logrus.Fields{
			"user_id": id,
			"points":  user.Points,
		}).Info("User points retrieved")
		return user.Points, nil
	}
}

func AddUserReceipt(receipt models.Receipt, userId int) (models.User, error) {
	user, err := getUserIfExists(userId)
	if err != nil {
		return models.User{}, err
	}
	user.Receipts = append(user.Receipts, receipt)
	user.Points += receipt.Points

	users[userId] = user
	logger.WithFields(logrus.Fields{
		"user_id":        userId,
		"receipt_points": receipt.Points,
		"total_points":   user.Points,
	}).Info("Receipt added to user")

	return user, nil
}

func GetUserReceipts(userId int) ([]models.Receipt, error) {
	user, err := getUserIfExists(userId)
	if err != nil {
		return []models.Receipt{}, err
	}
	logger.WithFields(logrus.Fields{
		"user_id":       userId,
		"receipt_count": len(user.Receipts),
	}).Info("User receipts retrieved")
	return user.Receipts, nil

}

func ReddemPoints(userId int, points int) (int, error) {
	user, err := getUserIfExists(userId)
	if err != nil {
		return 0, err
	}

	if user.Points < points {
		logger.WithFields(logrus.Fields{
			"user_id":          userId,
			"requested_points": points,
			"available_points": user.Points,
		}).Warn("Insufficient points for redemption")
		return 0, fmt.Errorf("user doesn't have enough points")
	}
	pointsLeft := user.Points - points
	user.Points = pointsLeft
	users[userId] = user

	logger.WithFields(logrus.Fields{
		"user_id":          userId,
		"redeemed_points":  points,
		"remaining_points": pointsLeft,
	}).Info("Points redeemed")

	return pointsLeft, nil
}
