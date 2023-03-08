package main

import (
	"fmt"
	"os"

	"github.com/lukkas-lukkas/go-api-rest/src/domain"
	console "github.com/lukkas-lukkas/go-api-rest/src/ui/console"
)

func main() {
	args := os.Args[1:]

	commands := []domain.Command{
		console.NewMonitorCommand(),
	}

	if len(args) < 1 {
		fmt.Println("You must pass the command")
		os.Exit(1)
	}

	subCommand := args[0]

	for _, cmd := range commands {
		if cmd.Name() == subCommand {
			cmd.Init(os.Args[2:])
			err := cmd.Exec()
			if err != nil {
				fmt.Println(err)
			}
			os.Exit(0)
		}
	}

	fmt.Println("ERROR: Command not found:", subCommand)
	os.Exit(1)
}
