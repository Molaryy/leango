package Token

import (
	"fmt"
	"leango/Variable"
	"slices"
)

var availableTokens = []string{"let", "const", "lean"}
var availableMutableVariables = make(map[string]Variable.Variable)
var availableNonMutableVariables = make(map[string]Variable.Variable)

func TokenHandler() map[string]func([]string, int) {
	var TokenMapFunction = make(map[string]func([]string, int))

	TokenMapFunction["let"] = isLet
	TokenMapFunction["const"] = isConst

	return TokenMapFunction
}

func isLet(rows []string, line int) {
	index := slices.Index(rows, "=")

	if len(rows) == 1 {
		panic(fmt.Sprintf("Nothing found after the declaration of let at line %d", line))
	}

	if index == -1 {
		panic(fmt.Sprintf("A value was not set for %s", rows[1]))
	}

	value := rows[index+1:]
	fmt.Println(value)
	fmt.Println("Is a let variable and the value is = ")
}

func isConst(rows []string, line int) {
	fmt.Printf("Index of = is at %d\n", slices.Index(rows, "="))
	fmt.Println("Is a const variable")
}
