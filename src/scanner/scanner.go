package scanner

import (
	"fmt"
	"leango/pkg/debugger"
	arguments "leango/src/args"
	"leango/src/token"
	"math"
	"strings"
)

func scanDelimiterAndOperator(b byte) (token.Token, bool) {
	var tok token.Token
	found := false

	switch b {
	case '{':
		tok = token.Token{Type: "DELIMITER_OPEN_BRACE"}
	case '}':
		tok = token.Token{Type: "DELIMITER_CLOSE_BRACE"}
	case '[':
		tok = token.Token{Type: "DELIMITER_OPEN_BRACKET"}
	case ']':
		tok = token.Token{Type: "DELIMITER_CLOSE_BRACKET"}
	case '(':
		tok = token.Token{Type: "DELIMITER_OPEN_PARENTHESES"}
	case ')':
		tok = token.Token{Type: "DELIMITER_CLOSE_PARENTHESES"}
	case ';':
		tok = token.Token{Type: "DELIMITER_SEMICOLON"}
	case '=':
		tok = token.Token{Type: "OPERATOR_ASSIGN"}
	case '+':
		tok = token.Token{Type: "OPERATOR_ADDITION"}
	case '-':
		tok = token.Token{Type: "OPERATOR_SUBTRACTION"}
	case '*':
		tok = token.Token{Type: "OPERATOR_MULTIPLICATION"}
	case '/':
		tok = token.Token{Type: "OPERATOR_DIVISION"}
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

func cutWord(indexSrc int, src []byte, lenSrc int) (int, string) {
	var buffer strings.Builder
	finalIndex := indexSrc

	for finalIndex < lenSrc {
		_, stop := scanDelimiterAndOperator(src[finalIndex])
		if src[finalIndex] == ' ' || src[finalIndex] == '"' || src[finalIndex] == '\n' || stop {
			break
		}
		buffer.WriteByte(src[finalIndex])
		finalIndex++
	}
	return finalIndex, buffer.String()
}

func scanKeyword(str string) (token.Token, bool) {
	var tok token.Token
	found := false

	switch str {
	case "let":
		tok = token.Token{Type: "DECLARATION_LET"}
	case "const":
		tok = token.Token{Type: "DECLARATION_CONST"}
	case "fn":
		tok = token.Token{Type: "DECLARATION_FN"}
	case "if":
		tok = token.Token{Type: "STATEMENT_IF"}
	case "while":
		tok = token.Token{Type: "STATEMENT_WHILE"}
	}

	if tok.Type != "" {
		found = true
	}
	return tok, found
}

func scanIntNumber(str string) (token.Token, bool, bool) {
	var tok token.Token
	var foundNumber float64 = 0
	found := false
	sign := 1
	strIndex := 0
	strLen := len(str)
	nbOfSigns := 0
	hasError := false

	if str[0] != '-' && str[0] != '+' && str[0] < '0' || str[0] > '9' {
		return tok, false, false
	}

	if str[0] == '-' {
		sign = -1
		nbOfSigns += 1
	}
	for i := 1; i < strLen && str[i] == '-' || str[i] == '+'; i++ {
		if str[0] == '-' {
			sign *= -1
		} else {
			sign *= 1
		}
		strIndex = i
		nbOfSigns += 1
	}
	remainingLen := strLen - nbOfSigns
	var unexpextedStr strings.Builder
	for strIndex < strLen && remainingLen > 0 {
		if str[strIndex] < '0' || str[strIndex] > '9' && hasError == false {
			hasError = true
		}
		if hasError {
			unexpextedStr.WriteByte(str[strIndex])
		} else {
			foundNumber += (float64(str[strIndex]) - '0') * math.Pow10(remainingLen-1)
		}
		strIndex++
		remainingLen--
	}

	if hasError {
		fmt.Println("syntax error found unexpexted value: ", unexpextedStr.String())
	} else {
		found = true
		tok = token.Token{
			Type:     "INTEGER",
			Value:    int(foundNumber),
			HasValue: true,
		}
	}
	return tok, found, hasError
}

type StoredDeclarations struct {
	lineNumber int
	name       string
}

func ScanFile(flags map[string]arguments.Flag, file arguments.File) ([]token.Token, bool) {
	var tokens []token.Token
	isReadingString := false
	var tmpToken string = ""
	lenSrc := len(file.Src)
	lastTokenType := ""
	errorFound := false

	for fileIndex := 0; fileIndex < lenSrc; fileIndex++ {
		if file.Src[fileIndex] == '"' {
			if isReadingString == false {
				isReadingString = true
				continue
			} else if fileIndex-1 > 0 && file.Src[fileIndex-1] != '\\' {
				tok := token.Token{Type: "VALUE_STRING", Value: tmpToken, HasValue: true}
				tokens = append(tokens, tok)
				tmpToken = ""
				isReadingString = false
				if tok.Type != "" {
					lastTokenType = tok.Type
				}
				continue
			}
		}
		if isReadingString {
			tmpToken += string(file.Src[fileIndex])
			continue
		}
		tok, ok := scanDelimiterAndOperator(file.Src[fileIndex])
		if tok.Type != "" {
			lastTokenType = tok.Type
		}
		if ok {
			tokens = append(tokens, tok)
		} else {
			idx, word := cutWord(fileIndex, file.Src, lenSrc)
			if idx != fileIndex {
				fileIndex = idx
				tok, found := scanKeyword(word)
				if found {
					tokens = append(tokens, tok)
					if tok.Type != "" {
						lastTokenType = tok.Type
					}
					continue
				}
				tok, found, err := scanIntNumber(word)
				if err {
					errorFound = true
					break
				}
				if found {
					tokens = append(tokens, tok)
					if tok.Type != "" {
						lastTokenType = tok.Type
					}
					continue
				}
				if lastTokenType == "OPERATOR_ASSIGN" ||
					lastTokenType == "OPERATOR_ADDITION" ||
					lastTokenType == "OPERATOR_MULTIPLICATION" ||
					lastTokenType == "OPERATOR_SUBTRACTION" ||
					lastTokenType == "OPERATOR_DIVISION" ||
					lastTokenType == "DECLARATION_LET" ||
					lastTokenType == "DECLARATION_CONST" ||
					lastTokenType == "DECLARATION_FN" {
					tokens = append(tokens, token.Token{Type: "IDENTIFIER", IsIdentifier: true, IdentifierName: word})
				} else {
					errorFound = true
					fmt.Printf("unexpected name %s\n", word)
					break
				}
			}
		}
	}

	if debugger.IsDebugActivated(flags) {
		for _, t := range tokens {
			debugger.PrintToken(t)
		}
	}

	return tokens, errorFound
}
