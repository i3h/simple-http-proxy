package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init_log() {
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	Log.SetFormatter(formatter)
	Log.Out = os.Stdout
}
