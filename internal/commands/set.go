package commands

import (
	"bufio"
	"os"
	"strings"
)

func Set(args string, vars map[string]string) map[string]string {
	parameter, rawVariable := getData(args)
	name, value := getVariable(rawVariable)

	switch parameter {
	case "/a":
		vars[name] = value
	case "/p":
		print(value)
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line := scanner.Text()
			vars[name] = line
		}
	default:
		panic("Unknown parameter: " + parameter)
	}

	return vars
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

func getVariable(rawVariable string) (name string, value string) {
	variableSplit := strings.SplitN(rawVariable, "=", 2)

	name = strings.ToLower(variableSplit[0])
	value = variableSplit[1]
	return name, value
}
