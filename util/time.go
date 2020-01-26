package util

import "time"

func GetISO8601TimeStamp() string {
	return Date("2006-01-02T15:04:05+0800", Time())
}

func GetNowDatetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Time() int64 {
	return time.Now().Unix()
}

func Date(format string, timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format)
}
