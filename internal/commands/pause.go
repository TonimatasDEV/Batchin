package commands

import (
	"os"
	"time"

	"golang.org/x/term"
)

func ExecutePause() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		println("Error pausing. Waiting 1 second instead.")
		time.Sleep(1 * time.Second)
		return
	}

	print("Press any key to continue...")

	var b = make([]byte, 1)
	_, _ = os.Stdin.Read(b)
	_ = term.Restore(int(os.Stdin.Fd()), oldState)

	println()
}
