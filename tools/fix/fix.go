package fix

import (
	"time"
)

type Fix struct {
	TimeZone *struct {
		Name       string `yaml:"Name"`
		OffsetHour int    `yaml:"OffsetHour"`
	} `yaml:"TimeZone"`
}

func (p *Fix) Init() {
	// fix timezone
	time.Local = time.FixedZone(p.TimeZone.Name, p.TimeZone.OffsetHour*3600)
}
