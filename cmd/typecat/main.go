package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
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
		fmt.Printf("File does not exist at %s\n", absPath)
		return
	}
	if info.IsDir() {
		fmt.Printf("%s is a directory, not a file\n", absPath)
		return
	}

	// Read the file
	dataByte, err := os.ReadFile(absPath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	data := string(dataByte)

	// Read the dictionary file as dictionary
	dictData, err := os.ReadFile("./data/dictionary.yml")
	if err != nil {
		fmt.Println("Error looking up dictionary: ", err)
		return
	}
	dictionary := make(map[string]string)

	// Unmarshal dictionary data into a map
	err = yaml.Unmarshal(dictData, &dictionary)
	if err != nil {
		fmt.Println("Error looking up dictionary: ", err)
		return
	}

	// This regex [a-zA-Z]+ finds sequences of letters.
	re := regexp.MustCompile(`[a-zA-Z0-9']+`)

	// Replace matches while preserving everything else
	result := re.ReplaceAllStringFunc(data, func(word string) string {
		if replacement, exists := dictionary[strings.ToUpper(word)]; exists {
			return replacement
		}
		return strings.ToUpper(word)
	})

	fmt.Printf("File content:\n%s", result)
}
