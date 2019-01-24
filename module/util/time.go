package util

import "time"

const (
	TimeLayoutDefault = "2006-01-02 15:04:05"
)

func TimeNowFormat(layout ...string) string {
	if len(layout) == 0 {
		return time.Now().Format(TimeLayoutDefault)
	}
	return time.Now().Format(layout[0])
}

func TimeFormat(value string, layout ...string) string {
	if len(layout) == 0 {
		return time.Now().Format(TimeLayoutDefault)
	}
	return time.Now().Format(layout[0])
}
