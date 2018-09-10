package command

import (
	"strconv"
)

func HandleCommands(options Commands, args []string) {
	for _, arg := range args {
		println(arg)
	}
	println("Verbose: " + strconv.FormatBool(options.Verbose))
}
