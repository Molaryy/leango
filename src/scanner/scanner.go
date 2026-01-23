package scanner

import (
	"fmt"
	"leango/pkg/debugger"
	arguments "leango/src/Args"
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
	buffer := ""
	finalIndex := indexSrc

	for finalIndex < lenSrc {
		_, stop := scanDelimiterAndOperator(src[finalIndex])
		if src[finalIndex] == ' ' || src[finalIndex] == '"' || src[finalIndex] == '\n' || stop {
			break
		}
		buffer += string(src[finalIndex])
		finalIndex++
	}
	return finalIndex, buffer
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

func scanIntNumber(str string) (token.Token, bool) {
	var tok token.Token
	found := false
	sign := 1
	strIndex := 0
	strLen := len(str)
	nbOfSigns := 0
	var foundNumber float64 = 0
	hasError := false

	if str[0] != '-' && str[0] != '+' && str[0] < '0' || str[0] > '9' {
		return tok, false
	}

	if str[0] == '-' {
		sign = -1
		 nbOfSigns += 1
	}
	for i := 1; i < strLen && str[i] == '-' ||  str[i] == '+'; i++ {
		if str[0] == '-' {
			sign *= -1
		} else {
			sign *= 1
		}
		strIndex = i
		nbOfSigns += 1
	}
	remainingLen:= strLen - nbOfSigns
	var unexpextedStr strings.Builder
	for strIndex < strLen {
		if str[strIndex] < '0' || str[strIndex] > '9' && hasError == false {
			hasError = true
		}
		if remainingLen > 0 && hasError == false {
			foundNumber += (float64(str[strIndex]) - 48) * math.Pow10(remainingLen - 1)
		}
		if hasError {
			unexpextedStr.WriteByte(str[strIndex])
		}
		strIndex++
		remainingLen--
	}

	if hasError {
		fmt.Println("Found unexpextedStr == ", unexpextedStr.String())
	}

	return tok, found
}

func ScanFile(flags map[string]arguments.Flag, file arguments.File) []token.Token {
	var tokens []token.Token
	isReadingString := false
	var tmpToken string = ""
	lenSrc := len(file.Src)

	for fileIndex := 0; fileIndex < lenSrc; fileIndex++ {
		if file.Src[fileIndex] == '"' {
			if isReadingString == false {
				isReadingString = true
				continue
			} else if fileIndex-1 > 0 && file.Src[fileIndex-1] != '\\' {
				tokens = append(tokens, token.Token{Type: "VALUE_STRING", Value: tmpToken, HasValue: true})
				tmpToken = ""
				isReadingString = false
				continue
			}
		}
		if isReadingString {
			tmpToken += string(file.Src[fileIndex])
		}
		if isReadingString == false {
			tok, ok := scanDelimiterAndOperator(file.Src[fileIndex])
			if ok {
				tokens = append(tokens, tok)
			} else {
				idx, word := cutWord(fileIndex, file.Src, lenSrc)
				if idx != fileIndex {
					fmt.Printf("Found word = [%s]\n", word)
					fileIndex = idx
					tok, found := scanKeyword(word)
					if found {
						tokens = append(tokens, tok)
						continue
					}
					tok, found = scanIntNumber(word)
					if found {
						tokens = append(tokens, tok)
						continue
					}
					fmt.Println("We don't handle this word = ", word)

					// There are word tokens that are numbers to check that, just check word[0].isnumeric() or word[0] == '-' || word[0] == '-', until it finds the number
				}
				continue
			}
		}
	}

	if debugger.IsDebugActivated(flags) {
		for _, t := range tokens {
			debugger.PrintToken(t)
		}
	}

	return tokens
}
