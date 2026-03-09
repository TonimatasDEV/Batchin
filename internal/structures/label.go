package structures

import (
	"regexp"
	"strings"
)

var labelPattern = regexp.MustCompile(`^:[A-Za-z0-9_]*$`)
var gotoPattern = regexp.MustCompile(`\s*goto\s[A-Za-z0-9_]*$`)
var labels = make(map[string]int)

func ExecuteGoto(line string) int {
	line = gotoPattern.FindString(line)
	split := strings.Split(line, " ")
	return labels[split[1]]
}

func SetLabel(label string, line int) {
	label = strings.TrimSpace(label)
	label = strings.TrimPrefix(label, ":")
	labels[label] = line
}

func IsGoto(line string) bool {
	return gotoPattern.MatchString(line)
}

func IsLabel(line string) bool {
	return labelPattern.MatchString(line)
}
