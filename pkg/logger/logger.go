package logger

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var gloabalLogger *logrus.Logger
var once sync.Once

type customLogger struct{}

func (c *customLogger) Format(entry *logrus.Entry) ([]byte, error) {

	logMessage := fmt.Sprintf("%s - %s: %s\n",
		time.Now().Format("2006-01-03 15:04:05"),
		strings.ToUpper(entry.Level.String()),
		entry.Message)
	return []byte(logMessage), nil

}

func initiateLogger() *logrus.Logger {

	once.Do(func() {
		gloabalLogger = logrus.New()
		gloabalLogger.SetFormatter(&customLogger{})

		// You can set the log level here
		gloabalLogger.SetLevel(logrus.DebugLevel)
	})

	return gloabalLogger

}

func Info(message string) {

	logger := initiateLogger()
	logger.Info(message)

}

func Warn(message string) {

	logger := initiateLogger()
	logger.Warn(message)

}

func Error(message string) {

	logger := initiateLogger()
	logger.Error(message)

}

func Debug(message string) {

	logger := initiateLogger()
	logger.Debug(message)

}
