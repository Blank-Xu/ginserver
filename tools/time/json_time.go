package time

import (
	"time"
)

type JSONTime time.Time

func NewJSONTime()JSONTime  {
	return JSONTime(time.Now())
}

func (p JSONTime) String() string {
	return time.Time(p).Format(LayoutDefault)
}

func (p JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(p).Format(LayoutDefault) + `"`), nil
}

func (p *JSONTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	s := string(data)
	if s == "null" {
		return nil
	}

	t, err := time.ParseInLocation(`"`+LayoutDefault+`"`, s, time.Local)
	if err != nil {
		return err
	}
	
	*p = JSONTime(t)
	return nil
}
