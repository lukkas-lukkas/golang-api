package ui_console

import (
	"bufio"
	"flag"
	"io"
	"os"
	"strings"

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
	file, err := os.Open("./sites.txt")

	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)

	var sites []string

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}
	file.Close()

	application.Monitor(sites, mc.tries, mc.delay)
	return nil
}
