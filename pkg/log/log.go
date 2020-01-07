package log

import (
	"os"
	"path/filepath"
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
		logDir := filepath.Dir(p.FileName)
		if err := os.MkdirAll(logDir, 0766); err != nil {
			return err
		}

		linkDir := filepath.Dir(p.LinkName)
		if err := os.MkdirAll(linkDir, 0766); err != nil {
			return err
		}

		linkName, err := filepath.Abs(p.LinkName)
		if err != nil {
			return err
		}

		rotate, err := rotatelogs.New(
			p.FileName,
			// WithLinkName为最新的日志建立软连接
			rotatelogs.WithLinkName(linkName),
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
