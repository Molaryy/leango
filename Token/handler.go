package Token

import (
	"fmt"
	"leango/Logger"
	"leango/Variable"
	"slices"
	"strconv"
	"strings"
)

var availableTokens = []string{"let", "const", "lean"}
var availableVariables = make(map[string]Variable.Variable)

func Handler() map[string]func([]string, int) {
	var TokenMapFunction = make(map[string]func([]string, int))

	TokenMapFunction["let"] = isLet
	TokenMapFunction["const"] = isConst

	return TokenMapFunction
}

func handleString(nbLine int, value string) (string, bool) {
	if value[0] == '"' && strings.HasSuffix(value, "\"") {
		fmt.Println(value)
		return strings.TrimSuffix(strings.TrimLeft(value, "\""), "\""), true
	} else if value[0] == '"' && !strings.HasSuffix(value, "\"") ||
		value[0] != '"' && strings.HasSuffix(value, "\"") {
		Logger.Fatal(nbLine, "Value must start with a \" and end with a \"")
	}
	return "", false
}

func handleVariableType(nbLine int, rows []string, indexEqualSign int) {
	var tokenValue = rows[indexEqualSign+1:]

	if intValue, err := strconv.Atoi(tokenValue[0]); err == nil {
		availableVariables[rows[1]] = Variable.Variable{IsNumeric: true, Value: intValue}
	} else if strValue, isStr := handleString(nbLine, tokenValue[0]); isStr == true {
		availableVariables[rows[1]] = Variable.Variable{IsString: true, Value: strValue}
	}
}

func isLet(rows []string, nbLine int) {
	indexEqualSign := slices.Index(rows, "=")
	lenRows := len(rows)

	if lenRows < 2 {
		Logger.Fatal(nbLine, "Nothing found after the declaration of let")
	}
	if indexEqualSign == 2 && lenRows == 3 {
		Logger.Fatal(nbLine, fmt.Sprintf("A value was not set for %s", rows[1]))
	}
	if _, varNameExists := availableVariables[rows[1]]; varNameExists {
		Logger.Fatal(nbLine, fmt.Sprintf("A variable aleady exists with the same name %s at", rows[1]))
	}
	if indexEqualSign == 2 {
		handleVariableType(nbLine, rows, indexEqualSign)
	} else {
		availableVariables[rows[1]] = Variable.Variable{Value: nil}
	}
}

func isConst(rows []string, line int) {
	fmt.Println("Is a const variable")
}
