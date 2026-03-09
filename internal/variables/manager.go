package variables

import (
	"strconv"
	"strings"
)

var variables = make(map[string]string)
var Echo = true

func ParseLine(line string) string {
	for name := range variables {
		line = ParseVariable(line, name)
	}

	return line
}

func ParseVariable(line, varName string) string {
	toReplace := "%" + varName + "%"
	replacement := variables[varName]
	return strings.ReplaceAll(line, toReplace, replacement)
}

func GetVariable(name string) string {
	return variables[name]
}

func GetNumberVariable(name string) int {
	value, err := strconv.Atoi(variables[name])

	if err != nil {
		panic(err)
	}

	return value
}

func SetVariable(name string, value string) {
	variables[name] = value
}

func SetRawVariable(rawVariable string) {
	variableSplit := strings.SplitN(rawVariable, "=", 2)

	name := strings.ToLower(variableSplit[0])
	value := variableSplit[1]
	variables[name] = value
}

func GetRawVariable(rawVariable string) (string, string) {
	variableSplit := strings.SplitN(rawVariable, "=", 2)

	name := strings.ToLower(variableSplit[0])
	value := variableSplit[1]

	return name, value
}

func IsEquation(value string) bool {
	//equationRegex :=
	return false
}
