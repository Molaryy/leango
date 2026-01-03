package scanner

import (
	"fmt"
	"leango/pkg/debugger"
	arguments "leango/src/Args"
	"leango/src/Token"
)

func getSupportedKeywords() []string {
	return []string{
		"let",
		"const",
		"function",
		"for",
		"return",
	}
}

func getSupportedTokenTypes() []string {
	return []string{
		"KEYWORD",
		"IDENTIFIER",
		"ASSIGN",
		"VALUE",
		"DELIMITER",
	}
}

func ScanFile(flags map[string]arguments.Flag, file arguments.File) []Token.Token {
	var tokens []Token.Token
	isReadingString := false
	var token string = ""

	for fileIndex, b := range file.Src {
		if isReadingString == false {
			switch b {
			case '{', '}', '[', ']', '(', ')', ';', '=', '+', '*', '/', '-':
				tokens = append(tokens, Token.Token{Type: "DELIMITER", Value: b})
				continue
			}
		}
		if b == '"' {
			if isReadingString == false {
				isReadingString = true
				continue
			} else if fileIndex-1 > 0 && file.Src[fileIndex-1] != '\\' {
				tokens = append(tokens, Token.Token{Type: "STRING_VALUE", Value: token})
				token = ""
				isReadingString = false
				continue
			}
		}
		if isReadingString {
			token += string(b)
		}
	}
	if debugger.IsDebugActivated(flags) {
		for _, t := range tokens {
			debugger.PrintToken(t)
		}
		fmt.Println()
	}

	return tokens
}
