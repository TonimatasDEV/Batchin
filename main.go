package main

import (
	"batch-interpreter/internal"
	"batch-interpreter/internal/commands"
	"batch-interpreter/internal/util"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("Ejercicio1.bat")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	ifRequireLine := false
	for _, line := range lines {
		if util.IsBlank(line) {
			continue
		}

		if commands.IsIf(line) || ifRequireLine {
			ifRequireLine = commands.ExecuteIf(line)
			continue
		}

		internal.ExecuteFunc(line)
	}
}
