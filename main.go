package main

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"leango/Logger"
	"leango/Token"
	"log"
	"os"
	"strings"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func Split(r rune) bool {
	return r == ' ' || r == '\t'
}

func parseFile(filePath string) {
	var tokens []string
	var tokenFunctions map[string]func([]string, int) = Token.Handler()
	data, err := os.ReadFile(filePath)

	checkErr(err)
	lines := strings.Split(string(data), "\n")

	for idxLine, line := range lines {
		if line != "" {
			tokens = strings.FieldsFunc(line, Split)
			for idx, token := range tokens {
				_, varNameExists := tokenFunctions[token]
				if idx > 0 && varNameExists {
					Logger.Fatal(idxLine, "you can't name a variable with the same name as a type")
				}
				if Token.IsTokenAvailable(token) {
					tokenFunctions[token](tokens, idx+1)
				}
			}
		}
	}
	Token.ShowVariables()
}

func main() {
	if len(os.Args) != 2 {
		log.Default().Fatal("not enough arguments\n./leango [filepath]")
	}
	if !strings.HasSuffix(os.Args[1], ".leango") {
		logger.WithFields(logger.Fields{}).Fatal(fmt.Sprintf("incorrect file extension [%s]", os.Args[1]))
	}

	parseFile(os.Args[1])
}
