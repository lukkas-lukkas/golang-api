package main

import (
	"flag"
	"fmt"
	"os"
)

type MonitorCommand struct {
	flagSet *flag.FlagSet
	tries   int
	delay   int
}

func (mc *MonitorCommand) Name() string {
	return mc.flagSet.Name()
}

func (mc *MonitorCommand) Init(args []string) error {
	return mc.flagSet.Parse(args)
}

func (mc *MonitorCommand) Run() error {
	fmt.Println("RUN MONITOR")
	fmt.Println("tries: ", mc.tries)
	fmt.Println("delay: ", mc.delay)
	return nil
}

func NewMonitorCommand() *MonitorCommand {
	mc := &MonitorCommand{
		flagSet: flag.NewFlagSet("monitor", flag.ContinueOnError),
	}

	mc.flagSet.IntVar(&mc.tries, "tries", 5, "Tries to each site")
	mc.flagSet.IntVar(&mc.delay, "delay", 5, "Delay between tries")

	return mc
}

type CommandRunner interface {
	Init([]string) error
	Run() error
	Name() string
}

func main() {
	args := os.Args[1:]

	commands := []CommandRunner{
		NewMonitorCommand(),
	}

	if len(args) < 1 {
		fmt.Println("You must pass the command")
		os.Exit(1)
	}

	subCommand := args[0]

	for _, cmd := range commands {
		if cmd.Name() == subCommand {
			cmd.Init(os.Args[2:])
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
			os.Exit(0)
		}
	}

	fmt.Println("ERROR: Command not found: ", subCommand)
	os.Exit(1)
}
