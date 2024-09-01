package main

type Token struct {
	Type    string
	Lexeme  string
	Literal string
	Line    int
}

const (
	EOF         = "EOF"
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"
	STAR        = "*"
	DOT         = "."
	PLUS        = "+"
	MINUS       = "-"
	COMMA       = ","
	SEMICOLON   = ";"
)

var TokenMap = map[string]string{
	LEFT_PAREN:  "LEFT_PAREN ( null",
	RIGHT_PAREN: "RIGHT_PAREN ) null",
	LEFT_BRACE:  "LEFT_BRACE { null",
	RIGHT_BRACE: "RIGHT_BRACE } null",
	STAR:        "STAR * null",
	DOT:         "DOT . null",
	PLUS:        "PLUS + null",
	MINUS:       "MINUS - null",
	COMMA:       "COMMA , null",
	SEMICOLON:   "SEMICOLON ; null",
}
