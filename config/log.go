package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	log *zap.Logger
}

// NewLog create new log with zap
func NewLog(logger *zap.Logger) Log {
	return Log{logger}
}

// Store log based on error value to stdout
func (l Log) Store(err error, message string, options map[string]interface{}) {
	var zapFields []zapcore.Field
	for k, v := range options {
		zapFields = append(zapFields, zap.Any(k, v))
	}

	if err != nil {
		zapFields = append(zapFields, zap.String("severity", "ERROR"))
		l.log.Error(err.Error(), zapFields...)
	} else {
		zapFields = append(zapFields, zap.String("severity", "INFO"))
		l.log.Info(message, zapFields...)
	}
}
