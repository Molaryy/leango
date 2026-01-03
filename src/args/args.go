package args

import (
	"fmt"
	"os"
	"strings"
)

type Flag struct {
	ExpectsValue bool
	Description  string
}

type File struct {
	Filepath string
	Src      []byte
}

type Arguments struct {
	Flags            map[string]Flag
	Files            []File
	HasProvidedFiles bool
}

func GetArguments(existingFlags map[string]Flag, args []string) (Arguments, error) {
	flags := map[string]Flag{}
	files := []File{}
	hasProvidedFiles := false

	for _, arg := range args {
		flag, exists := existingFlags[arg]
		if exists {
			flags[arg] = flag
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

			files = append(files, File{Filepath: arg, Src: content})

			if !hasProvidedFiles {
				hasProvidedFiles = true
			}
		} else {
			return Arguments{}, fmt.Errorf("named files must be .leango files: %s", arg)
		}
	}

	return Arguments{Flags: flags, Files: files, HasProvidedFiles: hasProvidedFiles}, nil
}
