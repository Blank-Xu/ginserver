package db

import (
	"fmt"

	"github.com/rs/zerolog"
	"xorm.io/core"
)

type SimpleLogger struct {
	logger  *zerolog.Logger
	level   core.LogLevel
	showSQL bool
}

func NewSimpleLogger(logger *zerolog.Logger, database string, logLevel core.LogLevel) *SimpleLogger {
	log := logger.With().Str("database", database).Logger()
	return &SimpleLogger{
		logger: &log,
		level:  logLevel,
	}
}

// Error implement core.ILogger
func (s *SimpleLogger) Error(v ...interface{}) {
	if s.level <= core.LOG_ERR && len(v) > 0 {
		s.logger.Error().Msg(fmt.Sprint(v...))
	}
	return
}

// Errorf implement core.ILogger
func (s *SimpleLogger) Errorf(format string, v ...interface{}) {
	if s.level <= core.LOG_ERR {
		s.logger.Error().Msg(fmt.Sprintf(format, v...))
	}
	return
}

// Debug implement core.ILogger
func (s *SimpleLogger) Debug(v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		s.logger.Debug().Msg(fmt.Sprint(v...))
	}
	return
}

// Debugf implement core.ILogger
func (s *SimpleLogger) Debugf(format string, v ...interface{}) {
	if s.level <= core.LOG_DEBUG {
		s.logger.Debug().Msg(fmt.Sprintf(format, v...))
	}
	return
}

// Info implement core.ILogger
func (s *SimpleLogger) Info(v ...interface{}) {
	if s.level <= core.LOG_INFO {
		s.logger.Info().Msg(fmt.Sprint(v...))
	}
	return
}

// Infof implement core.ILogger
func (s *SimpleLogger) Infof(format string, v ...interface{}) {
	if s.level <= core.LOG_INFO {
		s.logger.Info().Msg(fmt.Sprintf(format, v...))
	}
	return
}

// Warn implement core.ILogger
func (s *SimpleLogger) Warn(v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		s.logger.Warn().Msg(fmt.Sprint(v...))
	}
	return
}

// Warnf implement core.ILogger
func (s *SimpleLogger) Warnf(format string, v ...interface{}) {
	if s.level <= core.LOG_WARNING {
		s.logger.Warn().Msg(fmt.Sprintf(format, v...))
	}
	return
}

// Level implement core.ILogger
func (s *SimpleLogger) Level() core.LogLevel {
	return s.level
}

// SetLevel implement core.ILogger
func (s *SimpleLogger) SetLevel(l core.LogLevel) {
	s.level = l
	return
}

// ShowSQL implement core.ILogger
func (s *SimpleLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		s.showSQL = true
		return
	}
	s.showSQL = show[0]
}

// IsShowSQL implement core.ILogger
func (s *SimpleLogger) IsShowSQL() bool {
	return s.showSQL
}
