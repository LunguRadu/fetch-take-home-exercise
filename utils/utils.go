package utils

import (
	"time"
)

func IsAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func GetDay(dateStr string) (int, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return 0, err
	}
	return date.Day(), nil
}

func GetTime(timeStr string) (int, int, error) {
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return 0, 0, err
	}
	return t.Hour(), t.Minute(), nil
}
