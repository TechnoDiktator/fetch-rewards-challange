package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

// InitializeLogger initializes the logger with necessary configurations
func InitializeLogger() {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetLevel(logrus.DebugLevel) // or change to InfoLevel for less verbosity
}
