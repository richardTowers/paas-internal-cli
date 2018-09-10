package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/richardTowers/paas-internal-cli/command"
)

func main() {
	err := Main(os.Args[1:])
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func Main(args []string) error {
	var commands command.Commands
	parser := flags.NewParser(&commands, flags.Default)
	args, err := parser.ParseArgs(args)

	if isHelp(err) {
		return nil
	} else if err != nil {
		var b bytes.Buffer
		parser.WriteHelp(&b)
		return fmt.Errorf(b.String())
	}

	command.HandleCommands(commands, args)

	return nil
}

func isHelp(err error) bool {
	flagsErr, isFlagsErr := err.(*flags.Error)
	return isFlagsErr && flagsErr.Type == flags.ErrHelp
}
