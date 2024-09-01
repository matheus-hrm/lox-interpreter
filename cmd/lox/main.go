package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]
	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	rawfile, err := os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fileContents := string(rawfile)
	hasError := false
	for _, current := range fileContents {
		tokenStr, ok := TokenMap[string(current)]
		if ok {
			fmt.Println(tokenStr)
		} else {
			fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %c\n", current)
			hasError = true
		}
	}
	fmt.Println("EOF  null")
	if !hasError {
		os.Exit(1)
	}
	os.Exit(65)
}
