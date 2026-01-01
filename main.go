package main

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	arguments "leango/src/Args"
	"leango/src/Token"
	"os"
	"strings"
)

func Split(r rune) bool {
	return r == ' ' || r == '\t'
}

func parseFile(filePath string) {
	var tokens []string
	var tokenFunctions map[string]func([]string, int) = Token.Handler()
	data, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
		return
	}
	lines := strings.Split(string(data), "\n")

	for idxLine, line := range lines {
		if line != "" {
			tokens = strings.FieldsFunc(line, Split)
			for idx, token := range tokens {
				_, varNameExists := tokenFunctions[token]
				if idx > 0 && varNameExists {
					logger.Fatal(idxLine, "you can't name a variable with the same name as a type")
				}
				if Token.IsTokenAvailable(token) {
					tokenFunctions[token](tokens, idx+1)
				}
			}
		}
	}
	Token.ShowVariables()
}

func checkFileSuffix(filename string) bool {
	if !strings.HasSuffix(filename, ".leango") {
		logger.WithFields(logger.Fields{}).Fatal("incorrect file extension, should match .leango")
		return false
	}
	return true
}

func main() {
	existingFlags := []arguments.Flag{
		{
			Name:         "--debug",
			Shorthand:    "-d",
			ExpectsValue: false,
			Description:  "debugging mode is enabled for leango",
		},
	}
	args := os.Args[1:]
	arguments, err := arguments.GetArguments(existingFlags, args)
	if err != nil {
		logger.Fatal(err)
	}
	flags := arguments.Flags
	files := arguments.Files


	fmt.Println(flags)
	fmt.Println(files)

	/*if checkFileSuffix(os.Args[1]) {
	    parseFile(os.Args[1])
	}*/
}
