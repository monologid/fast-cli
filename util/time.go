package util

import "time"

// GetDateTime is used to get current time
func GetDateTime() string {
	layout := "2006-01-02 15:04:05"
	return time.Now().Format(layout)
}
