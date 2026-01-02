package scanner

import (
	arguments "leango/src/Args"
	"leango/src/Token"
	"strings"
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

func isDelimiter(c rune) bool {
	if strings.ContainsRune(delimiters, c) {
		return true
	}
	return false
}

func ScanFile(flags []arguments.Flag, file arguments.File) {
	var tokens []Token.Token

	for  _, b := range file.Src {
		if isDelimiter(rune(b)) {
			tokens = append(tokens, Token.Token{Type: "DELIMITER", Value: b})
		}
	}
}

