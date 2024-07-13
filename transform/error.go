package transform

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func loggerInit() *zap.Logger {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(zap.InfoLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))
	return logger
}
func ErrApi(err error) error {
	var l = loggerInit()
	l.Error(err.Error())
	defer l.Sync()
	return err
}
