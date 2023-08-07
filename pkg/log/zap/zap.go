package zap

import (
	"os"

	"github.com/professionsforall/hexagonal-template/pkg/log/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	core *zap.SugaredLogger
}

func NewZapLogger() logger.AppLogger {
	zapLogger := new(zapLogger)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	enc := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewCore(enc, os.Stdout, zapcore.InfoLevel)
	z := zap.New(core)
	sugarLogger := z.Sugar()
	zapLogger.core = sugarLogger
	return zapLogger
}

func (z *zapLogger) Error(err ...any) {
	z.core.Error(err...)
}

// Fatal implements log.Logger.
func (z *zapLogger) Fatal(msg string) {
	z.core.Fatal(msg)
}

// Info implements log.Logger.
func (z *zapLogger) Info(msg string) {
	z.core.Info(msg)
}

// Panic implements log.Logger.
func (z *zapLogger) Panic(err error) {
	z.core.Panic(err)
}
