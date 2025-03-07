package logger

import (
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() error {
	Log = logrus.New()
	Log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	Log.SetLevel(logrus.InfoLevel)
	return nil
}
