package util

import "time"

var TimeDefaultLayout = "2006-01-02 15:04:05"

func TimeNowFormat(layout ...string) string {
	if len(layout) == 0 {
		return time.Now().Format(TimeDefaultLayout)
	}
	return time.Now().Format(layout[0])
}

func TimeFormat(value string, layout ...string) string {
	if len(layout) == 0 {
		return time.Now().Format(TimeDefaultLayout)
	}
	return time.Now().Format(layout[0])
}
