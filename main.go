package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

func main() {
	err := Main()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func Main() error {
	var opts struct{}
	parser := flags.NewParser(&opts, flags.Default)
	args, err := parser.Parse()

	if isHelp(err) {
		return nil
	} else if err != nil {
		var b bytes.Buffer
		parser.WriteHelp(&b)
		return fmt.Errorf(b.String())
	}

	for _, arg := range args {
		println(arg)
	}

	return nil
}

func isHelp(err error) bool {
	flagsErr, isFlagsErr := err.(*flags.Error)
	return isFlagsErr && flagsErr.Type == flags.ErrHelp
}
