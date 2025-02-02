package utils

import "time"

func GetCurrentFormattedDateTime() string {
	currentTime := time.Now()

	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	return formattedTime
}

func GetFormattedDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
