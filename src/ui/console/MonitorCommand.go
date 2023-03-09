package ui_console

import (
	"flag"

	"github.com/lukkas-lukkas/go-api-rest/src/application"
)

type MonitorCommand struct {
	flagSet *flag.FlagSet
	tries   int
	delay   int
}

func NewMonitorCommand() *MonitorCommand {
	mc := &MonitorCommand{
		flagSet: flag.NewFlagSet("monitor", flag.ContinueOnError),
	}

	mc.flagSet.IntVar(&mc.tries, "tries", 5, "Tries to each site")
	mc.flagSet.IntVar(&mc.delay, "delay", 5, "Delay between tries")

	return mc
}

func (mc *MonitorCommand) Name() string {
	return mc.flagSet.Name()
}

func (mc *MonitorCommand) Init(args []string) error {
	return mc.flagSet.Parse(args)
}

func (mc *MonitorCommand) Exec() error {
	application.Monitor([]string{"site", "site2"}, mc.tries, mc.delay)
	return nil
}
