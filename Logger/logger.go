package Logger

import (
	log "github.com/sirupsen/logrus"
)

func Fatal(nbLine int, mes string) {
	log.WithFields(log.Fields{
		"line": nbLine,
	}).Fatal(mes)
}

func Info(nbLine int, mes string) {
	log.WithFields(log.Fields{
		"line": nbLine,
	}).Fatal(mes)
}
