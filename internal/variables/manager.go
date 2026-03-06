package variables

import (
	"strconv"
	"strings"
)

var variables = make(map[string]string)
var Echo = true

func ParseVariable(percentageName string) string {
	nameWithoutPercentage := strings.ReplaceAll(percentageName, "%", "")
	return variables[nameWithoutPercentage]
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

func SetRawVariable(rawVariable string) {
	variableSplit := strings.SplitN(rawVariable, "=", 2)

	name := strings.ToLower(variableSplit[0])
	value := variableSplit[1]
	variables[name] = value
}
