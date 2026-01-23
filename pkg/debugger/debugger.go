package debugger

import (
	"fmt"
	arguments "leango/src/args"
	"leango/src/token"
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

func PrintToken(token token.Token) {
	str := fmt.Sprintf("Type: %s ", token.Type)

	if token.HasValue {
		switch v := token.Value.(type) {
		case int:
			str += fmt.Sprintf("Value: %d", v)
		case string:
			str += fmt.Sprintf("Value: %s", v)
		case bool:
			str += fmt.Sprintf("Value: %t", v)
		case byte:
			str += fmt.Sprintf("Value: %c", v)
		default:
			str += fmt.Sprintf("Unknown type: %T", v)
		}
	}

	fmt.Printf("%s\n", str)
}

func PrintTokenAndSleep(token token.Token, seconds int64) {
	PrintToken(token)
	time.Sleep(time.Second * time.Duration(seconds))
}
