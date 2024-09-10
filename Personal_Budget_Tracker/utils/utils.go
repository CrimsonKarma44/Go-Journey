package utils

import "time"

func DaysToHour(days int) time.Duration {
	return time.Duration(days) * (time.Hour * 24)
}
func HoursToDays(hours time.Duration) int {
	return int(hours / 24)
}
