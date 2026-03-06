package operators

import (
	"strconv"
)

func calculate(vars map[string]string) {
	// TODO: Read expressions %var1%+%var2%*%var3%
}

func getVariable(name string, vars map[string]string) int {
	value, err := strconv.Atoi(vars[name])

	if err != nil {
		panic(err)
	}

	return value
}

func operation(operator string, x, y int) int {
	switch operator {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	case "%":
		return x % y
	default:
		panic("unknown operator: " + operator)
	}
}
