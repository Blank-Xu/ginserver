package db

import (
	"time"

	"ginserver/tools/utils"
)

type JSONTime time.Time

func (p JSONTime) String() string {
	return time.Time(p).Format(utils.TimeLayoutDefault)
}

func (p JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(p).Format(utils.TimeLayoutDefault) + `"`), nil
}

func (p *JSONTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}

	t, err := time.ParseInLocation(`"`+utils.TimeLayoutDefault+`"`, string(data), time.Local)
	if err != nil {
		return err
	}
	*p = JSONTime(t)
	return nil
}
