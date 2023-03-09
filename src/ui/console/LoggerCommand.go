package ui_console

import (
	"flag"
	"fmt"
)

type LoggerCommand struct {
	flagSet *flag.FlagSet
	filter  string
}

func NewLoggerCommand() *LoggerCommand {
	lc := &LoggerCommand{
		flagSet: flag.NewFlagSet("logs", flag.ContinueOnError),
	}

	lc.flagSet.StringVar(&lc.filter, "filter", "", "Filter logs for site")

	return lc
}

func (lc *LoggerCommand) Name() string {
	return lc.flagSet.Name()
}

func (lc *LoggerCommand) Init(args []string) error {
	return lc.flagSet.Parse(args)
}

func (lc *LoggerCommand) Exec() error {
	fmt.Println("EXEC LOGGER")
	fmt.Println("filter:", lc.filter)
	return nil
}
