package operators

import "strings"

func calculate(equation string) {
	equation = strings.TrimSpace(equation)
	strings.Count(equation, "=")
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
