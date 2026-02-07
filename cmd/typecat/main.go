package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: typecat <path_to_file>")
		return
	}
	var path string = os.Args[1]

	// Convert to absolute path to be sure where we are looking
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("Error resolving path: %v\n", err)
		return
	}

	// Check if the file exists and isn't a directory
	info, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		fmt.Printf("Error: File does not exist at %s\n", absPath)
		return
	}
	if info.IsDir() {
		fmt.Printf("Error: %s is a directory, not a file\n", absPath)
		return
	}

	// 3. Read the file
	data, err := os.ReadFile(absPath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("File content:\n%s", string(data))
}
