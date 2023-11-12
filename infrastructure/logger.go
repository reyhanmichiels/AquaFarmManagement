package infrastructure

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func CreateLogger() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	Logger = logger
}
