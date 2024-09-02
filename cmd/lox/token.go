package main

type Token struct {
	Type    string
	Lexeme  string
	Literal string
	Line    int
}

type TokenType string

const (
	EOF         TokenType = "EOF"
	LEFT_PAREN  TokenType = "("
	RIGHT_PAREN TokenType = ")"
	LEFT_BRACE  TokenType = "{"
	RIGHT_BRACE TokenType = "}"
	STAR        TokenType = "*"
	DOT         TokenType = "."
	PLUS        TokenType = "+"
	MINUS       TokenType = "-"
	COMMA       TokenType = ","
	SEMICOLON   TokenType = ";"
	EQUAL       TokenType = "="
	EQUAL_EQUAL TokenType = "=="
	BANG        TokenType = "!"
	BANG_EQUAL  TokenType = "!="
)

var TokenMap = map[string]string{
	"(":  "LEFT_PAREN",
	")":  "RIGHT_PAREN",
	"{":  "LEFT_BRACE",
	"}":  "RIGHT_BRACE",
	"*":  "STAR",
	".":  "DOT",
	"+":  "PLUS",
	"-":  "MINUS",
	",":  "COMMA",
	";":  "SEMICOLON",
	"=":  "EQUAL",
	"==": "EQUAL_EQUAL",
	"!":  "BANG",
	"!=": "BANG_EQUAL",
}
