package commands

import "strings"

var conditions = make(map[string][]string) // If and else if
var noCondition string                     // Else
var currentCondition string

func IsIf(line string) bool {
	return strings.HasPrefix(line, "if")
}

func IsElseIf(line string) bool {
	return strings.HasPrefix(line, "else if")
}

func IsElse(line string) bool {
	return strings.HasPrefix(line, "else")
}

func ExecuteIf(line string) bool {
	if IsIf(line) {
		// TODO
	} else if IsElseIf(line) {
		// TODO
	} else if IsElse(line) {
		// TODO
	} else {
		conditions[currentCondition] = append(conditions[currentCondition], strings.TrimSpace(line))
	}
	return false
}
