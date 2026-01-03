package debugger

import (
	"fmt"
	arguments "leango/src/Args"
	"leango/src/Token"
	"time"
)

const (
	RESET  = "\033[0m"
	BLUE   = "\033[34m"
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
)

func IsDebugActivated(flags map[string]arguments.Flag) bool {
	_, existsLong := flags["--debug"]

	if existsLong {
		return true
	}
	return false
}

func PrintToken(token Token.Token) {
	str := fmt.Sprintf("Type: %s ", token.Type)
	switch v := token.Value.(type) {
	case int:
		str += fmt.Sprintf("Value: %d\n", v)
	case string:
		str += fmt.Sprintf("Value: %s\n", v)
	case bool:
		str += fmt.Sprintf("Value: %t\n", v)
	case byte:
		str += fmt.Sprintf("Value: %c\n", v)
	default:
		str += fmt.Sprintf("Unknown type: %T\n", v)
	}
	fmt.Printf(str)
}

func PrintTokenAndSleep(token Token.Token, seconds int64) {
	PrintToken(token)
	time.Sleep(time.Second * time.Duration(seconds))
}
