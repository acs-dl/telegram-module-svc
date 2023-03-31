package main

import (
	"os"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
