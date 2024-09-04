package main

import (
	"fmt"
	"strconv"
)

const (
	LexicalError = 65
)

type ScannerError struct {
	Line    int
	Message string
}

func (e *ScannerError) Error() string {
	return fmt.Sprintf("[line %d] Error: %s", e.Line, e.Message)
}

type Scanner struct {
	Source  string
	Tokens  []Token
	Start   int
	Current int
	Line    int
	Errors  []error
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		Source:  source,
		Tokens:  []Token{},
		Start:   0,
		Current: 0,
		Line:    1,
		Errors:  []error{},
	}
}

func (s *Scanner) ScanTokens() []Token {
	for s.Current < len(s.Source) {
		s.Start = s.Current
		s.ScanToken()
	}
	s.Tokens = append(s.Tokens, Token{Type: "EOF", Lexeme: "", Literal: "null", Line: s.Line})

	return s.Tokens
}

func (s *Scanner) ScanToken() {
	c := s.Advance()
	switch c {
	case '(', ')', '{', '}', ',', '.', '-', '+', ';', '*':
		s.AddToken(TokenType(c), nil)
	case '!':
		s.matchAndAddToken('=', BANG_EQUAL, BANG)
	case '=':
		s.matchAndAddToken('=', EQUAL_EQUAL, EQUAL)
	case '<':
		s.matchAndAddToken('=', LESS_EQUAL, LESS)
	case '>':
		s.matchAndAddToken('=', GREATER_EQUAL, GREATER)
	case '/':
		if s.Match('/') {
			for s.Peek() != '\n' && !s.isAtEnd() {
				s.Advance()
			}
		} else {
			s.AddToken(SLASH, nil)
		}
	case ' ', '\r', '\t':
		// ignore whitespace
	case '\n':
		s.Line++
	case '"':
		s.scanString()
	default:
		if isDigit(c) {
			s.scanNumber()
		} else if isAlpha(c) {
			s.scanIdentifier()
		} else {
			s.AddError(fmt.Sprintf("Unexpected character: %c", c))
		}
	}
}

func (s *Scanner) AddToken(tokenType TokenType, literal interface{}) {
	text := s.Source[s.Start:s.Current]
	var literalStr string
	if literal != nil {
		literalStr = fmt.Sprintf("%v", literal)
	} else {
		literalStr = "null"
	}
	tokentypeStr, ok := TokenMap[string(tokenType)]
	if !ok {
		tokentypeStr = string(tokenType)
	}
	s.Tokens = append(s.Tokens, Token{
		Type:    tokentypeStr,
		Lexeme:  text,
		Literal: literalStr,
		Line:    s.Line,
	})
}

func (s *Scanner) scanString() {
	for s.Peek() != '"' && !s.isAtEnd() {
		if s.Peek() == '\n' {
			s.Line++
		}
		s.Advance()
	}

	if s.isAtEnd() {
		s.AddError("Unterminated string.")
		return
	}

	s.Advance()

	value := s.Source[s.Start+1 : s.Current-1]
	s.AddToken(STRING, value)
}

func (s *Scanner) scanNumber() {
	for isDigit(s.Peek()) {
		s.Advance()
	}

	if s.Peek() == '.' && isDigit(s.PeekNext()) {
		s.Advance()
		for isDigit(s.Peek()) {
			s.Advance()
		}
	}

	value := s.Source[s.Start:s.Current]
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		s.AddError(fmt.Sprintf("Invalid number: %s", value))
		return
	}

	fmtdNumber := formatFloat(value, number)
	s.AddToken(NUMBER, fmtdNumber)
}

func (s *Scanner) scanIdentifier() {
	for isAlphaNumeric(s.Peek()) {
		s.Advance()
	}

	text := s.Source[s.Start:s.Current]
	tokenType, ok := keywords[text]
	if !ok {
		tokenType = IDENTIFIER
	}
	s.AddToken(tokenType, nil)
}

func (s *Scanner) AddError(message string) {
	s.Errors = append(s.Errors, &ScannerError{Line: s.Line, Message: message})
}

func (s *Scanner) matchAndAddToken(expected byte, matchToken, noMatchToken TokenType) {
	if s.Match(expected) {
		s.AddToken(matchToken, nil)
	} else {
		s.AddToken(noMatchToken, nil)
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

func (s *Scanner) Match(expected byte) bool {
	if s.isAtEnd() || s.Source[s.Current] != expected {
		return false
	}
	s.Current++
	return true
}

func (s *Scanner) Peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.Source[s.Current]
}

func (s *Scanner) PeekNext() byte {
	if s.Current+1 >= len(s.Source) {
		return 0
	}
	return s.Source[s.Current+1]
}

func (s *Scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}
