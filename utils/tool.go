package utils

import "time"

func GetTimeDaysAgo(days int) time.Time {
	return time.Now().AddDate(0, 0, -days)
}

func GetTimeHoursAgo(hours int) time.Time {
	return time.Now().Add(-time.Duration(hours) * time.Hour)
}
