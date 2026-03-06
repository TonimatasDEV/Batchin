package internal

import (
	"batch-interpreter/internal/commands"
	"strings"
)

var vars = make(map[string]string)

func ExecuteFunc(line string) {
	lineSplit := strings.SplitN(line, " ", 2)
	command := lineSplit[0]
	args := ""

	if len(lineSplit) > 1 {
		args = lineSplit[1]
	}

	switch command {
	case "Rem":
		break
	case "@echo":
		commands.SetEcho(args)
	case "echo":
		commands.PrintEcho(args)
	case "set":
		vars = commands.Set(args, vars)
	default:
		panic("Unknown command: " + command)
	}
}
