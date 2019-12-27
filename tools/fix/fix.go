package fix

import (
	"time"
)

type Fix struct {
	TimeZone *struct {
		Name   string `yaml:"Name"`
		Offset int    `yaml:"Offset"`
	} `yaml:"TimeZone"`
}

func (p *Fix) Set() {
	// fix timezone
	time.Local = time.FixedZone(p.TimeZone.Name, p.TimeZone.Offset*3600)
}
