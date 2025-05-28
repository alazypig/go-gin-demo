package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func InitLogger() {
	appLogger, err := InitZapLogger("error.log")
	if err != nil {
		panic(err)
	}
	defer appLogger.Sync()

	log = appLogger

	log.Info("Logger initialized")
}

func Info(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	log.Error(msg, fields...)
}

func Debug(msg string, fields ...zapcore.Field) {
	log.Debug(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	log.Warn(msg, fields...)
}
