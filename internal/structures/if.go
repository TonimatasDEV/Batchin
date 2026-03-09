package structures

import "strings"

type Condition struct {
	Condition string
	Code      string
}

type If struct {
	Conditions []Condition
	ElseCode   string
}

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
	return false
}
