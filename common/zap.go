package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.SugaredLogger
)

func init() {
	// log file name
	fileName := "micro.log"
	syncWriter := zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    512,
			MaxBackups: 0,
			LocalTime:  true,
			Compress:   true,
		})

	// create encoder
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		syncWriter,
		zap.NewAtomicLevelAt(zap.DebugLevel))
	log := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1))
	logger = log.Sugar()
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}
