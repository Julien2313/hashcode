package main

import (
	"os"

	"runtime"

	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func startLogger() {
	log.SetLevel(logLevel)

	// wrap the log output
	if runtime.GOOS == "windows" {
		// force colors on for TextFormatter
		log.Formatter = &logrus.TextFormatter{
			ForceColors: true,
		}

		log.Out = ansicolor.NewAnsiColorWriter(os.Stdout)
	} else {
		log.Out = os.Stdout
	}
}
