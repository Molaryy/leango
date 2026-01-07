package helper

import "fmt"

func Helper() {
	fmt.Printf("Leango is a tool for managing leango source code.\n\n")
	fmt.Println("Usage\n\n\tleango [options] file...")
	fmt.Println("OPTIONS:")
	fmt.Println("--debug\tdebugging mode is enabled for leango")
}
