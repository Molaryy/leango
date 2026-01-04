package scanner

import (
	"fmt"
	"leango/pkg/debugger"
	arguments "leango/src/Args"
	"leango/src/Token"
)

func scanDelimiterAndOperator(b byte) (Token.Token, bool) {
	var tok Token.Token
	found := false

	switch b {
		case '{':
			tok = Token.Token{Type: "DELIMITER_OPEN_BRACE", Value: b}
		case '}':
			tok = Token.Token{Type: "DELIMITER_CLOSE_BRACE", Value: b}
		case '[':
			tok = Token.Token{Type: "DELIMITER_OPEN_BRACKET", Value: b}
		case ']':
			tok = Token.Token{Type: "DELIMITER_CLOSE_BRACKET", Value: b}
		case '(':
			tok = Token.Token{Type: "DELIMITER_OPEN_PARENTHESES", Value: b}
		case ')':
			tok = Token.Token{Type: "DELIMITER_CLOSE_PARENTHESES", Value: b}
		case ';':
			tok = Token.Token{Type: "DELIMITER_SEMICOLON", Value: b}
		case '=':
			tok = Token.Token{Type: "OPERATOR_ASSIGN", Value: b}
		case '+':
			tok = Token.Token{Type: "OPERATOR_ADDITION", Value: b}
		case '-':
			tok = Token.Token{Type: "OPERATOR_SUBTRACTION", Value: b}
		case '*':
			tok = Token.Token{Type: "OPERATOR_MULTIPLICATION", Value: b}
		case '/':
			tok = Token.Token{Type: "OPERATOR_DIVISION", Value: b}
		}
		if tok.Type != "" {
			found = true
		}
		return tok, found
}

// TODO: having environnment for functions and variables as maps for checking
//       res, ok := env.Variables["foo"]
//       if the foo variable is used without existing then an error will occur
//       if not it will just access it, same thing for functions

func ScanFile(flags map[string]arguments.Flag, file arguments.File) []Token.Token {
	var tokens []Token.Token
	isReadingString := false
	var token string = ""

	for fileIndex, b := range file.Src {
		if isReadingString == false {
			tok, ok := scanDelimiterAndOperator(b)
			if ok {
				tokens = append(tokens, tok)
			}
		}
		if b == '"' {
			if isReadingString == false {
				isReadingString = true
				continue
			} else if fileIndex-1 > 0 && file.Src[fileIndex-1] != '\\' {
				tokens = append(tokens, Token.Token{Type: "VALUE_STRING", Value: token})
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
