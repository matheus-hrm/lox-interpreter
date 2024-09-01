package main

import (
	"fmt"
	"os"
)

const (
	EOF = "EOF"
	LEFT_PAREN rune = '('
	RIGHT_PAREN rune = ')'
)

func main() {
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

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
	for _, current := range fileContents {
		switch current {
		case LEFT_PAREN:
			fmt.Println("LEFT_PAREN ( null")
		case RIGHT_PAREN:
			fmt.Println("RIGHT_PAREN ) null")
		}
	}
	fmt.Println("EOF  null")
}
