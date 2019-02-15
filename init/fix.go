package init

import (
	"time"

	"ginserver/init/config"
)

func fix() {
	var cfg = config.GetConfig().Fix
	// fix timezone
	time.Local = time.FixedZone(cfg.TimeZone.Name, cfg.TimeZone.Offset*3600)
}
