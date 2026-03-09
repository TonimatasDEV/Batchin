package main

import (
	"batch-interpreter/internal"
	"batch-interpreter/internal/structures"
	"batch-interpreter/internal/util"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("example.bat")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	ifRequireLine := false

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSuffix(lines[i], "\r")

		if util.IsBlank(line) {
			continue
		}

		if structures.IsIf(line) || ifRequireLine {
			ifRequireLine = structures.ExecuteIf(line)
			continue
		}

		if structures.IsLabel(line) {
			structures.SetLabel(line, i)
			continue
		}

		if structures.IsGoto(line) {
			i = structures.ExecuteGoto(line)
			continue
		}

		internal.ExecuteFunc(line)
	}
}
