package internal

import (
	"batch-interpreter/internal/commands"
	"strings"
)

func ExecuteFunc(line string) {
	lineSplit := strings.SplitN(line, " ", 2)
	command := strings.TrimSpace(lineSplit[0])
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
		commands.Set(args)
	case "pause":
		commands.ExecutePause()
	case "cls":
		commands.ExecuteCLS()
	default:
		panic("Unknown command: " + command)
	}
}
