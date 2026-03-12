package internal

import (
	"batch-interpreter/internal/commands"
	"batch-interpreter/internal/variables"
	"strings"
)

func ExecuteFunc(line string) {
	lineSplit := strings.SplitN(line, " ", 2)
	command := strings.TrimSpace(lineSplit[0])
	args := ""

	if len(lineSplit) > 1 {
		args = lineSplit[1]
	}

	if variables.Echo && !strings.Contains(line, "@echo off") {
		println(line)
	}

	switch command {
	case "Rem":
		break
	case "@echo":
		commands.ExecuteEcho(args)
	case "echo":
		commands.ExecuteEcho(args)
	case "set":
		commands.Set(args)
	case "pause":
		commands.ExecutePause()
	case "cls":
		commands.ExecuteCLS()
	case "exit":
		commands.ExecuteExit(args)
	default:
		panic("Unknown command: " + command)
	}
}
