package commands

import "strings"

var vars = make(map[string]any)

func Set(args string) {
	parameter := ""

	if strings.HasPrefix(args, "/") {
		rawParameter := strings.Split(args, " ")[0]
		parameter = strings.ToLower(rawParameter)
	}

	switch parameter {
	case "/a":
		panic("Not implemented yet")
	case "/p":
		panic("Not implemented yet")
	case "":
		panic("Not implemented yet")
	default:
		panic("Unknown parameter: " + parameter)
	}
}
