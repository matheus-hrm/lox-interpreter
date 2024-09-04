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
	if command != "tokenize" && command != "parse" {
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

	switch command {
	case "tokenize":
		for _, token := range tokens {
			switch token.Type {
			case "EOF":
				fmt.Printf("%s  %s\n", token.Type, token.Literal)
			case "STRING":
				fmt.Printf("%s %s %s\n", token.Type, token.Lexeme, token.Literal)
			case "NUMBER":
				fmt.Printf("%s %s %s\n", token.Type, token.Lexeme, token.Literal)
			case "IDENTIFIER":
				fmt.Printf("%s %s %s\n", token.Type, token.Lexeme, token.Literal)
			default:
				fmt.Printf("%s %s %s\n", token.Type, token.Lexeme, token.Literal)
			}
		}

		if len(scanner.Errors) > 0 {
			for _, err := range scanner.Errors {
				fmt.Fprintln(os.Stderr, err)
			}
			os.Exit(LexicalError)
		}
		os.Exit(0)

	case "parse":
		parser := NewParser(fileContents, tokens)
		ast, err := parser.Parse()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, node := range ast.Nodes {
			fmt.Println(node.String())
		}
	}
}
