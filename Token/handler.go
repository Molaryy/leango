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
var singleVariableValue = 1

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

func getArithmeticOperation(firstValue int, operator string, secondValue int, nbLine int) int {
    switch operator {
        case "+": return firstValue + secondValue
        case "-": return firstValue - secondValue
        case "*": return firstValue * secondValue
        case "/":
            if secondValue == 0 {
                Logger.Fatal(nbLine, "You tried to divise by 0 at ")
            }
            return firstValue / secondValue
        case "%": return firstValue % secondValue
    }
    Logger.Fatal(nbLine, "An error has occured")
    return 0
}

func handleNumericOperations(rows []string, tokens []string, nbLine int) {
    tokensLen := len(tokens)
    sum := 0
    firstOp := true
    var opt string

    for idx := 0; idx < tokensLen;  {
        fstValue, err := strconv.Atoi(tokens[idx])
        if err != nil {
            Logger.Fatal(nbLine, "Trying to do arithmetic operation with another type than an Int")
        }
        if idx + 1 >= tokensLen {
            break
        }
        opt = tokens[idx + 1]
        if idx + 2 >= tokensLen {
            Logger.Fatal(nbLine, "There is no value after the aritmetic operation")
        }
        scdValue, err := strconv.Atoi(tokens[idx + 2])
        if err != nil {
            Logger.Fatal(nbLine, "Trying to do arithmetic operation with another type than an Int")
        }
        if firstOp {
            sum += getArithmeticOperation(fstValue, opt, scdValue, nbLine)
            firstOp = false
        } else {
            sum = getArithmeticOperation(sum, opt, scdValue, nbLine)
        }
        idx += 2
    }
    availableVariables[rows[1]] = Variable.Variable{IsNumeric: true, Value: sum, IsMutable: true}
}

func handleArithmeticOperations(tokens []string, rows []string, nbLine int) {
    if _, err := strconv.Atoi(tokens[0]); err == nil {
        handleNumericOperations(rows, tokens, nbLine)
	} else if strValue, isStr := handleString(nbLine, tokens[0]); isStr == true {
	   	availableVariables[rows[1]] = Variable.Variable{IsString: true, Value: strValue}
	}
}

func handleVariableType(nbLine int, rows []string, indexEqualSign int) {
	var tokenValues = rows[indexEqualSign+1:]
	
	if (len(tokenValues) == singleVariableValue) {
	   if intValue, err := strconv.Atoi(tokenValues[0]); err == nil {
	   	availableVariables[rows[1]] = Variable.Variable{IsNumeric: true, Value: intValue}
	   } else if strValue, isStr := handleString(nbLine, tokenValues[0]); isStr == true {
	   	availableVariables[rows[1]] = Variable.Variable{IsString: true, Value: strValue}
	   }
	} else {
	   handleArithmeticOperations(tokenValues, rows, nbLine)
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
		availableVariables[rows[1]] = Variable.Variable{Value: nil, IsMutable: true}
	}
}

func isConst(rows []string, line int) {
	fmt.Println("Is a const variable")
}
