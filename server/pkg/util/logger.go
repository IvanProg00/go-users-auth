package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func NewLogger() {
	Log = logrus.New()
	Log.Out = os.Stdout
	Log.Formatter = &logrus.TextFormatter{
		FullTimestamp:             true,
		EnvironmentOverrideColors: true,
	}
}
