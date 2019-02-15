package utils

import (
	"time"
)

const (
	TimeLayoutDefault = "2006-01-02 15:04:05"
	TimeLayoutUTC     = "2006-01-02 15:04:05 (UTC-07:00)"
)

func TimeNowFormat(layout ...string) string {
	if len(layout) == 0 {
		return time.Now().Format(TimeLayoutDefault)
	}
	return time.Now().Format(layout[0])
}

func TimeFormat(time time.Time, layout ...string) string {
	if len(layout) == 0 {
		return time.Format(TimeLayoutDefault)
	}
	return time.Format(layout[0])
}

func GetDaysBetween(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int(t2.Sub(t1).Hours() / 24)
}

func GetAge(birth time.Time) int {
	today := time.Now()
	age := today.Year() - birth.Year()
	if time.Now().Month() < birth.Month() {
		age--
	}
	if today.Month() == birth.Month() && today.Day() > birth.Day() {
		age--
	}
	return age
}
