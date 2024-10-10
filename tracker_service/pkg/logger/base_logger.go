package logger

import (
	"fmt"
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
	LevelPanic = "panic"
	LevelFatal = "fatal"
)

type Logger struct {
	logger *zap.Logger
}

type LoggerConfig struct {
	Level           string
	LogDirectory    string
	AppName         string
	SamplingEnabled bool
}

func NewLogger(cfg LoggerConfig) (*Logger, error) {
	encConf := zap.NewProductionEncoderConfig()
	encConf.TimeKey = "timestamp"
	encConf.MessageKey = "message"
	encConf.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConf := zap.Config{
		Level:            parseLevel(cfg.Level),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    ecszap.ECSCompatibleEncoderConfig(encConf),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields: map[string]interface{}{
			"service_name": cfg.AppName,
		},
	}
	if cfg.SamplingEnabled {
		zapConf.Sampling = &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		}
	}
	if cfg.LogDirectory != "" {
		os.MkdirAll(cfg.LogDirectory, os.ModePerm)
		zapConf.OutputPaths = append(zapConf.OutputPaths, cfg.LogDirectory+"/app.log")
		zapConf.ErrorOutputPaths = append(zapConf.ErrorOutputPaths, cfg.LogDirectory+"/app.log")
	}

	zapLogger, err := zapConf.Build(ecszap.WrapCoreOption(), zap.AddCaller(), zap.AddCallerSkip(1))
	return &Logger{logger: zapLogger}, err
}

func parseLevel(level string) zap.AtomicLevel {
	switch level {
	case zapcore.DebugLevel.String():
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case zapcore.InfoLevel.String():
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case zapcore.WarnLevel.String():
		return zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case zapcore.ErrorLevel.String():
		return zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case zapcore.PanicLevel.String():
		return zap.NewAtomicLevelAt(zapcore.PanicLevel)
	case zapcore.FatalLevel.String():
		return zap.NewAtomicLevelAt(zapcore.FatalLevel)
	default:
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
}

type Field struct {
	Key   string
	Value interface{}
}

func (l *Logger) Debug(message string, fields ...Field) {
	if l.logger != nil {
		l.logger.Debug(message, convertToZapFields(fields)...)
	}
}
func (l *Logger) Info(message string, fields ...Field) {
	if l.logger != nil {
		l.logger.Info(message, convertToZapFields(fields)...)
	}
}
func (l *Logger) Warn(message string, fields ...Field) {
	if l.logger != nil {
		l.logger.Warn(message, convertToZapFields(fields)...)
	}
}
func (l *Logger) Error(message string, fields ...Field) {
	if l.logger != nil {
		l.logger.Error(message, convertToZapFields(fields)...)
	}
}
func (l *Logger) Panic(message string, fields ...Field) {
	if l.logger != nil {
		l.logger.Panic(message, convertToZapFields(fields)...)
	}
}
func (l *Logger) Fatal(message string, fields ...Field) {
	if l.logger != nil {
		l.logger.Fatal(message, convertToZapFields(fields)...)
	}
}

func convertToZapFields(fields []Field) []zap.Field {
	values := []zap.Field{}
	for _, val := range fields {
		if _, strType := val.Value.(fmt.Stringer); strType {
			values = append(values, zap.Reflect(val.Key, val.Value))
		} else {
			values = append(values, zap.Any(val.Key, val.Value))
		}
	}
	return values
}

func ConvertMapToFields(data map[string]interface{}) (res []Field) {
	for i, v := range data {
		res = append(res, Field{
			Key:   i,
			Value: v,
		})
	}
	return
}
