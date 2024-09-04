package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	Source  string
	Tokens  []Token
	Current int
}

func NewParser(source string, tokens []Token) *Parser {
	return &Parser{
		Source:  source,
		Tokens:  tokens,
		Current: 0,
	}
}

type AST struct {
	Statements []Stmt
	Nodes      []Expr
}

func (p *Parser) Parse() (*AST, error) {
	ast := &AST{}
	for !p.isAtEnd() {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		ast.Nodes = append(ast.Nodes, expr)
	}
	return ast, nil
}

type ParserError struct {
	Message string
	Token   Token
}

func (e *ParserError) Error() string {
	return e.Message
}

func (e *ParserError) ErrorToken() Token {
	return e.Token
}

func (p *Parser) expression() (Expr, error) {
	return p.assign()
}

type Expr interface {
	expr()
	String() string
}

type Stmt interface {
	stmt()
}

type AssignExpr struct {
	Name  string
	Value Expr
}

func (a AssignExpr) expr() {}

func (a AssignExpr) String() string {
	return fmt.Sprintf("(%s = %s)", a.Name, a.Value.String())
}

type BinaryExpr struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (b BinaryExpr) expr() {}

func (b BinaryExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Operator.Lexeme, b.Left.String(), b.Right.String())
}

type LogicalExpr struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (l LogicalExpr) expr() {}

func (l LogicalExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", l.Operator.Lexeme, l.Left.String(), l.Right.String())
}

type UnaryExpr struct {
	Operator Token
	Right    Expr
}

func (u UnaryExpr) expr() {}

func (u UnaryExpr) String() string {
	return fmt.Sprintf("(%s %s)", u.Operator.Lexeme, u.Right.String())
}

type Literal struct {
	Value interface{}
}

func (l Literal) expr() {}

func (l Literal) String() string {
	switch v := l.Value.(type) {
	case float64:
		val := fmt.Sprintf("%v", l.Value)
		return formatFloat(val, v)
	case string:
		return fmt.Sprintf("%v", strings.Trim(fmt.Sprintf("%v", v), `"`))
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (p *Parser) assign() (Expr, error) {
	expr, err := p.logicalOr()
	if err != nil {
		return nil, err
	}

	if p.match("EQUAL") {
		equals := p.previous()
		value, err := p.assign()
		if err != nil {
			return nil, err
		}

		if assignExpr, ok := expr.(AssignExpr); ok {
			return AssignExpr{
				Name:  assignExpr.Name,
				Value: value,
			}, nil
		}

		return nil, &ParserError{Message: "Invalid assignment target.", Token: equals}
	}

	return expr, nil
}

func (p *Parser) logicalOr() (Expr, error) {
	expr, err := p.logicalAnd()
	if err != nil {
		return nil, err
	}

	for p.match("OR") {
		operator := p.previous()
		right, err := p.logicalAnd()
		if err != nil {
			return nil, err
		}
		expr = LogicalExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) logicalAnd() (Expr, error) {
	expr, err := p.equality()
	if err != nil {
		return nil, err
	}

	for p.match("AND") {
		operator := p.previous()
		right, err := p.equality()
		if err != nil {
			return nil, err
		}
		expr = LogicalExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) equality() (Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match("BANG_EQUAL", "EQUAL_EQUAL") {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		expr = BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) comparison() (Expr, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.match("GREATER", "GREATER_EQUAL", "LESS", "LESS_EQUAL") {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) term() (Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match("MINUS", "PLUS") {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) factor() (Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.match("SLASH", "STAR") {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) unary() (Expr, error) {
	if p.match("BANG", "MINUS") {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return UnaryExpr{Operator: operator, Right: right}, nil
	}

	return p.primary()
}

func (p *Parser) primary() (Expr, error) {
	switch {
	case p.match("FALSE"):
		return Literal{Value: "false"}, nil
	case p.match("TRUE"):
		return Literal{Value: "true"}, nil
	case p.match("NIL"):
		return Literal{Value: "nil"}, nil
	case p.match("NUMBER"):
		num, err := strconv.ParseFloat(p.previous().Lexeme, 64)
		if err != nil {
			return nil, &ParserError{Message: ("Invalid number format: " + p.previous().Lexeme), Token: p.previous()}
		}

		return Literal{Value: num}, nil
	case p.match("STRING"):
		return Literal{Value: p.previous().Lexeme}, nil
	case p.match("LEFT_PAREN"):
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		p.consume("RIGHT_PAREN", "Expect ')' after expression.")
		return UnaryExpr{Operator: Token{Type: "LEFT_PAREN"}, Right: expr}, nil
	default:
		return nil, &ParserError{Message: "Expect expression.", Token: p.peek()}
	}
}

func (p *Parser) match(types ...string) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(t string) bool {
	return !p.isAtEnd() && p.peek().Type == t
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.Current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == "EOF"
}

func (p *Parser) peek() Token {
	return p.Tokens[p.Current]
}

func (p *Parser) previous() Token {
	return p.Tokens[p.Current-1]
}

func (p *Parser) consume(t string, message string) Token {
	if p.check(t) {
		return p.advance()
	}
	panic(&ParserError{Message: message, Token: p.peek()})
}
