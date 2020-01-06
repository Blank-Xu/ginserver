package log

import (
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Option struct {
	RecordFile   bool          `yaml:"RecordFile"`
	FileName     string        `yaml:"FileName"`
	LinkName     string        `yaml:"LinkName"`
	RotateHour   int           `yaml:"RotateHour"`
	ReportCaller bool          `yaml:"ReportCaller"`
	TimeLayout   string        `yaml:"TimeLayout"`
	Level        zerolog.Level `yaml:"Level"`
}

func (p *Option) Init() error {
	zerolog.SetGlobalLevel(p.Level)
	zerolog.TimeFieldFormat = p.TimeLayout

	var ctx zerolog.Context
	if p.RecordFile {
		rotate, err := rotatelogs.New(
			p.FileName,
			// WithLinkName为最新的日志建立软连接
			rotatelogs.WithLinkName(p.LinkName),
			// WithRotationTime设置日志分割的时间
			rotatelogs.WithRotationTime(time.Hour*time.Duration(p.RotateHour)),
		)
		if err != nil {
			return err
		}

		ctx = zerolog.New(rotate).With().Timestamp()
	} else {
		ctx = zerolog.New(os.Stdout).With().Timestamp()
	}
	if p.ReportCaller {
		ctx = ctx.Caller()
	}
	log.Logger = ctx.Logger()

	log.Info().Msg("log load success")

	return nil
}
