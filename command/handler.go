package command

import (
	"strconv"
	"fmt"
)

func HandleCommands(options Commands, args []string) {
	for _, arg := range args {
		fmt.Println(arg)
	}
	fmt.Println("Verbose: " + strconv.FormatBool(options.Verbose))
}
