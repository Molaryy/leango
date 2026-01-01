package Logger

import (
	log "github.com/sirupsen/logrus"
)

func Fatal(nbLine int, mes string) {
	if nbLine > -1 {
		log.WithFields(log.Fields{
			"line": nbLine,
		})
	}
	log.Fatal(mes)
}

func Info(nbLine int, mes string) {
	if nbLine > -1 {
		log.WithFields(log.Fields{
			"line": nbLine,
		})
	}
	log.Info(mes)
}

func Debug(mes string) {
	log.SetLevel(log.DebugLevel)
	log.Debug(mes)
}
