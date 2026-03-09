package commands

import (
	"os"
	"strconv"
	"strings"
)

func ExecuteExit(args string) {
	if strings.HasPrefix(args, "/b") { // Remove /b, should be implemented?
		args = strings.TrimPrefix(args, "/b")
		args = strings.TrimSpace(args)
	}

	value, err := strconv.Atoi(args)
	if err != nil {
		println("Error code should be an integer.")
	}

	os.Exit(value)
}
