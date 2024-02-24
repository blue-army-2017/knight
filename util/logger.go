package util

import "go.uber.org/zap"

func GetLogger() *zap.SugaredLogger {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer l.Sync()

	return l.Sugar()
}
