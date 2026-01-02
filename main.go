package main

import (
	"fmt"
	arguments "leango/src/Args"
	"leango/src/Scanner"
	"os"
)

func main() {
	existingFlags := make(map[string]arguments.Flag)

	existingFlags["--debug"] = arguments.Flag {
			ExpectsValue: false,
			Description:  "debugging mode is enabled for leango",
	}
	args := os.Args[1:]

	arguments, err := arguments.GetArguments(existingFlags, args)
	if err != nil {
		// TODO: improve logging
		fmt.Println(err)
		return
	}

	// TODO: at the moment, if a flag is set with the shorthanded one, both of them will be stored
	flags := arguments.Flags
	files := arguments.Files

	for _, file := range files {
		scanner.ScanFile(flags, file)
	}
}

