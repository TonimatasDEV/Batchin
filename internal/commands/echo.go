package commands

import "batch-interpreter/internal/variables"

func SetEcho(args string) {
	if args == "off" {
		variables.EchoOff()
	}
}

func PrintEcho(args string) {
	println(args)
}
