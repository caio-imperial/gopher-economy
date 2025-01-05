package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type AppLogger struct {
	Sugar *zap.SugaredLogger
}

var appLogger *AppLogger

func GetLogger() *AppLogger {
	return appLogger
}

func NewLogger() (err error) {
	loggerLevel := zap.NewAtomicLevelAt(zap.InfoLevel)
	loggerDevelopment := false
	loggerEncoding := "json"
	if os.Getenv("APP_ENV") == "development" {
		loggerLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
		loggerDevelopment = true
		loggerEncoding = "console"
	}
	config := zap.Config{
		Level:            loggerLevel,
		Development:      loggerDevelopment,
		Encoding:         loggerEncoding,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			MessageKey:     "msg",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		},
	}

	zapLogger, err := config.Build()
	if err != nil {
		return
	}

	appLogger = &AppLogger{Sugar: zapLogger.Sugar()}
	appLogger.Debug("Development Mode", "development", true)
	return
}

func (l *AppLogger) Sync() {
	_ = l.Sugar.Sync()
}

func loggerWithCorrectCaller(logger *AppLogger) *zap.SugaredLogger {
	return logger.Sugar.Desugar().WithOptions(zap.AddCallerSkip(1)).Sugar()
}

func (l *AppLogger) Info(message string, fields ...interface{}) {
	loggerWithCorrectCaller(l).Infow(message, fields...)
}

func (l *AppLogger) Error(message string, fields ...interface{}) {
	loggerWithCorrectCaller(l).Errorw(message, fields...)
}

func (l *AppLogger) Debug(message string, fields ...interface{}) {
	loggerWithCorrectCaller(l).Debugw(message, fields...)
}

func (l *AppLogger) Warn(message string, fields ...interface{}) {
	loggerWithCorrectCaller(l).Warnw(message, fields...)
}
