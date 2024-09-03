package main

type Token struct {
	Type    string
	Lexeme  string
	Literal interface{}
	Line    int
}

type TokenType string

const (
	EOF           TokenType = "EOF"
	LEFT_PAREN    TokenType = "("
	RIGHT_PAREN   TokenType = ")"
	LEFT_BRACE    TokenType = "{"
	RIGHT_BRACE   TokenType = "}"
	STAR          TokenType = "*"
	DOT           TokenType = "."
	PLUS          TokenType = "+"
	MINUS         TokenType = "-"
	COMMA         TokenType = ","
	SEMICOLON     TokenType = ";"
	EQUAL         TokenType = "="
	EQUAL_EQUAL   TokenType = "=="
	BANG          TokenType = "!"
	BANG_EQUAL    TokenType = "!="
	LESS          TokenType = "<"
	GREATER       TokenType = ">"
	LESS_EQUAL    TokenType = "<="
	GREATER_EQUAL TokenType = ">="
	SLASH         TokenType = "/"
	WHITESPACE    TokenType = " "
	TAB           TokenType = "\t"
	NEWLINE       TokenType = "\n"
	DOUBLE_QUOTE  TokenType = "\""
	AND           TokenType = "AND"
	CLASS         TokenType = "CLASS"
	ELSE          TokenType = "ELSE"
	FALSE         TokenType = "FALSE"
	FUN           TokenType = "FUN"
	FOR           TokenType = "FOR"
	IF            TokenType = "IF"
	NIL           TokenType = "NIL"
	OR            TokenType = "OR"
	PRINT         TokenType = "PRINT"
	RETURN        TokenType = "RETURN"
	SUPER         TokenType = "SUPER"
	THIS          TokenType = "THIS"
	TRUE          TokenType = "TRUE"
	VAR           TokenType = "VAR"
	WHILE         TokenType = "WHILE"
)

const (
	STRING     TokenType = "STRING"
	IDENTIFIER TokenType = "IDENTIFIER"
	NUMBER     TokenType = "NUMBER"
)

var TokenMap = map[string]string{
	"(":      "LEFT_PAREN",
	")":      "RIGHT_PAREN",
	"{":      "LEFT_BRACE",
	"}":      "RIGHT_BRACE",
	"*":      "STAR",
	".":      "DOT",
	"+":      "PLUS",
	"-":      "MINUS",
	",":      "COMMA",
	";":      "SEMICOLON",
	"=":      "EQUAL",
	"==":     "EQUAL_EQUAL",
	"!":      "BANG",
	"!=":     "BANG_EQUAL",
	"<":      "LESS",
	">":      "GREATER",
	"<=":     "LESS_EQUAL",
	">=":     "GREATER_EQUAL",
	"/":      "SLASH",
	" ":      "WHITESPACE",
	"\t":     "TAB",
	"\n":     "NEWLINE",
	"STRING": "STRING",
}
var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}
