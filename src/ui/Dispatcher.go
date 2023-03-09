package ui

import (
	"fmt"

	"github.com/lukkas-lukkas/go-api-rest/src/domain"
	console "github.com/lukkas-lukkas/go-api-rest/src/ui/console"
)

func commands() []domain.Command {
	return []domain.Command{
		console.NewMonitorCommand(),
		console.NewLoggerCommand(),
	}
}

func Dispatch(command string, args []string) error {
	for _, cmd := range commands() {
		if cmd.Name() != command {
			continue
		}

		cmd.Init(args)

		return cmd.Exec()
	}

	return fmt.Errorf("ERROR: Command not found: %q", command)
}
