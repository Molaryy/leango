package main

import (
	"fmt"
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
	var tokenFunctions map[string]func([]string, int) = Token.TokenHandler()
	data, err := os.ReadFile(filePath)

	checkErr(err)
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		tokens = strings.FieldsFunc(line, Split)
		for idx, token := range tokens {
			if Token.IsTokenAvailable(token) {
				tokenFunctions[token](tokens, idx+1)
			}
			fmt.Println(token)
		}
		fmt.Println("")
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Default().Fatal("Not enough arguments\n./leango [filepath]")
	}
	parseFile(os.Args[1])
}
