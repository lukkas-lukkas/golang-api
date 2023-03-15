package ui

import (
	"fmt"

	"github.com/lukkas-lukkas/go-api-rest/src/application"
	"github.com/lukkas-lukkas/go-api-rest/src/domain"
	"github.com/lukkas-lukkas/go-api-rest/src/infrastructure"
	console "github.com/lukkas-lukkas/go-api-rest/src/ui/console"
)

func commands() []domain.Command {
	logger := infrastructure.NewFileLogger()
	monitorService := application.NewMonitorService(logger)
	return []domain.Command{
		console.NewMonitorCommand(monitorService),
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
