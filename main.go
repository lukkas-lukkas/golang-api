package main

import (
	"fmt"
	"os"

	"github.com/lukkas-lukkas/go-api-rest/src/ui"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("You must pass the command")
		os.Exit(1)
	}

	subCommand := args[0]

	err := ui.Dispatch(subCommand, args[1:])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
