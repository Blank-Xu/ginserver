package util

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

type JSONTime time.Time

func (p JSONTime) String() string {
	return time.Time(p).Format(TimeLayoutDefault)
}

func (p JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(p).Format(TimeLayoutDefault) + `"`), nil
}

func (p *JSONTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}

	t, err := time.ParseInLocation(`"`+TimeLayoutDefault+`"`, string(data), time.Local)
	if err != nil {
		return err
	}
	*p = JSONTime(t)
	return nil
}
