package commands

import (
	"batch-interpreter/internal/variables"
	"bufio"
	"os"
	"strings"
)

func Set(args string) {
	parameter, rawVariable := getData(args)

	switch parameter {
	case "/a":
		variable
		variables.SetRawVariable(rawVariable)
	case "/p":
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			line := scanner.Text()
			variables.SetRawVariable(line)
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
