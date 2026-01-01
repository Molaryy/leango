package args

import (
	"fmt"
	"os"
	"strings"
)

type Flag struct {
	Name string
	Shorthand string
	ExpectsValue bool
	Description string
}

type File struct {
	Filepath string
	Src []byte
}

type Arguments struct {
	Flags []Flag
	Files []File
}

// Compares the givent argument with the existingFlags argument and retrieves the result if it exists based on the matching name or shorthanded name
func findFlag(existingFlags []Flag, arg string) (*Flag, bool) {
	for _, flag := range existingFlags {
		if flag.Name == arg || flag.Shorthand == arg {
			return &flag, true
		}
	}
	return nil, false
}

// Get flags & files from leango arguments
func GetArguments(existingFlags []Flag, args []string) (Arguments, error) {
	flags := []Flag{}
	files := []File{}

	for _, arg := range args {
		flag, exists := findFlag(existingFlags, arg)
		if exists {
			flags = append(flags, *flag)
			continue
		}
		info, err := os.Stat(arg)
			if err != nil {
    			return Arguments{}, fmt.Errorf("input file %q: %w", arg, err)
			}
			if info.IsDir() {
    			return Arguments{}, fmt.Errorf("input file %q is a directory", arg)
			}
		if strings.HasSuffix(arg, ".leango") {
			content, err := os.ReadFile(arg)
			if err != nil {
				return Arguments{}, err
			}

			files = append(files, File{Filepath: arg,  Src: content})
		} else {
			return Arguments{}, fmt.Errorf("named files must be .leango files: %s", arg)
		}
	}

	return Arguments{Flags: flags, Files: files}, nil
}

