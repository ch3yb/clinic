package service

import (
	"database/sql"
	db "github.com/ch3yb/clinic/api/database"
	"github.com/ch3yb/clinic/api/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

type Service struct {
	db     *sql.DB
	Logger *zap.Logger
	Err    *errors.Err
}

func New() *Service {
	database, err := db.StartDatabase()
	if err != nil {
		log.Println("err when start  Database on service.go : ", err)
		return nil
	}

	return &Service{
		db:     database,
		Logger: loggerInit(),
		Err: &errors.Err{
			Lang: "fr", //get it auto
		},
	}
}
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
