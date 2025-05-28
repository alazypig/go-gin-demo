package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZapLogger(logPath string) (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	cfg.OutputPaths = []string{logPath}
	cfg.ErrorOutputPaths = []string{logPath}
	cfg.InitialFields = map[string]interface{}{
		"service": "go-gin",
	}

	logger, err := cfg.Build()

	if err != nil {
		return nil, err
	}

	logger.Info("Zap logger initialized", zap.String("file", logPath))

	return logger, nil
}
