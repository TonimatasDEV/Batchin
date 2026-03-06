package internal

import (
	"batch-interpreter/internal/commands"
	"strings"
)

func ExecuteFunc(line string) {
	lineSplit := strings.SplitN(line, " ", 1)
	command := lineSplit[0]
	args := ""

	if len(lineSplit) > 1 {
		args = lineSplit[1]
	}

	switch command {
	case "@echo":
		panic("Not implemented yet")
	case "set":
		commands.Set(args)
	}
}
