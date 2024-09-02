package main

import (
	"fmt"
	"os"
)

type Scanner struct {
	Source  string
	Tokens  []Token
	Start   int
	Current int
	Line    int

	ExitCode int
}

const (
	LexicalError = 65
)

func NewScanner(source string) *Scanner {
	return &Scanner{
		Source:   source,
		Tokens:   []Token{},
		Start:    0,
		Current:  0,
		Line:     1,
		ExitCode: 0,
	}
}

func (s *Scanner) ScanTokens() []Token {
	for s.Current < len(s.Source) {
		s.Start = s.Current
		s.ScanToken()
	}
	s.Tokens = append(s.Tokens, Token{Type: "EOF", Lexeme: "", Literal: "", Line: s.Line})

	return s.Tokens
}

func (s *Scanner) ScanToken() {
	c := s.Advance()
	switch TokenType(c) {
	case LEFT_PAREN:
		s.AddToken(LEFT_PAREN)
	case RIGHT_PAREN:
		s.AddToken(RIGHT_PAREN)
	case LEFT_BRACE:
		s.AddToken(LEFT_BRACE)
	case RIGHT_BRACE:
		s.AddToken(RIGHT_BRACE)
	case STAR:
		s.AddToken(STAR)
	case DOT:
		s.AddToken(DOT)
	case PLUS:
		s.AddToken(PLUS)
	case MINUS:
		s.AddToken(MINUS)
	case COMMA:
		s.AddToken(COMMA)
	case SEMICOLON:
		s.AddToken(SEMICOLON)
	case EQUAL:
		s.matchAndAddToken('=', EQUAL_EQUAL, EQUAL)
	case EQUAL_EQUAL:
		s.AddToken(EQUAL_EQUAL)
	case BANG:
		s.matchAndAddToken('=', BANG_EQUAL, BANG)
	case BANG_EQUAL:
		s.AddToken(BANG_EQUAL)
	case LESS:
		s.matchAndAddToken('=', LESS_EQUAL, LESS)
	case GREATER:
		s.matchAndAddToken('=', GREATER_EQUAL, GREATER)
	default:
		fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", s.Line, c)
		s.ExitCode = LexicalError
	}
}

func (s *Scanner) AddToken(tokenType TokenType) {
	lexeme := ""
	if s.Current <= len(s.Source) {
		lexeme = s.Source[s.Start:s.Current]
	}
	s.Tokens = append(s.Tokens, Token{
		Type:    TokenMap[string(tokenType)],
		Lexeme:  lexeme,
		Literal: "",
		Line:    s.Line,
	})
}

func (s *Scanner) matchAndAddToken(expected byte, matchToken, noMatchToken TokenType) {
	if s.Current < len(s.Source) && s.Source[s.Current] == expected {
		s.Advance()
		s.AddToken(matchToken)
	} else {
		s.AddToken(noMatchToken)
	}
}

func (s *Scanner) Advance() byte {
	if s.Current >= len(s.Source) {
		return 0
	}
	c := s.Source[s.Current]
	s.Current++
	return c
}
