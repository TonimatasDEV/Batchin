package commands

import (
	"batch-interpreter/internal/variables"
	"strings"
)

func ExecuteEcho(args string) {
	argsSplit := strings.Split(args, " ")

	if args == "" {
		if variables.Echo {
			println("ECHO is enabled.")
		} else {
			println("ECHO is disabled.")
		}
	} else if len(argsSplit) == 1 {
		noSpacePlease := strings.TrimSpace(argsSplit[0])

		if noSpacePlease == "on" {
			variables.Echo = true
		} else if noSpacePlease == "off" {
			variables.Echo = false
		} else {
			parsedArgs := variables.ParseLine(args)
			println(parsedArgs)
		}
	} else {
		parsedArgs := variables.ParseLine(args)
		println(parsedArgs)
	}
}
