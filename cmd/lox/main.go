package main

import (
	"fmt"
	"os"
)

const LexicalError = 65

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]
	if command != "tokenize" && command != "parse" && command != "evaluate" {
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
			if parserError, ok := err.(*ParserError); ok {
				fmt.Fprintf(os.Stderr, "Error at line %d: %s\n", parserError.Token.Line, parserError.Message)
				os.Exit(LexicalError)
			}
			os.Exit(1)
		}
		if scannerError, ok := err.(*ScannerError); ok {
			fmt.Fprintf(os.Stderr, "Error at line %d: %s\n", scannerError.Line, scannerError.Message)
			os.Exit(LexicalError)
		}
		if len(scanner.Errors) > 0 {
			for _, err := range scanner.Errors {
				fmt.Fprintln(os.Stderr, err)
			}
			os.Exit(LexicalError)
		}
		for _, node := range ast.Nodes {
			fmt.Println(node.String())
		}
	case "evaluate":
		parser := NewParser(fileContents, tokens)
		ast, err := parser.Parse()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			if parserError, ok := err.(*ParserError); ok {
				fmt.Fprintf(os.Stderr, "Error at line %d: %s\n", parserError.Token.Line, parserError.Message)
				os.Exit(LexicalError)
			}
			os.Exit(1)
		}
		if scannerError, ok := err.(*ScannerError); ok {
			fmt.Fprintf(os.Stderr, "Error at line %d: %s\n", scannerError.Line, scannerError.Message)
			os.Exit(LexicalError)
		}
		if len(scanner.Errors) > 0 {
			for _, err := range scanner.Errors {
				fmt.Fprintln(os.Stderr, err)
			}
			os.Exit(LexicalError)
		}
		evaluator := NewEvaluator(ast)
		res, err := evaluator.Evaluate()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			//error code 70
			if runtimeError, ok := err.(*RuntimeError); ok {
				fmt.Fprintf(os.Stderr, "Error at line %d: %s\n", runtimeError.Token.Line, runtimeError.Message)
				os.Exit(70)
			}
			os.Exit(1)
		}
		for _, r := range res.([]interface{}) {
			fmt.Println(formatOutput(r))
		}
	}
}
