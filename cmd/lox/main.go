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
	scanner := NewScanner(fileContents)
	tokens := scanner.ScanTokens()

	hasError := false
	for _, token := range tokens {
		if token.Type == "EOF" {
			fmt.Println("EOF  null")
		} else {
			fmt.Printf("%s %s null\n", token.Type, token.Lexeme)
		}
		if token.Type == "ERROR" {
			fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %s\n", token.Line, token.Lexeme)
			hasError = true
		}
	}

	if hasError {
		os.Exit(65)
	}
	os.Exit(0)
}
