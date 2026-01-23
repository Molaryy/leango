package main

import (
	"fmt"
	"leango/pkg/debugger"
	"leango/pkg/helper"
	arguments "leango/src/args"
	"leango/src/scanner"
	"os"
)

func main() {
	existingFlags := make(map[string]arguments.Flag)

	existingFlags["--debug"] = arguments.Flag{
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

	flags := arguments.Flags
	files := arguments.Files

	_, ok := flags["--help"]
	if ok || !arguments.HasProvidedFiles {
		helper.Helper()
		return
	}

	for _, file := range files {
		if debugger.IsDebugActivated(flags) {
			fmt.Printf("Scanning %s%s%s\n", debugger.BLUE, file.Filepath, debugger.RESET)
		}
		scanner.ScanFile(flags, file)
	}
}
