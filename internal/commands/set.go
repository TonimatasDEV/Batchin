package commands

import (
	"batch-interpreter/internal/operators"
	"batch-interpreter/internal/variables"
	"bufio"
	"os"
	"strings"
)

func Set(args string) {
	parameter, rawVariable := getData(args)

	switch parameter {
	case "/a":
		name, value := variables.GetRawVariable(rawVariable)
		parsedValue := variables.ParseLine(value)
		result := operators.Calculate(parsedValue)
		variables.SetVariable(name, result)
	case "/p":
		name, value := variables.GetRawVariable(rawVariable)
		toPrint := strings.ReplaceAll(value, "\r", "")
		print(toPrint)

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			line := scanner.Text()
			variables.SetVariable(name, line)
		}

		if scanner.Err() != nil {
			panic(scanner.Err())
		}
	default:
		panic("Unknown parameter: " + parameter)
	}
}

func getData(args string) (parameter string, rawVariable string) {
	if strings.HasPrefix(args, "/") {
		rawParameter := strings.SplitN(args, " ", 2)
		parameter = strings.ToLower(rawParameter[0])
		rawVariable = rawParameter[1]
	} else {
		rawVariable = args
	}

	return parameter, rawVariable
}
