package scanner

import (
	"fmt"
	arguments "leango/src/Args"
	"leango/src/Token"
)

func getSupportedKeywords() []string {
	return []string {
		"let",
		"const",
		"function",
		"for",
		"return",
	}
}

func getSupportedTokenTypes() []string {
	return []string {
		"KEYWORD",
		"IDENTIFIER",
		"ASSIGN",
		"VALUE",
		"DELIMITER",
	}
}

const (
	delimiters = "{}[]();=+*/"
	
)

func ScanFile(flags map[string]arguments.Flag, file arguments.File) {
	var tokens []Token.Token

	for  _, b := range file.Src {
		switch b {
			case '{', '}', '[', ']', '(', ')', ';', '=', '+', '*', '/', '-':
				tokens = append(tokens, Token.Token{Type: "DELIMITER", Value: b})
				continue
			}
	}
	fmt.Println(tokens)
}

