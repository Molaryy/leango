package Token

import (
	"fmt"
	"leango/Logger"
	"leango/Variable"
	"slices"
	"strconv"
)

var availableTokens = []string{"let", "const", "lean"}
var availableVariables = make(map[string]Variable.Variable)

func TokenHandler() map[string]func([]string, int) {
	var TokenMapFunction = make(map[string]func([]string, int))

	TokenMapFunction["let"] = isLet
	TokenMapFunction["const"] = isConst

	return TokenMapFunction
}

func isLet(rows []string, nbLine int) {
	indexEqualSign := slices.Index(rows, "=")
	lenRows := len(rows)
	var tokenValue []string

	if lenRows < 2 {
		Logger.Fatal(nbLine, "Nothing found after the declaration of let")
	}
	fmt.Println(indexEqualSign)
	if indexEqualSign == 2 && lenRows == 3 {
		Logger.Fatal(nbLine, fmt.Sprintf("A value was not set for %s", rows[1]))
	}
	if _, varNameExists := availableVariables[rows[1]]; varNameExists {
		Logger.Fatal(nbLine, fmt.Sprintf("A variable aleady exists with the same name %s: line %d", rows[1], nbLine))
	}
	if indexEqualSign == 2 {
		tokenValue = rows[indexEqualSign+1:]

		if intValue, err := strconv.Atoi(tokenValue[0]); err == nil {
			availableVariables[rows[1]] = Variable.Variable{IsNumeric: true, Value: intValue}
		} else {
			availableVariables[rows[1]] = Variable.Variable{IsNumeric: true, Value: intValue}
		}
	} else {
		availableVariables[rows[1]] = Variable.Variable{Value: nil}
	}
}

func isConst(rows []string, line int) {
	fmt.Printf("Index of = is at %d\n", slices.Index(rows, "="))
	fmt.Println("Is a const variable")
}
