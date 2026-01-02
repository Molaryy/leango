package debugger

import (
	arguments "leango/src/Args"
	"leango/src/Token"
	"log"
	"time"
)

func IsDebugActivated(flags map[string]arguments.Flag) bool {
	_, existsLong := flags["--debug"]

	if existsLong {
		return true
	}
	return false
}

func PrintToken(token Token.Token) {
	log.Printf("%+v",token)
}

func PrintTokenAndSleep(token Token.Token, seconds int64) {
	PrintToken(token)
	time.Sleep(time.Second * time.Duration(seconds))
}

