package operators

func calculate(vars map[string]string) {
	// TODO: Read expressions %var1%+%var2%*%var3%
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
