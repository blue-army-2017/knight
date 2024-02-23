package util

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func init() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer l.Sync()

	logger = l.Sugar()
}

func LogDebug(msg string, data ...any) {
	logger.Debugw(msg, data...)
}

func LogInfo(msg string, data ...any) {
	logger.Infow(msg, data...)
}

func LogWarn(msg string, data ...any) {
	logger.Warnw(msg, data...)
}

func LogError(msg string, data ...any) {
	logger.Errorw(msg, data...)
}

func LogFatal(msg string, data ...any) {
	logger.Fatalw(msg, data...)
}
