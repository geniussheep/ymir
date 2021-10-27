package utils

import "time"

const DateTimeLayout = "2006-01-02 15:04:05"
const DateLayout = "2006-01-02"
const TimeLayout = "15:04:05"

// FormatDateTime format DateTimeLayout 2006-01-02 15:04:05
func FormatDateTime(t time.Time) string {
	return t.Format(DateTimeLayout)
}

// FormatDateTimeUTC format DateTimeLayout 2006-01-02 15:04:05
func FormatDateTimeUTC(t time.Time) string {
	return t.UTC().Format(DateTimeLayout)
}

func FormatDate(t time.Time) string {
	return t.Format(DateLayout)
}

func FormatTime(t time.Time) string {
	return t.Format(TimeLayout)
}

// ConvertDateTime format DateTimeLayout 2006-01-02 15:04:05
func ConvertDateTime(ts string) (time.Time, error) {
	return time.Parse(DateTimeLayout, ts)
}

func ConvertDate(ts string) (time.Time, error) {
	return time.Parse(DateLayout, ts)
}
