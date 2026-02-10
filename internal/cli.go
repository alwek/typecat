package internal

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func ParseArgs(args []string) (string, error) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: typecat [options] <path-to-file>\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	// The file path would be a "positional argument" (the part after flags)
	args = flag.Args()
	if len(args) < 1 {
		return "", fmt.Errorf("File not specified")
	}

	path := strings.TrimSpace(args[0])
	if path == "" {
		return "", fmt.Errorf("File path cannot be empty or whitespace")
	}
	return path, nil
}
